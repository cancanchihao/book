# book

go-zero api,grpc服务实践

项目包含 borrow-api，library-rpc，user-api，user-rpc，borrow 对 library 与 user 进行 rpc 服务调用

使用 etcd，go zero管理，user-api 使用 jwt 作为 middleware 进行鉴权

TODO:

1.在项目中新增 Redis 客户端配置（支持单点/集群）

2.使用 go-zero 提供的 sqlc.CachedConn 做数据库查询缓存，减少重复 SQL 查询

3.在高并发关键操作（如借书、还书）中，使用 RedSync 实现分布式锁，保证业务互斥

4.配置至少 5 个 Redis 节点，传入 RedSync 多池，实现多数决策容错
