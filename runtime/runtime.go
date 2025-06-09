package runtime

import (
	"context"
	"enclave/config"
	"errors"

	extism "github.com/extism/go-sdk"
)

type EnclaveHost interface {
	Generate() ([]byte, error)
	Unlock() ([]byte, error)
	Sign(message []byte) ([]byte, error)
	Verify(pubKey []byte, message []byte, sig []byte) (bool, error)
}

type enclaveHost struct {
	plugin *extism.Plugin
}

func NewEnclaveHost(ctx context.Context) EnclaveHost {
	h := &enclaveHost{}
	manifest := config.GetManifest()
	cfg := extism.PluginConfig{}
	plugin, err := extism.NewPlugin(ctx, manifest, cfg, []extism.HostFunction{})
	if err != nil {
		panic(err)
	}
	h.plugin = plugin
	return h
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
