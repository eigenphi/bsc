// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package native

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/eth/tracers/protobuf"
	"github.com/ethereum/go-ethereum/eth/tracers/types"
	"github.com/ethereum/go-ethereum/signer/fourbyte"
	"github.com/holiman/uint256"
	"math/big"
	"strconv"
	"strings"
	"sync/atomic"
)

var labelDb *fourbyte.Database

const (
	LabelTransfer         = "Transfer"
	LabelInternalTransfer = "Internal-Transfer"
)

type OpsCallFrame struct {
	Type            string          `json:"type"`
	Label           string          `json:"label"`
	From            string          `json:"from"`
	To              string          `json:"to,omitempty"`
	ContractCreated string          `json:"contract_created,omitempty"`
	Value           string          `json:"value,omitempty"`
	GasIn           string          `json:"gasIn"`
	GasCost         string          `json:"gasCost"`
	Input           string          `json:"input,omitempty"`
	Error           string          `json:"error,omitempty"`
	Calls           []*OpsCallFrame `json:"calls,omitempty"`
	parent          *OpsCallFrame   `json:"-"`
	code            []byte          `json:"-"` // for calculating CREATE2 contract address
	salt            *uint256.Int    `json:"-"` // for calculating CREATE2 contract address
}

func init() {
	labelDb, _ = fourbyte.New()
	register("OpsTracerNative", NewOpsTracer)
	register("OpsPlainTracerNative", NewOpsPlainTracer)
}

func NewOpsTracer(ctx *tracers.Context, message json.RawMessage) (tracers.Tracer, error) {
	return &OpsTracer{}, nil
}

func NewOpsPlainTracer(ctx *tracers.Context, message json.RawMessage) (tracers.Tracer, error) {
	return &OpsTracer{plain: true}, nil
}

type OpsTracer struct {
	plain bool

	evm          *vm.EVM
	callstack    OpsCallFrame
	currentDepth int
	currentFrame *OpsCallFrame
	interrupt    uint32 // Atomic flag to signal execution interruption
	reason       error  // Textual reason for the interruption
	initialized  bool
}

func (t *OpsTracer) CaptureTxStart(gasLimit uint64) {
	return
}

func (t *OpsTracer) CaptureTxEnd(restGas uint64) {
	return
}

func (t *OpsTracer) CaptureStart(env *vm.EVM, from common.Address, to common.Address, create bool, input []byte, gas uint64, value *big.Int) {
	if t.initialized {
		return
	}
	t.evm = env
	t.callstack = OpsCallFrame{
		Type:  "CALL",
		From:  addrToHex(from),
		To:    addrToHex(to),
		GasIn: uintToHex(gas),
		Value: bigToHex(value),
		Input: bytesToHex(input),
	}
	if create {
		t.callstack.Type = "CREATE"
	}
	t.currentDepth = env.Depth() + 1 // depth is the value before "CALL" or "CREATE"
	t.currentFrame = &t.callstack
	t.initialized = true
	return
}

// Note the result has no "0x" prefix
func getLogValueHex(stack *vm.Stack, memory *vm.Memory) string {
	offset := stack.Back(0).Uint64()
	length := stack.Back(1).Uint64()
	if memory.Len() < int(offset+length) {
		memory.Resize(offset + length)
	}
	return hex.EncodeToString(memory.Data()[offset : offset+length])
}
func (t *OpsTracer) isPrecompiled(env *vm.EVM, addr common.Address) bool {
	// `isPostMerge` is stolen from internal/ethapi/api.go
	isPostMerge := env.Context.Difficulty.Cmp(common.Big0) == 0
	activePrecompiles := vm.ActivePrecompiles(
		env.ChainConfig().Rules(env.Context.BlockNumber, isPostMerge))

	for _, p := range activePrecompiles {
		if p == addr {
			return true
		}
	}
	return false
}

func (t *OpsTracer) getLabel(topic0 string) string {
	//if op != vm.LOG0 {
	topic0Bs, _ := hex.DecodeString(topic0)
	label, _ := labelDb.Selector(topic0Bs)
	//}
	return label
}

func (t *OpsTracer) CaptureState(pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, rData []byte, depth int, err error) {

	if err != nil {
		t.reason = err
		if t.currentFrame != nil {
			t.currentFrame.Error = err.Error()
		}
		return
	}
	// Fix txs like 0x3494b6a2f62a558c46660691f68e4e2a47694e0b02fad1969e1f0dc725fc9ee5,
	// where a sub-CALL is failed but the whole tx is not reverted.
	if t.currentDepth == depth+1 && (t.currentFrame.Type == vm.CALL.String() ||
		t.currentFrame.Type == vm.CALLCODE.String() ||
		t.currentFrame.Type == vm.DELEGATECALL.String() ||
		t.currentFrame.Type == vm.STATICCALL.String() ||
		t.currentFrame.Type == vm.CREATE.String() ||
		t.currentFrame.Type == vm.CREATE2.String()) {
		t.currentFrame.Error = "Subcall reverted"
		t.currentFrame = t.currentFrame.parent
		t.currentDepth -= 1
	}

	if op == vm.LOG0 || op == vm.LOG1 || op == vm.LOG2 || op == vm.LOG3 || op == vm.LOG4 {
		var topic0, topic1, topic2, topic3, logInput string
		switch op {
		case vm.LOG1:
			topic0 = scope.Stack.Back(2).String()[2:] // remove "0x" prefix
			logInput = topic0
		case vm.LOG2:
			topic0 = scope.Stack.Back(2).String()[2:] // remove "0x" prefix
			topic1 = scope.Stack.Back(3).String()[2:] // remove "0x" prefix
			logInput = strings.Join([]string{topic0, topic1}, " ")
		case vm.LOG3:
			topic0 = scope.Stack.Back(2).String()[2:] // remove "0x" prefix
			topic1 = scope.Stack.Back(3).String()[2:] // remove "0x" prefix
			topic2 = scope.Stack.Back(4).String()[2:] // remove "0x" prefix
			logInput = strings.Join([]string{topic0, topic1, topic2}, " ")
		case vm.LOG4:
			topic0 = scope.Stack.Back(2).String()[2:] // remove "0x" prefix
			topic1 = scope.Stack.Back(3).String()[2:] // remove "0x" prefix
			topic2 = scope.Stack.Back(4).String()[2:] // remove "0x" prefix
			topic3 = scope.Stack.Back(5).String()[2:] // remove "0x" prefix
			logInput = strings.Join([]string{topic0, topic1, topic2, topic3}, " ")
		}
		var label = t.getLabel(topic0)
		frame := OpsCallFrame{
			Type:    op.String(),
			Label:   label,
			From:    strings.ToLower(scope.Contract.Address().String()),
			Input:   logInput,
			Value:   getLogValueHex(scope.Stack, scope.Memory),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			parent:  t.currentFrame,
		}
		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
		return
	}

	switch op {
	case vm.CREATE, vm.CREATE2:
		value := scope.Stack.Back(0)
		from := scope.Contract.Address()
		frame := OpsCallFrame{
			Type:    op.String(),
			From:    strings.ToLower(from.String()),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			Value:   value.String(),
			parent:  t.currentFrame,
		}
		if op == vm.CREATE {
			nonce := t.evm.StateDB.GetNonce(from)
			frame.ContractCreated = crypto.CreateAddress(from, nonce).String()
		}
		if op == vm.CREATE2 {
			frame.salt = scope.Stack.Back(3)
		}
		if !value.IsZero() {
			frame.Label = LabelInternalTransfer
		}
		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
		t.currentFrame = &frame
		t.currentDepth += 1
	case vm.SELFDESTRUCT:
		value := t.evm.StateDB.GetBalance(scope.Contract.Address())
		var to common.Address = scope.Stack.Back(0).Bytes20()
		frame := OpsCallFrame{
			Type:    op.String(),
			From:    strings.ToLower(scope.Contract.Address().String()),
			To:      strings.ToLower(to.String()),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			Value:   value.String(),
			parent:  t.currentFrame,
		}
		if value.Uint64() != 0 {
			frame.Label = LabelInternalTransfer
		}
		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
	case vm.CALL, vm.CALLCODE:
		var to common.Address = scope.Stack.Back(1).Bytes20()
		if t.isPrecompiled(t.evm, to) {
			return
		}
		value := scope.Stack.Back(2)
		frame := OpsCallFrame{
			Type:    op.String(),
			From:    strings.ToLower(scope.Contract.Address().String()),
			To:      strings.ToLower(to.String()),
			Value:   value.String(),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			parent:  t.currentFrame,
		}
		if !value.IsZero() {
			frame.Label = LabelInternalTransfer
		}
		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
		t.currentFrame = &frame
		t.currentDepth += 1
	case vm.DELEGATECALL, vm.STATICCALL:
		var to common.Address = scope.Stack.Back(1).Bytes20()
		if t.isPrecompiled(t.evm, to) {
			return
		}

		frame := OpsCallFrame{
			Type:    op.String(),
			From:    strings.ToLower(scope.Contract.Address().String()),
			To:      strings.ToLower(to.String()),
			GasIn:   uintToHex(gas),
			GasCost: uintToHex(cost),
			parent:  t.currentFrame,
		}

		t.currentFrame.Calls = append(t.currentFrame.Calls, &frame)
		t.currentFrame = &frame
		t.currentDepth += 1
	}
	return

}

func (t *OpsTracer) CaptureEnter(typ vm.OpCode, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {
	if typ == vm.CREATE2 {
		create2Frame := t.currentFrame
		codeHash := crypto.Keccak256Hash(input)
		contractAddr := crypto.CreateAddress2(
			common.HexToAddress(create2Frame.From),
			common.Hash(create2Frame.salt.Bytes32()),
			codeHash.Bytes())
		create2Frame.ContractCreated = contractAddr.String()
	}
	return
}

func (t *OpsTracer) CaptureExit(output []byte, gasUsed uint64, err error) {
	if t.evm.Depth() == t.currentDepth {
		return
	}
	t.currentFrame.GasCost = uintToHex(gasUsed)
	if err != nil {
		t.currentFrame.Error = err.Error()
	}

	t.currentFrame = t.currentFrame.parent
	t.currentDepth -= 1
}

func (t *OpsTracer) CaptureFault(pc uint64, op vm.OpCode, gas, cost uint64, scope *vm.ScopeContext, depth int, err error) {
}

func (t *OpsTracer) CaptureEnd(output []byte, gasUsed uint64, err error) {
	if t.evm.Depth() == t.currentDepth {
		return
	}
	t.currentFrame.GasCost = uintToHex(gasUsed)
	if err != nil {
		t.currentFrame.Error = err.Error()
	}

	t.currentFrame = t.currentFrame.parent
	t.currentDepth -= 1
}

func (t *OpsTracer) GetCallStack() *OpsCallFrame {
	if len(t.callstack.Error) != 0 {
		t.callstack.Calls = []*OpsCallFrame{}
	}
	if t.reason != nil {
		t.callstack.Error = t.reason.Error()
		t.callstack.Calls = []*OpsCallFrame{}
	}
	return &t.callstack
}

func toPbCallTrace(in *OpsCallFrame) *protobuf.StackFrame {
	if in == nil {
		return &protobuf.StackFrame{}
	}

	var calls []*protobuf.StackFrame
	if in.Calls != nil {
		calls = make([]*protobuf.StackFrame, len(in.Calls))
		for i, c := range in.Calls {
			calls[i] = toPbCallTrace(c)
		}
	}
	return &protobuf.StackFrame{
		Type:            in.Type,
		Label:           in.Label,
		From:            in.From,
		To:              in.To,
		ContractCreated: in.ContractCreated,
		Value:           in.Value,
		Input:           in.Input,
		Error:           in.Error,
		Calls:           calls,
	}

}

func dfs(node *protobuf.StackFrame, prefix string, sks *[]types.PlainStackFrame) {
	if node == nil {
		return
	}
	*sks = append(*sks, types.PlainStackFrame{
		FrameId:         prefix,
		Type:            node.Type,
		Label:           node.Label,
		From:            node.From,
		To:              node.To,
		ContractCreated: node.ContractCreated,
		Value:           node.Value,
		Input:           node.Input,
		Error:           node.Error,
		ChildrenCount:   int32(len(node.GetCalls())),
	})
	for i, call := range node.GetCalls() {
		cPrefix := fmt.Sprintf("%s_%d", prefix, i)
		dfs(call, cPrefix, sks)
	}
}

func (t *OpsTracer) GetResult() (json.RawMessage, error) {
	result := t.GetCallStack()
	if t.plain {
		pbTrace := toPbCallTrace(result)
		ptxs := make([]types.PlainStackFrame, 0)
		dfs(pbTrace, "0", &ptxs)
		res, err := json.Marshal(ptxs)
		if err != nil {
			return nil, err
		}
		return res, t.reason
	}

	res, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	return res, t.reason
}

func (t *OpsTracer) Stop(err error) {
	t.reason = err
	atomic.StoreUint32(&t.interrupt, 1)
}

func uintToHex(n uint64) string {
	return "0x" + strconv.FormatUint(n, 16)
}

func addrToHex(a common.Address) string {
	return strings.ToLower(a.Hex())
}
