package eth

import (
	"encoding/hex"
	"fmt"

	"github.com/blockinsight/abi4go/eth/erc20"
	"github.com/blockinsight/rpc4go/driver/internal/jsonrpc"
	"github.com/libs4go/errors"
	"github.com/libs4go/fixed"
	"github.com/libs4go/slf4go"
)

// Client .
type Client interface {
	BlockNumber() (int64, error)
	BlockByNumber(number int64, block interface{}) error
	BalanceOf(address string, asset string) (string, error)
	SendTransaction(data []byte) (string, error)
	GetTransactionReceipt(tx string) (val *TransactionReceipt, err error)
}

type clientImpl struct {
	*jsonrpc.RPCClient
	slf4go.Logger
}

// New create eth rpc client with remote node url
func New(url string) Client {
	return &clientImpl{
		RPCClient: jsonrpc.NewRPCClient(url),
		Logger:    slf4go.Get("ethrpc-client"),
	}
}

// GetTransactionReceipt ...
func (client *clientImpl) GetTransactionReceipt(tx string) (val *TransactionReceipt, err error) {

	err = client.Call2("eth_getTransactionReceipt", &val, tx)

	return
}

func (client *clientImpl) BlockNumber() (int64, error) {
	var data string

	err := client.Call2("eth_blockNumber", &data)

	if err != nil {
		return 0, err
	}

	val, err := fixed.New(0, fixed.HexRawValue(data))

	if err != nil {
		return 0, errors.Wrap(err, "decode %s error", data)
	}

	return val.RawValue.Int64(), nil
}
func (client *clientImpl) BlockByNumber(number int64, val interface{}) (err error) {
	err = client.Call2("eth_getBlockByNumber", val, fmt.Sprintf("0x%x", number), true)

	return
}

func (client *clientImpl) Call(callsite *CallSite) (val string, err error) {

	err = client.Call2("eth_call", &val, callsite, "latest")

	return
}

func (client *clientImpl) BalanceOf(address string, asset string) (string, error) {

	if asset == "" || asset == "0x0000000000000000000000000000000000000000" {
		data := erc20.BalanceOf(address)

		valstr, err := client.Call(&CallSite{
			To:   asset,
			Data: data,
		})

		if err != nil {
			return "", errors.Wrap(err, "eth_call error")
		}

		return valstr, nil

	}

	var data string

	err := client.Call2("eth_getBalance", &data, address, "latest")

	if err != nil {
		return "", err
	}

	return data, nil

}

func (client *clientImpl) SendTransaction(data []byte) (val string, err error) {
	err = client.Call2("eth_sendRawTransaction", &val, "0x"+hex.EncodeToString(data))

	return
}
