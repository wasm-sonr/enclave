# Enclave

This is a simple enclave that can be used to generate and sign data.

## Usage

To use the enclave, you can run the following command:

```bash
extism call ./enclave.wasm generate --wasi
```

This will generate a new enclave and output the public key and the serialized data of the enclave.

To sign data with the enclave, you can run the following command:

```bash
extism call ./enclave.wasm sign --wasi --input ./tmp/enclave.json
```

This will sign the data using the enclave and output the signature.
