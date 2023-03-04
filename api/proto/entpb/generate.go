package entpb
//go:generate protoc -I=.. --go_out=.. --go-grpc_out=.. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --entgrpc_out=.. --entgrpc_opt=paths=source_relative,schema_path=../../../graph/entgen/schema,ent_package=github.com/woocoos/knockout/ent entpb/entpb.proto
