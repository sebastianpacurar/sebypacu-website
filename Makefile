magic:
	echo rm -f web/app.wasm website
	GOOS=js GOARCH=wasm go build -o web/app.wasm
	go build -o website
	./website


