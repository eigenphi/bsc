// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: Block.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StackFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type            string        `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`   // OP code
	Label           string        `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"` // empty or "Transfer" or "Internal-Transfer"
	From            string        `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To              string        `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	ContractCreated string        `protobuf:"bytes,5,opt,name=contractCreated,proto3" json:"contractCreated,omitempty"`
	Value           string        `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	Input           string        `protobuf:"bytes,7,opt,name=input,proto3" json:"input,omitempty"`
	Error           string        `protobuf:"bytes,8,opt,name=error,proto3" json:"error,omitempty"`
	Calls           []*StackFrame `protobuf:"bytes,9,rep,name=calls,proto3" json:"calls,omitempty"`
}

func (x *StackFrame) Reset() {
	*x = StackFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StackFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StackFrame) ProtoMessage() {}

func (x *StackFrame) ProtoReflect() protoreflect.Message {
	mi := &file_Block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StackFrame.ProtoReflect.Descriptor instead.
func (*StackFrame) Descriptor() ([]byte, []int) {
	return file_Block_proto_rawDescGZIP(), []int{0}
}

func (x *StackFrame) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *StackFrame) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *StackFrame) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *StackFrame) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *StackFrame) GetContractCreated() string {
	if x != nil {
		return x.ContractCreated
	}
	return ""
}

func (x *StackFrame) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *StackFrame) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *StackFrame) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *StackFrame) GetCalls() []*StackFrame {
	if x != nil {
		return x.Calls
	}
	return nil
}

type InternalTxn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionHash string `protobuf:"bytes,1,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	From            string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To              string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Value           string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *InternalTxn) Reset() {
	*x = InternalTxn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InternalTxn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InternalTxn) ProtoMessage() {}

func (x *InternalTxn) ProtoReflect() protoreflect.Message {
	mi := &file_Block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InternalTxn.ProtoReflect.Descriptor instead.
func (*InternalTxn) Descriptor() ([]byte, []int) {
	return file_Block_proto_rawDescGZIP(), []int{1}
}

func (x *InternalTxn) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *InternalTxn) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *InternalTxn) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *InternalTxn) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type EventLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain            string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	Address          string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Topic0           string `protobuf:"bytes,3,opt,name=topic0,proto3" json:"topic0,omitempty"`
	Topic1           string `protobuf:"bytes,4,opt,name=topic1,proto3" json:"topic1,omitempty"`
	Topic2           string `protobuf:"bytes,5,opt,name=topic2,proto3" json:"topic2,omitempty"`
	Topic3           string `protobuf:"bytes,6,opt,name=topic3,proto3" json:"topic3,omitempty"`
	Data             string `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
	BlockNumber      int64  `protobuf:"varint,8,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	BlockTimestamp   int64  `protobuf:"varint,9,opt,name=blockTimestamp,proto3" json:"blockTimestamp,omitempty"`
	TransactionHash  string `protobuf:"bytes,10,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	TransactionIndex int32  `protobuf:"varint,11,opt,name=transactionIndex,proto3" json:"transactionIndex,omitempty"`
	BlockHash        string `protobuf:"bytes,12,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	LogIndex         int32  `protobuf:"varint,13,opt,name=logIndex,proto3" json:"logIndex,omitempty"`
	Removed          bool   `protobuf:"varint,14,opt,name=removed,proto3" json:"removed,omitempty"`
	SenderInfo       string `protobuf:"bytes,15,opt,name=senderInfo,proto3" json:"senderInfo,omitempty"`
}

func (x *EventLog) Reset() {
	*x = EventLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventLog) ProtoMessage() {}

func (x *EventLog) ProtoReflect() protoreflect.Message {
	mi := &file_Block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventLog.ProtoReflect.Descriptor instead.
func (*EventLog) Descriptor() ([]byte, []int) {
	return file_Block_proto_rawDescGZIP(), []int{2}
}

func (x *EventLog) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *EventLog) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *EventLog) GetTopic0() string {
	if x != nil {
		return x.Topic0
	}
	return ""
}

func (x *EventLog) GetTopic1() string {
	if x != nil {
		return x.Topic1
	}
	return ""
}

func (x *EventLog) GetTopic2() string {
	if x != nil {
		return x.Topic2
	}
	return ""
}

func (x *EventLog) GetTopic3() string {
	if x != nil {
		return x.Topic3
	}
	return ""
}

func (x *EventLog) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *EventLog) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *EventLog) GetBlockTimestamp() int64 {
	if x != nil {
		return x.BlockTimestamp
	}
	return 0
}

func (x *EventLog) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *EventLog) GetTransactionIndex() int32 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

func (x *EventLog) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *EventLog) GetLogIndex() int32 {
	if x != nil {
		return x.LogIndex
	}
	return 0
}

func (x *EventLog) GetRemoved() bool {
	if x != nil {
		return x.Removed
	}
	return false
}

func (x *EventLog) GetSenderInfo() string {
	if x != nil {
		return x.SenderInfo
	}
	return ""
}

type Receipt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain                    string      `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	TransactionHash          string      `protobuf:"bytes,2,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	TransactionIndex         int32       `protobuf:"varint,3,opt,name=transactionIndex,proto3" json:"transactionIndex,omitempty"`
	BlockHash                string      `protobuf:"bytes,4,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	BlockNumber              int64       `protobuf:"varint,5,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	GasUsed                  int64       `protobuf:"varint,6,opt,name=gasUsed,proto3" json:"gasUsed,omitempty"`
	ContractAddress          string      `protobuf:"bytes,7,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	TransactionReceiptStatus bool        `protobuf:"varint,8,opt,name=transactionReceiptStatus,proto3" json:"transactionReceiptStatus,omitempty"`
	EventCount               int32       `protobuf:"varint,9,opt,name=eventCount,proto3" json:"eventCount,omitempty"`
	BlockTimestamp           int64       `protobuf:"varint,10,opt,name=blockTimestamp,proto3" json:"blockTimestamp,omitempty"`
	EventLogs                []*EventLog `protobuf:"bytes,11,rep,name=eventLogs,proto3" json:"eventLogs,omitempty"`
	SenderInfo               string      `protobuf:"bytes,12,opt,name=senderInfo,proto3" json:"senderInfo,omitempty"`
}

func (x *Receipt) Reset() {
	*x = Receipt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Receipt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Receipt) ProtoMessage() {}

func (x *Receipt) ProtoReflect() protoreflect.Message {
	mi := &file_Block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Receipt.ProtoReflect.Descriptor instead.
func (*Receipt) Descriptor() ([]byte, []int) {
	return file_Block_proto_rawDescGZIP(), []int{3}
}

func (x *Receipt) GetChain() string {
	if x != nil {
		return x.Chain
	}
	return ""
}

func (x *Receipt) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *Receipt) GetTransactionIndex() int32 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

func (x *Receipt) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *Receipt) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *Receipt) GetGasUsed() int64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *Receipt) GetContractAddress() string {
	if x != nil {
		return x.ContractAddress
	}
	return ""
}

func (x *Receipt) GetTransactionReceiptStatus() bool {
	if x != nil {
		return x.TransactionReceiptStatus
	}
	return false
}

func (x *Receipt) GetEventCount() int32 {
	if x != nil {
		return x.EventCount
	}
	return 0
}

func (x *Receipt) GetBlockTimestamp() int64 {
	if x != nil {
		return x.BlockTimestamp
	}
	return 0
}

func (x *Receipt) GetEventLogs() []*EventLog {
	if x != nil {
		return x.EventLogs
	}
	return nil
}

func (x *Receipt) GetSenderInfo() string {
	if x != nil {
		return x.SenderInfo
	}
	return ""
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionHash  string `protobuf:"bytes,1,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	TransactionIndex int32  `protobuf:"varint,2,opt,name=transactionIndex,proto3" json:"transactionIndex,omitempty"`
	BlockNumber      int64  `protobuf:"varint,3,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	FromAddress      string `protobuf:"bytes,4,opt,name=fromAddress,proto3" json:"fromAddress,omitempty"`
	ToAddress        string `protobuf:"bytes,5,opt,name=toAddress,proto3" json:"toAddress,omitempty"`
	Nonce            int64  `protobuf:"varint,6,opt,name=nonce,proto3" json:"nonce,omitempty"`
	TransactionValue string `protobuf:"bytes,7,opt,name=transactionValue,proto3" json:"transactionValue,omitempty"`
	BlockTimestamp   int64  `protobuf:"varint,8,opt,name=blockTimestamp,proto3" json:"blockTimestamp,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_Block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_Block_proto_rawDescGZIP(), []int{4}
}

func (x *Transaction) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *Transaction) GetTransactionIndex() int32 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

func (x *Transaction) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *Transaction) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *Transaction) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *Transaction) GetNonce() int64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *Transaction) GetTransactionValue() string {
	if x != nil {
		return x.TransactionValue
	}
	return ""
}

func (x *Transaction) GetBlockTimestamp() int64 {
	if x != nil {
		return x.BlockTimestamp
	}
	return 0
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber    int64  `protobuf:"varint,2,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	BlockHash      string `protobuf:"bytes,3,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	ParentHash     string `protobuf:"bytes,4,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	Miner          string `protobuf:"bytes,6,opt,name=miner,proto3" json:"miner,omitempty"`
	BlockSize      int32  `protobuf:"varint,7,opt,name=blockSize,proto3" json:"blockSize,omitempty"`
	GasLimit       int64  `protobuf:"varint,8,opt,name=gasLimit,proto3" json:"gasLimit,omitempty"`
	GasUsed        int64  `protobuf:"varint,9,opt,name=gasUsed,proto3" json:"gasUsed,omitempty"`
	BlockTimestamp int64  `protobuf:"varint,10,opt,name=blockTimestamp,proto3" json:"blockTimestamp,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Block_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_Block_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_Block_proto_rawDescGZIP(), []int{5}
}

func (x *Block) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *Block) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *Block) GetParentHash() string {
	if x != nil {
		return x.ParentHash
	}
	return ""
}

func (x *Block) GetMiner() string {
	if x != nil {
		return x.Miner
	}
	return ""
}

func (x *Block) GetBlockSize() int32 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

func (x *Block) GetGasLimit() int64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *Block) GetGasUsed() int64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *Block) GetBlockTimestamp() int64 {
	if x != nil {
		return x.BlockTimestamp
	}
	return 0
}

type TraceTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber      int64       `protobuf:"varint,1,opt,name=blockNumber,proto3" json:"blockNumber,omitempty"`
	TransactionHash  string      `protobuf:"bytes,2,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	TransactionIndex int32       `protobuf:"varint,3,opt,name=transactionIndex,proto3" json:"transactionIndex,omitempty"`
	FromAddress      string      `protobuf:"bytes,4,opt,name=fromAddress,proto3" json:"fromAddress,omitempty"`
	ToAddress        string      `protobuf:"bytes,5,opt,name=toAddress,proto3" json:"toAddress,omitempty"`
	GasPrice         int64       `protobuf:"varint,6,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"`
	Input            string      `protobuf:"bytes,7,opt,name=input,proto3" json:"input,omitempty"`
	Nonce            int64       `protobuf:"varint,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
	TransactionValue string      `protobuf:"bytes,9,opt,name=transactionValue,proto3" json:"transactionValue,omitempty"`
	Stack            *StackFrame `protobuf:"bytes,10,opt,name=stack,proto3" json:"stack,omitempty"`
	BlockTimestamp   int64       `protobuf:"varint,11,opt,name=blockTimestamp,proto3" json:"blockTimestamp,omitempty"`
}

func (x *TraceTransaction) Reset() {
	*x = TraceTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Block_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceTransaction) ProtoMessage() {}

func (x *TraceTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_Block_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceTransaction.ProtoReflect.Descriptor instead.
func (*TraceTransaction) Descriptor() ([]byte, []int) {
	return file_Block_proto_rawDescGZIP(), []int{6}
}

func (x *TraceTransaction) GetBlockNumber() int64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

func (x *TraceTransaction) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *TraceTransaction) GetTransactionIndex() int32 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

func (x *TraceTransaction) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *TraceTransaction) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *TraceTransaction) GetGasPrice() int64 {
	if x != nil {
		return x.GasPrice
	}
	return 0
}

func (x *TraceTransaction) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *TraceTransaction) GetNonce() int64 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *TraceTransaction) GetTransactionValue() string {
	if x != nil {
		return x.TransactionValue
	}
	return ""
}

func (x *TraceTransaction) GetStack() *StackFrame {
	if x != nil {
		return x.Stack
	}
	return nil
}

func (x *TraceTransaction) GetBlockTimestamp() int64 {
	if x != nil {
		return x.BlockTimestamp
	}
	return 0
}

var File_Block_proto protoreflect.FileDescriptor

var file_Block_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x65,
	0x72, 0x69, 0x67, 0x6f, 0x6e, 0x22, 0xf0, 0x01, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x28,
	0x0a, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x65, 0x72, 0x69, 0x67, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x52, 0x05, 0x63, 0x61, 0x6c, 0x6c, 0x73, 0x22, 0x71, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x54, 0x78, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xc2, 0x03, 0x0a, 0x08,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x30, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x30,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x32,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x33, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x33, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26,
	0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x6f,
	0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0xcd, 0x03, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2a, 0x0a, 0x10,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x61, 0x73, 0x55,
	0x73, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67, 0x61, 0x73, 0x55, 0x73,
	0x65, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x3a, 0x0a, 0x18,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x2e, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x0b, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x72, 0x69, 0x67, 0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0xaf, 0x02, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x72, 0x6f, 0x6d,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66,
	0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x2a,
	0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x22, 0xf9, 0x01, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x20, 0x0a, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05,
	0x6d, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x69, 0x6e,
	0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x67,
	0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x90,
	0x03, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x63, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x66,
	0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x67,
	0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x67,
	0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6e, 0x6f,
	0x6e, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x28, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x65, 0x72, 0x69, 0x67, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x46, 0x72, 0x61,
	0x6d, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x42, 0x37, 0x0a, 0x19, 0x69, 0x6f, 0x2e, 0x65, 0x69, 0x67, 0x65, 0x6e, 0x70, 0x68, 0x69,
	0x2e, 0x65, 0x72, 0x69, 0x67, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x42, 0x0c,
	0x45, 0x72, 0x69, 0x67, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x0a,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_Block_proto_rawDescOnce sync.Once
	file_Block_proto_rawDescData = file_Block_proto_rawDesc
)

func file_Block_proto_rawDescGZIP() []byte {
	file_Block_proto_rawDescOnce.Do(func() {
		file_Block_proto_rawDescData = protoimpl.X.CompressGZIP(file_Block_proto_rawDescData)
	})
	return file_Block_proto_rawDescData
}

var file_Block_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_Block_proto_goTypes = []interface{}{
	(*StackFrame)(nil),       // 0: erigon.stackFrame
	(*InternalTxn)(nil),      // 1: erigon.InternalTxn
	(*EventLog)(nil),         // 2: erigon.EventLog
	(*Receipt)(nil),          // 3: erigon.Receipt
	(*Transaction)(nil),      // 4: erigon.Transaction
	(*Block)(nil),            // 5: erigon.Block
	(*TraceTransaction)(nil), // 6: erigon.TraceTransaction
}
var file_Block_proto_depIdxs = []int32{
	0, // 0: erigon.stackFrame.calls:type_name -> erigon.stackFrame
	2, // 1: erigon.Receipt.eventLogs:type_name -> erigon.EventLog
	0, // 2: erigon.TraceTransaction.stack:type_name -> erigon.stackFrame
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Block_proto_init() }
func file_Block_proto_init() {
	if File_Block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StackFrame); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InternalTxn); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventLog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Receipt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Block_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Block_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Block_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Block_proto_goTypes,
		DependencyIndexes: file_Block_proto_depIdxs,
		MessageInfos:      file_Block_proto_msgTypes,
	}.Build()
	File_Block_proto = out.File
	file_Block_proto_rawDesc = nil
	file_Block_proto_goTypes = nil
	file_Block_proto_depIdxs = nil
}
