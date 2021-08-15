magic:
	rm -f web/app.wasm pwa
	GOOS=js GOARCH=wasm go build -o web/app.wasm
	go build -o pwa
	./pwa


