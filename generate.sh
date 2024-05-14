export PATH="$PATH:$(go env GOPATH)/bin"

# Generate proto files for Go
protoc -I proto proto/catalog/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
protoc -I proto proto/sale/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative 

# Generate proto files for TypeScript
npm install grpc-tools grpc_tools_node_protoc_ts
./node_modules/.bin/grpc_tools_node_protoc --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts --ts_out=grpc_js:./gen/ts --proto_path=./proto/sale/ ./proto/sale/quote.proto
./node_modules/.bin/grpc_tools_node_protoc --js_out=import_style=commonjs,binary:./gen/js --grpc_out=grpc_js:./gen/js --proto_path=./proto/sale/ ./proto/sale/quote.proto