all: lint gen build
with-webapp: gen lint bf build  

lint:
	@buf lint

gen: clean
	@buf generate

clean:
	@rm -rf pkg/api pkg/webapp/dist/*
	@rm -f bin/*

build:
	@go mod tidy
	@cd cmd/server && go build -o ../../bin/ic_server

# Build the frontend Vue App for production and create the dist directory.
# Copy the dist directory to pkg/webapp to compile the dist directory 
# into the ic-server binary. 
bf: build-frontend
build-frontend:
	@echo "*** $@"
	@cd ../frontend && yarn && yarn build
	@cp -r ../frontend/dist pkg/webapp

init:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest
