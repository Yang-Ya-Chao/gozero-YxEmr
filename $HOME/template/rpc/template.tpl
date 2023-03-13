syntax = "proto3";

package {{.package}};
option go_package="./{{.package}}";

message Req {
  int64 Ibrlx = 1;
  string Cbrh = 2;
  string Csqdh = 3;
  repeated string Cztbm  = 4;
}

message Resp {
  string Data = 1;
}

service {{.serviceName}} {
  rpc Do(Req) returns(Resp);
  rpc Co(Req) returns(Resp);
}
//使用YxEmr模板生成
//goctl rpc protoc reg.proto --go_out=. --go-grpc_out=. --zrpc_out=. --home D:/GoWork/src/YxEmr/$HOME/template
//默认模板生成
//goctl rpc protoc reg.proto --go_out=. --go-grpc_out=. --zrpc_out=.