package rpc4go

import (
	"sync"

	"github.com/libs4go/errors"
	"github.com/libs4go/sdi4go"
)

// Client .
type Client interface {
	// BlockNumber get best blocknumber
	BlockNumber() (int64, error)
	BlockByNumber(number int64, block interface{}) error
	BalanceOf(address string, asset string) (string, error)
	SendTransaction(data []byte) (string, error)
}

var injector sdi4go.Injector
var once sync.Once

// Driver .
type Driver interface {
	Create(url string) Client
}

func getInjector() sdi4go.Injector {
	once.Do(func() {
		injector = sdi4go.New()
	})

	return injector
}

func getDriver(name string) (Driver, error) {
	var driver Driver
	if err := getInjector().Create(name, &driver); err != nil {
		return nil, err
	}

	return driver, nil
}

// RegisterDriver .
func RegisterDriver(name string, driver Driver) {
	if err := getInjector().Bind(name, sdi4go.Singleton(driver)); err != nil {
		panic(errors.Wrap(err, "register driver %s error", name))
	}
}

// New create client object with provider driver and url
func New(driverName string, url string) (Client, error) {

	driver, err := getDriver(driverName)

	if err != nil {
		return nil, errors.Wrap(err, "rpc4go driver %s not found", driverName)
	}

	return driver.Create(url), nil
}
