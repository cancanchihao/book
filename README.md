# book

go-zero api,grpc服务实践

项目包含 borrow-api，library-rpc，user-api，user-rpc，borrow 对 library 与 user 进行 rpc 服务调用

使用 etcd，go zero 管理，user-api 使用 jwt 作为 middleware 进行鉴权，model 层使用 redis 进行缓存优化
