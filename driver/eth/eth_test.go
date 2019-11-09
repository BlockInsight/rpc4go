package eth

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCall(t *testing.T) {
	client := New("https://mainnet.infura.io/v3/44ab06a5fca644df953378ac1c16d2b9")

	receipt, err := client.GetTransactionReceipt("0x0663c3543ee3b91044215aec0fb7d95c2ea9615f1b4d3d6559de85089227b8a8")
	require.NoError(t, err)

	printObject(receipt)
}

func printObject(obj interface{}) {
	buff, _ := json.MarshalIndent(obj, "\t", "\t")

	println(string(buff))
}
