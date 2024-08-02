# gin-project-gen

This is a tool to generate a gin project template.

## Install
```shell
$ go install github.com/trumanwong/gin-project-gen@latest
```

## Create a gin project
```shell
# Create a template project
$ gin-project-gen new helloworld
```

## Adding Proto files
```shell
# Add a proto file
$ gin-project-gen proto add api/app/v1/demo.proto
```
Output:
api/app/v1/demo.proto
```
syntax = "proto3";

package api.app.v1;

option go_package = "gin-layout/api/app/v1;v1";

service Demo {
	rpc CreateDemo (CreateDemoRequest) returns (CreateDemoReply);
	rpc UpdateDemo (UpdateDemoRequest) returns (UpdateDemoReply);
	rpc DeleteDemo (DeleteDemoRequest) returns (DeleteDemoReply);
	rpc GetDemo (GetDemoRequest) returns (GetDemoReply);
	rpc ListDemo (ListDemoRequest) returns (ListDemoReply);
}

message CreateDemoRequest {}
message CreateDemoReply {}

message UpdateDemoRequest {}
message UpdateDemoReply {}

message DeleteDemoRequest {}
message DeleteDemoReply {}

message GetDemoRequest {}
message GetDemoReply {}

message ListDemoRequest {}
message ListDemoReply {}
```

## Generate Proto files
```shell
$ gin-project-gen proto client api/app/v1/demo.proto
```
output:
```shell
api/app/v1/demo.pb.go
api/app/v1/demo_grpc.pb.go
api/app/v1/demo_http.pb.go
```

## Tool upgrade
```shell
$ gin-project-gen upgrade
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/trumanwong/gin-project-gen/blob/main/LICENSE) file for details.

## Acknowledgements

The following project had particular influence on protoc-gen-go-gin's design.

- [go-kratos/kratos](https://github.com/go-kratos/kratos) is a microservice-oriented governance framework implemented by golang.
- [gin-gonic/gin](https://github.com/gin-gonic/gin) is a web framework written in Go.
- <a href="https://jb.gg/OpenSourceSupport"><img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg?_gl=1*1nuywz*_ga*NTcwMDkwNDIxLjE2ODQzMTI1Mzg.*_ga_9J976DJZ68*MTY4NDMxMjUzOC4xLjEuMTY4NDMxMjU1Mi4wLjAuMA.." width="60" height="60"><img src="https://resources.jetbrains.com/storage/products/company/brand/logos/GoLand.svg?_gl=1*1nuywz*_ga*NTcwMDkwNDIxLjE2ODQzMTI1Mzg.*_ga_9J976DJZ68*MTY4NDMxMjUzOC4xLjEuMTY4NDMxMjU1Mi4wLjAuMA.." width="60" height="60"><img src="https://resources.jetbrains.com/storage/products/company/brand/logos/GoLand_icon.svg?_gl=1*1b2zdbh*_ga*NTcwMDkwNDIxLjE2ODQzMTI1Mzg.*_ga_9J976DJZ68*MTY4NDMxMjUzOC4xLjEuMTY4NDMxMjU1Mi4wLjAuMA.." width="60" height="60"></a>
