package runtime

import (
	"errors"

	extism "github.com/extism/go-sdk"
)

type EnclaveHost interface {
	Generate() ([]byte, error)
	Unlock() ([]byte, error)
	Sign(message []byte) ([]byte, error)
	Verify(pubKey []byte, message []byte, sig []byte) (bool, error)
}

//	func NewEnclaveHost() EnclaveHost {
//		return EnclaveHost{
//			Manifest: extism.Manifest{
//				Wasm: []extism.Wasm{
//					extism.WasmUrl{
//						Url: "https://github.com/extism/plugins/releases/latest/download/count_vowels.wasm",
//					},
//				},
//				Config: map[string]string{
//					config.KeyChainID: "aeiouyAEIOUY",
//				},
//			},
//		}
//	}
type enclaveHost struct {
	plugin extism.Plugin
}

func (h *enclaveHost) Generate() ([]byte, error) {
	exit, out, err := h.plugin.Call("generate", nil)
	if err != nil {
		return nil, err
	}
	if exit != 0 {
		return nil, errors.New(string(out))
	}
	return out, nil
}

func (h *enclaveHost) Unlock() ([]byte, error) {
	exit, out, err := h.plugin.Call("unlock", nil)
	if err != nil {
		return nil, err
	}
	if exit != 0 {
		return nil, errors.New(string(out))
	}
	return out, nil
}

func (h *enclaveHost) Sign(message []byte) ([]byte, error) {
	exit, out, err := h.plugin.Call("sign", message)
	if err != nil {
		return nil, err
	}
	if exit != 0 {
		return nil, errors.New(string(out))
	}
	return out, nil
}

func (h *enclaveHost) Verify(pubKey []byte, message []byte, sig []byte) (bool, error) {
	exit, out, err := h.plugin.Call("verify", pubKey)
	if err != nil {
		return false, err
	}
	if exit != 0 {
		return false, errors.New(string(out))
	}
	return true, nil
}
