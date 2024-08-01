.PHONY: build
build:
	go build

.PHONY: release
release:
	@rm -rf ./dist
	@mkdir -p ./dist
	GOOS=darwin GOARCH=amd64 go build -o ./dist/git-branches_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build -o ./dist/git-branches_darwin_arm64
