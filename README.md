## Zero-Shop 微服务商城

## 模块划分

| 子项目名 | 项目路径                          |
|------|-------------------------------|
| 商品服务 | [/app/blog](./app/goods)      |
| 订单服务 | [/app/order](./app/order)     |
| 用户服务 | [/app/user](./app/user)       |
| 支付服务 | [/app/payment](./app/payment) |
| 统一网关 | [/app/gateway](./app/gateway) |
| 静态资源 | [/web](./web)                 |

## 技术栈
微服务
- RabbitMq
- redis
- mongo
- mysql
- jwt
- etcd

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

### mysql数据表设计
1. 用户表(user.sql)
``` sql
CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户名称',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '密码',
  `role_id` int unsigned NOT NULL COMMENT '角色',
  `phone` varchar(255) DEFAULT NULL COMMENT '手机号',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '电子邮箱',
  `user_id` varchar(255) DEFAULT NULL COMMENT '用户唯一ID',
  `sex` varchar(10) DEFAULT NULL COMMENT '性别',
  PRIMARY KEY (`id`),
  KEY `role` (`role_id`),
  CONSTRAINT `role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```
2. 角色表(role.sql)
```sql
CREATE TABLE `role` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `role_id` int unsigned DEFAULT NULL COMMENT '角色ID',
    `role_name` varchar(11) DEFAULT NULL COMMENT '角色名称',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```