CREATE TABLE `user`  (
                         `uid` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                         `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
                         `passcode` char(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '加密随机数',
                         `passwd` char(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'md5密码',
                         `status` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户账号状态。0-默认；1-冻结；2-停号',
                         `hardware` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'hardware',
                         `ctime` timestamp(0) NOT NULL DEFAULT '2013-03-15 14:38:09',
                         `mtime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
                         PRIMARY KEY (`uid`) USING BTREE,
                         UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户信息表' ROW_FORMAT = Dynamic;

INSERT INTO user(`uid`, `username`, `passcode`, `passwd`, `status`, `hardware`, `ctime`, `mtime`) VALUES (1, 'admin', '123', '30e33d5a27ecbefc44ffb36d02800791', 0, '', '2022-03-15 14:38:09', '2022-01-18 12:45:58');

CREATE TABLE `login_history`  (
                                  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                                  `uid` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户UID',
                                  `state` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '登录状态，0登录，1登出',
                                  `ctime` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '登录时间',
                                  `ip` varchar(31) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'ip',
                                  `hardware` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'hardware',
                                  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户登录表' ROW_FORMAT = Dynamic;

CREATE TABLE `login_last`  (
                               `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
                               `uid` int(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户UID',
                               `login_time` timestamp(0) NULL DEFAULT NULL COMMENT '登录时间',
                               `logout_time` timestamp(0) NULL DEFAULT NULL COMMENT '登出时间',
                               `ip` varchar(31) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'ip',
                               `is_logout` tinyint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否logout,1:logout，0:login',
                               `session` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '会话',
                               `hardware` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'hardware',
                               PRIMARY KEY (`id`) USING BTREE,
                               UNIQUE INDEX `uid`(`uid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '最后一次用户登录表' ROW_FORMAT = Dynamic;

