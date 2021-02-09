From tutorial at https://golangbot.com/webassembly-using-go/

```
cd cmd/wasm
GOOS=js GOARCH=wasm go build -o  ../../assets/primes.wasm
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ../../assets
```

Or use Tinygo

```
cd cmd/wasm
cp $(tinygo env TINYGOROOT)/targets/wasm_exec.js ../../assets
tinygo build -o ../../assets/primes.wasm -target wasm .
```