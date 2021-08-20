package downstream

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
)

type Tx struct {
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
}

func FromOrgTx(orgTx *types.Transaction, oBlock *types.Block, signer types.Signer) (tx *Tx, err error) {
	v, r, s := orgTx.RawSignatureValues()
	msg, err := orgTx.AsMessage(signer)
	if err != nil {
		return
	}

	oHeader := oBlock.Header()
	tx = &Tx{
		Hash:        orgTx.Hash().Hex(),
		BlockHash:   oHeader.Hash().Hex(),
		BlockNumber: oHeader.Number,
		From:        strings.ToLower(msg.From().Hex()),
		To:          msg.To(),
		Gas:         msg.Gas(),
		GasPrice:    msg.GasPrice(),
		Input:       hexutil.Encode(msg.Data()),
		Nonce:       msg.Nonce(),
		// Index: ,
		Value:          decimal.NewFromBigInt(msg.Value(), int32(-18)).String(),
		R:              hexutil.EncodeBig(r),
		S:              hexutil.EncodeBig(s),
		V:              hexutil.EncodeBig(v),
		BlockTimestamp: oHeader.Time,
	}

	return
}
