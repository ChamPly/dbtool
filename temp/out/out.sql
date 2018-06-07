
-- ----------------------------
-- Table structure for epg_sys_dictionary
-- ----------------------------
DROP TABLE IF EXISTS `epg_sys_dictionary`;
CREATE TABLE `epg_sys_dictionary` (
	`id` int(20) NOT NULL  AUTO_INCREMENT COMMENT '编号',
	`value` varchar(32) CHARACTER SET utf8 NOT NULL DEFAULT '-' COMMENT '枚举项的值',
	`description` varchar(64) CHARACTER SET utf8 NOT NULL DEFAULT '-' COMMENT '枚举项的描述',
	`belong_enum` varchar(32) CHARACTER SET utf8 NOT NULL DEFAULT '-' COMMENT '所属枚举',
	`sort_id` int(2) NOT NULL  COMMENT '枚举内排序',
	`group_id` varchar(32) CHARACTER SET utf8 DEFAULT '-' COMMENT '分组',
	`status` int(2) NOT NULL  COMMENT '状态',
	`remark` varchar(64) CHARACTER SET utf8 DEFAULT '-' COMMENT '备注',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci COMMENT='数据字典';