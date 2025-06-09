package main

import (
	"github.com/extism/go-pdk"
	"github.com/go-sonr/crypto/mpc"
)

type SignRequest struct {
	Message []byte `json:"message"`
	Enclave []byte `json:"enclave"`
}

type SignResponse struct {
	Signature []byte `json:"signature"`
}

type VerifyRequest struct {
	PubKey  []byte `json:"pub_key"`
	Message []byte `json:"message"`
	Sig     []byte `json:"sig"`
}

type VerifyResponse struct {
	Valid bool `json:"valid"`
}

func main() {
	export()
	unlock()
	sign()
	verify()
}

//go:wasmexport export
func export() int32 {
	pwdInput := pdk.InputString()
	e, err := mpc.NewEnclave()
	if err != nil {
		pdk.SetError(err)
		return 1
	}
	pdk.SetVar("PUB_KEY", e.PubKeyBytes())
	pdk.Log(pdk.LogInfo, "Generated enclave successfully")
	bz, err := e.Encrypt([]byte(pwdInput))
	if err != nil {
		pdk.SetError(err)
		return 1
	}
	pdk.SetVar("ENC_DATA", bz)
	pdk.SetVarInt("UNLOCKED", 0)
	pdk.Output(bz)
	return 0
}

//go:wasmexport unlock
func unlock() int32 {
	return 0
}

//go:wasmexport sign
func sign() int32 {
	req := SignRequest{}
	err := pdk.InputJSON(req)
	if err != nil {
		pdk.SetError(err)
		return 1
	}
	pdk.Log(pdk.LogInfo, "Deserialized request successfully")
	e, err := mpc.ImportEnclave(mpc.WithEnclaveJSON(req.Enclave))
	if err != nil {
		pdk.SetError(err)
		return 1
	}
	pdk.Log(pdk.LogInfo, "Imported enclave successfully")
	sig, err := e.Sign(req.Message)
	if err != nil {
		pdk.SetError(err)
		return 1
	}
	pdk.Log(pdk.LogInfo, "Signature successful")
	sigJSON := SignResponse{Signature: sig}
	pdk.OutputJSON(sigJSON)
	return 0
}

//go:wasmexport verify
func verify() int32 {
	req := VerifyRequest{}
	err := pdk.InputJSON(req)
	if err != nil {
		pdk.Log(pdk.LogError, err.Error())
		return 1
	}
	pdk.Log(pdk.LogInfo, "Deserialized request successfully")
	res := VerifyResponse{Valid: false}
	valid, err := mpc.VerifyWithPubKey(req.PubKey, req.Message, req.Sig)
	if err != nil {
		pdk.SetError(err)
		return 1
	}
	res.Valid = valid
	pdk.OutputJSON(res)
	return 0
}
