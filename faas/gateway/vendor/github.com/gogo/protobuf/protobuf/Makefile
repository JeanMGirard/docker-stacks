URL="https://raw.githubusercontent.com/google/protobuf/master/src/google/protobuf/"

update:
	go install github.com/gogo/protobuf/gogoreplace

	(cd ./google/protobuf && rm descriptor.proto; wget ${URL}/descriptor.proto)
	# gogoprotobuf requires users to import gogo.proto which imports descriptor.proto
	# The descriptor.proto is only compatible with proto3 just because of the reserved keyword.
	# We remove it to stay compatible with previous versions of protoc before proto3
	gogoreplace 'reserved 38;' '//reserved 38;' ./google/protobuf/descriptor.proto
	gogoreplace 'reserved 8;' '//reserved 8;' ./google/protobuf/descriptor.proto
	gogoreplace 'reserved 9;' '//reserved 9;' ./google/protobuf/descriptor.proto
	gogoreplace 'reserved 4;' '//reserved 4;' ./google/protobuf/descriptor.proto
	gogoreplace 'reserved 5;' '//reserved 5;' ./google/protobuf/descriptor.proto
	gogoreplace 'option go_package = "github.com/golang/protobuf/protoc-gen-go/descriptor;descriptor";' 'option go_package = "descriptor";' ./google/protobuf/descriptor.proto

	(cd ./google/protobuf/compiler && rm plugin.proto; wget ${URL}/compiler/plugin.proto)
	gogoreplace 'option go_package = "github.com/golang/protobuf/protoc-gen-go/plugin;plugin_go";' 'option go_package = "plugin_go";' ./google/protobuf/compiler/plugin.proto
	
	(cd ./google/protobuf && rm any.proto; wget ${URL}/any.proto)
	gogoreplace 'go_package = "github.com/golang/protobuf/ptypes/any";' 'go_package = "types";' ./google/protobuf/any.proto
	(cd ./google/protobuf && rm empty.proto; wget ${URL}/empty.proto)
	gogoreplace 'go_package = "github.com/golang/protobuf/ptypes/empty";' 'go_package = "types";' ./google/protobuf/empty.proto
	(cd ./google/protobuf && rm timestamp.proto; wget ${URL}/timestamp.proto)
	gogoreplace 'go_package = "github.com/golang/protobuf/ptypes/timestamp";' 'go_package = "types";' ./google/protobuf/timestamp.proto
	(cd ./google/protobuf && rm duration.proto; wget ${URL}/duration.proto)
	gogoreplace 'go_package = "github.com/golang/protobuf/ptypes/duration";' 'go_package = "types";' ./google/protobuf/duration.proto
	(cd ./google/protobuf && rm struct.proto; wget ${URL}/struct.proto)
	gogoreplace 'go_package = "github.com/golang/protobuf/ptypes/struct;structpb";' 'go_package = "types";' ./google/protobuf/struct.proto
	(cd ./google/protobuf && rm wrappers.proto; wget ${URL}/wrappers.proto)
	gogoreplace 'go_package = "github.com/golang/protobuf/ptypes/wrappers";' 'go_package = "types";' ./google/protobuf/wrappers.proto
	(cd ./google/protobuf && rm field_mask.proto; wget ${URL}/field_mask.proto)
	gogoreplace 'option go_package = "google.golang.org/genproto/protobuf/field_mask;field_mask";' 'option go_package = "types";' ./google/protobuf/field_mask.proto
	
	
