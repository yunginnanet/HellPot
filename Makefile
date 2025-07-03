all: deps check build
format :
	find . -iname "*.go" -exec gofmt -s -l -w {} \;
check :
	go vet ./...
run :
	go run cmd/HellPot/*.go
deps:
	go mod tidy -v
test : check
	go test -v ./...
build :
	go build -x -trimpath -ldflags "-s -w -X main.version=`git tag --sort=-version:refname | head -n 1`" cmd/HellPot/*.go
