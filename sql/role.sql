CREATE TABLE `role` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `role_id` int unsigned DEFAULT NULL COMMENT '角色ID',
    `role_name` varchar(11) DEFAULT NULL COMMENT '角色名称',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;