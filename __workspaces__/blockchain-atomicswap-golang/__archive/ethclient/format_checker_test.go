package ethclient

import (
	"testing"
)

func TestGetInfuraURL(t *testing.T) {
	url := GetInfuraURL("network", "apikeystring")
	if url != "https://network.infura.io/v3/apikeystring" {
		t.Fail()
	}
}

func TestIsValidAddress(t *testing.T) {
	const account = "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"
	if !IsValidAddress(account) {
		t.Fail()
	}
}
