This is a set of reproduction steps for an issue where roundtrip_js.go fails to handle Promise if RoundTrip context is canceled.

To reproduce:
```
./build.sh
./server
```
Go to [http://localhost:9090/] and open the console. Notice that after a second (when our request times out) we see the following error:
```
wasm_exec.js:536 Uncaught (in promise) Error: Go program has already exited
    at globalThis.Go._resume (wasm_exec.js:536:11)
    at wasm_exec.js:549:8
```