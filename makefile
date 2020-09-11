build:
	@GOARCH=wasm GOOS=js go build -o web/app.wasm
	@go build

run: build
	@./lofimusic

clean:
	@go clean
	@-rm web/app.wasm