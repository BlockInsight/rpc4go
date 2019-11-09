package eth

import (
	"github.com/blockinsight/model.proto/golang/eth"
)

// Block .
type Block eth.Blockchain

// Transaction .
type Transaction eth.Transaction

// TransactionReceipt .
type TransactionReceipt eth.TransactionReceipt

// CallSite .
type CallSite struct {
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Value    string `json:"value,omitempty"`
	GasPrice string `json:"gasPrice,omitempty"`
	Gas      string `json:"gas,omitempty"`
	Data     string `json:"data,omitempty"`
}
