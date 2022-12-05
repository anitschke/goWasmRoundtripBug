SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

cd $SCRIPT_DIR/cmd/wasm/  
GOOS=js GOARCH=wasm go build -o  $SCRIPT_DIR/assets/exitBeforeFetchFinished.wasm  
cp -f "$(go env GOROOT)/misc/wasm/wasm_exec.js" $SCRIPT_DIR/assets/

cd $SCRIPT_DIR/cmd/server
go build -o $SCRIPT_DIR/server