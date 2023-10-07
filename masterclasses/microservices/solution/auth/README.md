# Install Homebrew

/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

# Install protoc protobuf compiler

brew install protobuf
brew install protobuf-c

# Install Go gRPC library
go get google.golang.org/grpc

# Install Go protobuf library

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# Compile protobuf file

protoc ./apis/auth.proto --go_out=./apis --go-grpc_out=./apis

# Docker

docker build -t auth .
docker run -it -p 50051:50052 auth



