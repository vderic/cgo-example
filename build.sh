go generate
go clean -cache . 
CGO_CFLAGS_ALLOW='-f.*' go build .

