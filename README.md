## Zero-Shop 微服务商城

## 模块划分

| 子项目名 | 项目路径                                |
|------|-------------------------------------|
| 商品服务 | [/app/goods](./app/goods/README.md) |
| 订单服务 | [/app/order](./app/order)           |
| 用户服务 | [/app/user](./app/user)             |
| 支付服务 | [/app/payment](./app/payment)       |
| 统一网关 | [/app/gateway](./app/gateway)       |
| 静态资源 | [/web](./web)                       |

## 技术栈
微服务
- RabbitMq
- redis
- mongo
- mysql
- jwt
- etcd
- ElasticSearch

区块链
- [fisco-bcos联盟链平台](https://fisco-bcos-documentation.readthedocs.io/zh-cn/latest/index.html)
- solidity合约

前端静态资源
- vue.js

## 待办
- [ ] 项目数据库设计
- [ ] 项目基础架构搭建
- [ ] 登录注册加入第三方登录
- [ ] 商品信息采用区块链进行溯源
- [ ] 支付模块接入第三方支付
- [ ] 采用消息代理进行流量消峰
