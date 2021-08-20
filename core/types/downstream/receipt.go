package downstream

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type Receipt struct {
	TransactionHash   string      `json:"transactionHash"`
	TransactionIndex  int         `json:"transactionIndex"`
	BlockHash         string      `json:"blockHash"`
	BlockNumber       *big.Int    `json:"blockNumber"`
	CumulativeGasUsed string      `json:"cumulativeGasUsed"`
	GasUsed           uint64      `json:"gasUsed"`
	Address           string      `json:"contractAddress"`
	Bloom             string      `json:"logsBloom"`
	Root              *string     `json:"root,omitempty"`
	Status            *bool       `json:"transactionReceiptStatus,omitempty"`
	EventCount        int         `json:"eventCount"`
	BlockTimestamp    uint64      `json:"blockTimestamp"`
	Logs              []*EventLog `json:"eventLogs"`
}

type EventLog struct {
	Id               *big.Int     `json:"_id"`
	BlockNumber      *big.Int     `json:"blockNumber"`
	LogIndex         uint         `json:"logIndex"`
	Address          string       `json:"address"`
	Topic0           *common.Hash `json:"topic0,omitempty"`
	Topic1           *common.Hash `json:"topic1,omitempty"`
	Topic2           *common.Hash `json:"topic2,omitempty"`
	Topic3           *common.Hash `json:"topic3,omitempty"`
	Data             string       `json:"eventData"`
	TransactionIndex int          `json:"transactionIndex"`
	TransactionHash  string       `json:"transactionHash"`
	BlockHash        string       `json:"blockHash"`
	Removed          bool         `json:"removed"`
	BlockTimestamp   uint64       `json:"blockTimestamp"`
	Class            string       `json:"_class,omitempty"`
}

func FromOrgReceipt(oR *types.Receipt, transactionHash string, transactionIndex int,
	oBlock *types.Block) *Receipt {

	oHeader := oBlock.Header()
	result := &Receipt{
		TransactionHash:   transactionHash,
		TransactionIndex:  transactionIndex,
		BlockHash:         oHeader.Hash().Hex(),
		BlockNumber:       oHeader.Number,
		CumulativeGasUsed: fmt.Sprintf("0x%x", oR.CumulativeGasUsed),
		GasUsed:           oR.GasUsed,
		Address:           strings.ToLower(oR.ContractAddress.Hex()),
		Bloom:             hexutil.Encode(hexutil.Bytes(oR.Bloom[:])),
		EventCount:        len(oR.Logs),
		BlockTimestamp:    oHeader.Time,
		Logs:              make([]*EventLog, 0, len(oR.Logs)),
	}

	if len(oR.PostState) > 0 {
		root := hexutil.Encode(oR.PostState)
		result.Root = &root
	} else {
		stat := oR.Status == 1
		result.Status = &stat
	}

	for _, oLog := range oR.Logs {
		id := big.NewInt(0).Mul(oHeader.Number, big.NewInt(10000000))
		id.Add(id, new(big.Int).SetUint64(uint64(oLog.Index)))
		log := &EventLog{
			Id:               id,
			BlockNumber:      oHeader.Number,
			LogIndex:         oLog.Index,
			Address:          strings.ToLower(oLog.Address.Hex()),
			Data:             hexutil.Encode(oLog.Data),
			TransactionHash:  transactionHash,
			TransactionIndex: transactionIndex,
			BlockHash:        oHeader.Hash().Hex(),
			Removed:          oLog.Removed,
			BlockTimestamp:   oHeader.Time,
		}

		if len(oLog.Topics) > 0 {
			log.Topic0 = &oLog.Topics[0]
		}
		if len(oLog.Topics) > 1 {
			log.Topic1 = &oLog.Topics[1]
		}
		if len(oLog.Topics) > 2 {
			log.Topic2 = &oLog.Topics[2]
		}
		if len(oLog.Topics) > 3 {
			log.Topic3 = &oLog.Topics[3]
		}

		result.Logs = append(result.Logs, log)
	}

	return result
}
