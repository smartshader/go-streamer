compile:
	protoc api/v1/*.proto \
			--gogo_out=\
	Mgogoproto/gogo.proto=github.com/gogo/protobuf/proto:. \
			--proto_path=\
	$$(go list -f '{{ .Dir }}' -m github.com/gogo/protobuf) \
			--proto_path=.

install_deps:
	go get github.com/gogo/protobuf/proto && \
    go get github.com/gogo/protobuf/jsonpb && \
    go get github.com/gogo/protobuf/protoc-gen-gogo && \
    go get github.com/gogo/protobuf/gogoproto