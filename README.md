# kratos_grpc_error

kratos grpc 传递 error 问题最小复现。

## 使用方式

1. `go run ./server` 运行服务端。
2. `go run ./client` 运行客户端。

期望结果：客户端打印 "Is user not found error? true"。
实际结果：客户端打印 "Is user not found error? true"。