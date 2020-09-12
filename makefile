build:
	@GOARCH=wasm GOOS=js go build -o docs/web/app.wasm ./bin/lofimusic
	@go build -o docs/lofimusic ./bin/lofimusic

run: build
	@cd docs && ./lofimusic local


github: build
	@cd docs && ./lofimusic github

clean:
	@go clean ./...
	@-rm docs/web/app.wasm
	@-rm docs/lofimusic