package config

import (
	extism "github.com/extism/go-sdk"
)

func GetManifest() extism.Manifest {
	return extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmUrl{
				Url: "https://github.com/extism/plugins/releases/latest/download/count_vowels.wasm",
			},
		},
		Config: map[string]string{},
	}
}

func GetPluginConfig() extism.PluginConfig {
	return extism.PluginConfig{
		EnableWasi: true,
	}
}
