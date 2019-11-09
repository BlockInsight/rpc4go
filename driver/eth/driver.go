package eth

import (
	"github.com/blockinsight/rpc4go"
)

type drvierImpl struct {
	client Client
}

func (driver *drvierImpl) BlockNumber() (int64, error) {
	return driver.client.BlockNumber()
}
func (driver *drvierImpl) BlockByNumber(number int64, block interface{}) error {
	return driver.client.BlockByNumber(number, block)
}
func (driver *drvierImpl) BalanceOf(address string, asset string) (string, error) {
	return driver.client.BalanceOf(address, asset)
}

func (driver *drvierImpl) SendTransaction(data []byte) (string, error) {
	return driver.client.SendTransaction(data)
}

func init() {
	rpc4go.RegisterDriver("eth", func(url string) rpc4go.Client {
		return &drvierImpl{client: New(url)}
	})
}
