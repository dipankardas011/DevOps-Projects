GOOS=js
GOARCH=wasm

main:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $@.wasm
	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
# install goexec: go get -u github.com/shurcooL/goexec
	goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'