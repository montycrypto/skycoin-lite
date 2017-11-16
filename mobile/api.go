package mobile

import (
	"github.com/skycoin/skycoin-lite-gopher/liteclient"
	"encoding/json"
)

// Returns a series of addresses based on a seed and the number of addresses
func GetAddresses(seed string, amount int) (string) {
	addresses, _ := liteclient.Addresses(seed, 256)
	response, _ := json.Marshal(addresses)
	return string(response)
}

// Returns addresses with balances, based on an array with balances
func GetBalances(seed string, amount int) (string, error) {
	addresses, err := liteclient.Addresses(seed, amount)
	completeAddresses, err := liteclient.AddressesWithBalance(addresses)
	response, err := json.Marshal(completeAddresses)
	return string(response), err
}

// Returns outputs, based on an array with balances
func GetOutputs(seed string, amount int) (string, error) {
	addresses, err := liteclient.Addresses(seed, amount)
	outputs, err := liteclient.Outputs(addresses)
	response, err := json.Marshal(outputs)
	return string(response), err
}

// Returns a transaction ID
func PostTransaction(seed string, addresses int, destinationAddress string, amount int) (string, error) {
	wallet := liteclient.Wallet{seed, addresses}
	return liteclient.Send(wallet, destinationAddress, uint64(amount))
}
