# Enclave

This is a simple enclave that can be used to generate and sign data.

## Usage

To use the enclave, you can either run it inside the Cloudflare Durable Object or use it using the Golang Host Runtime.

### Cloudflare Durable Object

The Cloudflare Durable Object is a simple way to run the enclave in the cloud. You can find the code for the Durable Object under the `worker` directory.

### Golang Host Runtime

If you want to run the enclave locally, you can use the Golang Host Runtime. You can find the code for the Host Runtime under the `runtime` directory.

- The `runtime` directory leverages the extism sdk and also proto actor in order to maintain plugin specific functionality.

## API

### `generate()`

Generates a new key pair and returns the public key.

### `sign(data)`

Signs the provided data using the private key.

### `verify(data, signature)`

Verifies the provided signature using the public key.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
