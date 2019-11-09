package eth

import (
	"github.com/blockinsight/rpc4go"
)

type drvierImpl struct {
}

func (driver *drvierImpl) Create(url string) rpc4go.Client {
	return New(url)
}

func init() {
	rpc4go.RegisterDriver("eth", &drvierImpl{})
}
