.PHONY: build test

all: build publish

build:
	@GOOS=wasip1 GOARCH=wasm go build -buildmode=c-shared -o enclave.wasm ./plugin/main.go
	@gum log --level info --time kitchen "[ENCLAVE] Completed go build successfully."

test: build
	@mkdir -p ./tmp
	@extism call ./enclave.wasm export --wasi --input="testpwd" >> ./tmp/enclave.json
	@cat ./tmp/enclave.json || true
	@extism call ./enclave.wasm sign --wasi --input ./tmp/enclave.json >> ./tmp/enclave.json

publish: build
	@gum spin --show-error --title "[ENCLAVE] Uploading enclave.wasm to r2" -- sh -c "rclone copy ./enclave.wasm r2:cdn/bin/"
	@gum log --level info --time kitchen "[ENCLAVE] Completed rclone upload successfully."

clean:
	@rm -rf ./tmp
	@rm -rf ./bin
	@rm -rf ./enclave.wasm
	@go mod tidy
