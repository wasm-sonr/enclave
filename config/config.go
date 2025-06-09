package config

import (
	"errors"

	"github.com/extism/go-pdk"
	"github.com/go-sonr/crypto/mpc"
)

const (
	KeyChainID  = "chain_id"
	KeyPassword = "password"
	KeyGateway  = "gateway"
	KeyEnclave  = "enclave"
	KeyPubKey   = "pub_key"
)

// GetChainID returns the chain ID to use for unlocking the enclave
func GetChainID() string {
	v := pdk.GetVar(KeyChainID)
	if v == nil {
		return "sonr-testnet-1"
	}
	return string(v)
}

// GetPassword returns the key to use for unlocking the enclave
func GetPassword() []byte {
	v := pdk.GetVar(KeyPassword)
	if v == nil {
		return []byte("password")
	}
	return v
}

// GetGateway returns the gateway URL without the cid value
func GetGateway() string {
	v := pdk.GetVar(KeyGateway)
	if v == nil {
		return "https://ipfs.did.run/ipfs/"
	}
	return string(v)
}

// GetEnclave returns the Enclave interface
func GetEnclave() (mpc.Enclave, error) {
	v := pdk.GetVar(KeyEnclave)
	if v == nil {
		return nil, errors.New("enclave not found")
	}
	e, err := mpc.ImportEnclave(mpc.WithEncryptedData(v, GetPassword()))
	if err != nil {
		pdk.SetError(err)
		return nil, err
	}
	pdk.Log(pdk.LogInfo, "Imported enclave successfully")
	return e, nil
}
