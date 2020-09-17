build:
	@GOARCH=wasm GOOS=js go build -o docs/web/app.wasm ./bin/lofimusic
	@go build -o docs/lofimusic ./bin/lofimusic

run: build
	@cd docs && ./lofimusic local


build-github: build
	@cd docs && ./lofimusic github

github: build-github clean 

test:
	go test ./bin/lofimusic
	GOARCH=wasm GOOS=js go test ./bin/lofimusic

clean:
	@go clean ./...
	@-rm docs/lofimusic