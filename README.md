# 配置RabbitMQ

在节点启动的config.toml中添加如下配置
```
[Eth.Downstream]
URI= ["amqp://guest:guest@localhost:5672"]
Exchange = "test-exchange"
RoutingKey = "test-key"
RetryInterval = 500 #ms
TimeoutInterval = 5000 #ms
```

# 通过RPC导出Block信息到文件
通过http请求
```
curl --data '{"method":"eth_dumpBlock","params":["0x1", "0x11", "/home/op/blocks.json"],"id":1,"jsonrpc":"2.0"}' -H "Content-Type: application/json" -X POST localhost:8545
```
通过js控制台操作
```
geth attach data/eth.ipc

$ eth.dumpBlock(1, 100, "/home/op/blocks.json")
```

> 同步模式问题，如果最开始使用fast/snap模式进行区块同步，那么在节点追到最新区块后，需要将同步模式改为full，防止后续的区块执行不会出现不连续的情况；