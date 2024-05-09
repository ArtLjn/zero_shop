CREATE TABLE `goods` (
     `img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '图片',
     `preface` text COMMENT '简介',
     `price` varchar(255) DEFAULT NULL COMMENT '价格',
     `good_id` varchar(255) DEFAULT NULL COMMENT '商品ID',
     `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
     `seller_id` varchar(255) DEFAULT NULL COMMENT '商家ID',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;