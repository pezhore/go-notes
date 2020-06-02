test:
	go test -v $(shell go list ./... | grep -v /vendor/) 


build: deps
	gox -osarch="linux/amd64 windows/amd64 darwin/amd64" \
	-output="pkg/{{.OS}}_{{.Arch}}/gn" .

release: release_bump release_build

release_bump:
	scripts/release_bump.sh

release_build:
	scripts/release_build.sh

deps:

clean:
	rm -rf pkg/
