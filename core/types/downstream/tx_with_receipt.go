package downstream

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/shopspring/decimal"
)

type TxWithReceipt struct {
	Hash           string          `json:"_id"`
	BlockHash      string          `json:"blockHash"`
	BlockNumber    *big.Int        `json:"blockNumber"`
	From           string          `json:"fromAddress"`
	Gas            uint64          `json:"gas",string`
	GasPrice       *big.Int        `json:"gasPrice"`
	Input          string          `json:"input"`
	Nonce          uint64          `json:"nonce"`
	To             *common.Address `json:"toAddress"`
	Index          int             `json:"transactionIndex"`
	Value          string          `json:"transactionValue"`
	V              string          `json:"v"`
	R              string          `json:"r"`
	S              string          `json:"s"`
	BlockTimestamp uint64          `json:"blockTimestamp"`
	Receipt        *Receipt        `json:"receipt"`
	Class          string          `json:"_class"`
}

func FromOrgTxWithReceipt(orgTx *types.Transaction, oBlock *types.Block, receipt *Receipt, signer types.Signer) (tx *TxWithReceipt, err error) {
	v, r, s := orgTx.RawSignatureValues()
	msg, err := orgTx.AsMessage(signer)
	if err != nil {
		return
	}

	tx = &TxWithReceipt{
		Hash:           orgTx.Hash().Hex(),
		BlockHash:      oBlock.Header().Hash().Hex(),
		BlockNumber:    oBlock.Header().Number,
		From:           strings.ToLower(msg.From().Hex()),
		To:             msg.To(),
		Gas:            msg.Gas(),
		GasPrice:       msg.GasPrice(),
		Input:          hexutil.Encode(msg.Data()),
		Nonce:          msg.Nonce(),
		Index:          receipt.TransactionIndex,
		Value:          decimal.NewFromBigInt(msg.Value(), int32(-18)).String(),
		R:              hexutil.EncodeBig(r),
		S:              hexutil.EncodeBig(s),
		V:              hexutil.EncodeBig(v),
		BlockTimestamp: oBlock.Header().Time,
		Receipt:        receipt,
		Class:          "com.mingsi.data.connector.entity.Transaction",
	}

	return
}
