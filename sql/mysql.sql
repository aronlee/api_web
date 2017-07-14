CREATE TABLE IF NOT EXISTS `user_info` (
  `u_i_d` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL COMMENT '用户名',
  `email` varchar(128) NOT NULL DEFAULT '' COMMENT '电子邮箱',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '账号状态',
  `is_third` tinyint NOT NULL DEFAULT 0 COMMENT '是否第三方账号，1-是',
  `is_root` tinyint NOT NULL DEFAULT 0 COMMENT '是否是root账号，1-是',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`u_i_d`),
  UNIQUE KEY (`username`),
  UNIQUE KEY (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '用户信息表';

CREATE TABLE IF NOT EXISTS `user_login` (
  `u_i_d` int unsigned NOT NULL,
  `username` varchar(20) NOT NULL COMMENT '用户名',
  `passcode` char(12) NOT NULL DEFAULT '' COMMENT '加密随机数',
  `password` char(32) NOT NULL DEFAULT '' COMMENT 'md5密码',
  `login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后一次登录时间（主动登录或cookie登录）',
  PRIMARY KEY (`u_i_d`),
  UNIQUE KEY (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '用户登陆表';

CREATE TABLE IF NOT EXISTS `article` (
  `a_i_d` int unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL COMMENT '作者名',
  `title` varchar(127) NOT NULL DEFAULT '' COMMENT '文章标题',
  `content` mediumtext NOT NULL COMMENT '正文(带html)',
  `txt` mediumtext NOT NULL COMMENT '正文(纯文本)',
  `c_s_s` mediumtext NOT NULL COMMENT '需要额外引入的css样式',
  `view_num` int unsigned NOT NULL DEFAULT 0 COMMENT '浏览数',
  `comment_num` int unsigned NOT NULL DEFAULT 0 COMMENT '评论数',
  `like_num` int unsigned NOT NULL DEFAULT 0 COMMENT '赞数',
  `status` int unsigned NOT NULL DEFAULT 0 COMMENT '文章状态',
  `u_i_d` int unsigned NOT NULL COMMENT '作者id',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`a_i_d`),
  KEY (`u_i_d`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '文章表';

CREATE TABLE IF NOT EXISTS `comment` (
  `c_i_d` int unsigned NOT NULL AUTO_INCREMENT,
  `floor` int unsigned NOT NULL COMMENT '楼层',
  `content` text NOT NULL COMMENT '评论内容',
  `a_i_d` int unsigned NOT NULL COMMENT '评论的文章id',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论时间',
  `u_i_d` int unsigned NOT NULL COMMENT '评论用户ID',
  PRIMARY KEY (`c_i_d`),
  KEY (`u_i_d`, `a_i_d`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '评论表';

CREATE TABLE IF NOT EXISTS `article_tag` (
  `i_d` int unsigned NOT NULL AUTO_INCREMENT,
  `a_i_d` int unsigned NOT NULL COMMENT '文章id',
  `t_i_d` int unsigned NOT NULL COMMENT 'tagID',
  PRIMARY KEY (`i_d`),
  KEY (`t_i_d`, `a_i_d`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '文章拥有的tags表';

CREATE TABLE IF NOT EXISTS `tag` (
  `t_i_d` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL COMMENT 'tag名称',
  `u_i_d` int unsigned NOT NULL COMMENT '创建用户ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论时间',
  PRIMARY KEY (`t_i_d`),
  UNIQUE KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT 'tags表';

CREATE TABLE IF NOT EXISTS `session` (
  `session_key` char(64) NOT NULL,
  `session_data` blob,
  `session_expiry` int(11) unsigned NOT NULL,
  PRIMARY KEY (`session_key`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;