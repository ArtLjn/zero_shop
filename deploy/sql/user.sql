CREATE TABLE `user` (
    `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户名称',
    `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '密码',
    `role_id` int NOT NULL COMMENT '角色',
    `phone` varchar(255) DEFAULT NULL COMMENT '手机号',
    `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '电子邮箱',
    `user_id` varchar(255) DEFAULT NULL COMMENT '用户唯一ID',
    `sex` varchar(10) DEFAULT NULL COMMENT '性别',
    PRIMARY KEY (`id`),
    KEY `role` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;