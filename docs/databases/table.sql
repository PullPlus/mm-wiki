-- --------------------------------
-- MM-Wiki <LABEL_2200>
-- author: phachon
-- --------------------------------

-- <LABEL_2002>
-- CREATE DATABASE IF NOT EXISTS mm_wiki DEFAULT CHARSET utf8;

-- --------------------------------
-- <LABEL_2201>
-- --------------------------------
DROP TABLE IF EXISTS `mw_user`;
CREATE TABLE `mw_user` (
  `user_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2218> id',
  `username` varchar(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2202>',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '<LABEL_2219>',
  `given_name` varchar(50) NOT NULL DEFAULT '' COMMENT '<LABEL_2220>',
  `mobile` char(13) NOT NULL DEFAULT '' COMMENT '<LABEL_2203>',
  `phone` char(13) NOT NULL DEFAULT '' COMMENT '<LABEL_2221>',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '<LABEL_2222>',
  `department` char(50) NOT NULL DEFAULT '' COMMENT '<LABEL_2223>',
  `position` char(50) NOT NULL DEFAULT '' COMMENT '<LABEL_2224>',
  `location` char(50) NOT NULL DEFAULT '' COMMENT '<LABEL_2225>',
  `im` char(50) NOT NULL DEFAULT '' COMMENT '<LABEL_2068>',
  `last_ip` varchar(15) NOT NULL DEFAULT '' COMMENT '<LABEL_2163>ip',
  `last_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2069>',
  `role_id` tinyint(3) NOT NULL DEFAULT '0' COMMENT '<LABEL_2226> id',
  `is_forbidden` tinyint(3) NOT NULL DEFAULT '0' COMMENT '<LABEL_2164>，0 <LABEL_2250> 1 <LABEL_2251>',
  `is_delete` tinyint(3) NOT NULL DEFAULT '0' COMMENT '<LABEL_2165>，0 <LABEL_2250> 1 <LABEL_2251>',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2167>',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2201>';

-- ---------------------------------------------------------------
-- <LABEL_2087>
-- ---------------------------------------------------------------
DROP TABLE IF EXISTS `mw_role`;
CREATE TABLE `mw_role` (
  `role_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2226> id',
  `name` char(10) NOT NULL DEFAULT '' COMMENT '<LABEL_2168>',
  `type` tinyint(3) NOT NULL DEFAULT '0' COMMENT '<LABEL_2169> 0 <LABEL_2088>，1 <LABEL_2110>',
  `is_delete` tinyint(3) NOT NULL DEFAULT '0' COMMENT '<LABEL_2165>，0 <LABEL_2250> 1 <LABEL_2251>',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '<LABEL_2167>',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2087>';

-- -------------------------------------------------------
-- <LABEL_2089>
-- -------------------------------------------------------
DROP TABLE IF EXISTS `mw_privilege`;
CREATE TABLE `mw_privilege` (
  `privilege_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2227>id',
  `name` char(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2204>',
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '<LABEL_2228>',
  `type` enum('controller','menu') DEFAULT 'controller' COMMENT '<LABEL_2170>：<LABEL_2205>、<LABEL_2229>',
  `controller` char(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2205>',
  `action` char(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2230>',
  `icon` char(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2231>（<LABEL_2171>)',
  `target` char(200) NOT NULL DEFAULT '' COMMENT '<LABEL_2172>',
  `is_display` tinyint(1) NOT NULL DEFAULT '0' COMMENT '<LABEL_2173>：0<LABEL_2206> 1<LABEL_2232>',
  `sequence` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2233>(<LABEL_2090>)',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '<LABEL_2167>',
  PRIMARY KEY (`privilege_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2089>';

-- ------------------------------------------------------------------
-- <LABEL_2005>
-- ------------------------------------------------------------------
DROP TABLE IF EXISTS `mw_role_privilege`;
CREATE TABLE `mw_role_privilege` (
  `role_privilege_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2070> id',
  `role_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '<LABEL_2226>id',
  `privilege_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '<LABEL_2227>id',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  PRIMARY KEY (`role_privilege_id`),
  KEY (`role_id`),
  KEY (`privilege_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2005>';

-- --------------------------------
-- <LABEL_2207>
-- --------------------------------
DROP TABLE IF EXISTS `mw_space`;
CREATE TABLE `mw_space` (
  `space_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2234> id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '<LABEL_2235>',
  `description` varchar(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2236>',
  `tags` varchar(255) NOT NULL DEFAULT '' COMMENT '<LABEL_2237>',
  `visit_level` enum('private','public') NOT NULL DEFAULT 'public' COMMENT '<LABEL_2174>：private,public',
  `is_share` tinyint(3) NOT NULL DEFAULT '1' COMMENT '<LABEL_2021> 0 <LABEL_2250> 1 <LABEL_2251>',
  `is_export` tinyint(3) NOT NULL DEFAULT '1' COMMENT '<LABEL_2022> 0 <LABEL_2250> 1 <LABEL_2251>',
  `is_delete` tinyint(3) NOT NULL DEFAULT '0' COMMENT '<LABEL_2165> 0 <LABEL_2250> 1 <LABEL_2251>',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2167>',
  PRIMARY KEY (`space_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2207>';

-- --------------------------------
-- <LABEL_2091>
-- --------------------------------
DROP TABLE IF EXISTS `mw_space_user`;
CREATE TABLE `mw_space_user` (
  `space_user_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2071> id',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2218> id',
  `space_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2234> id',
  `privilege` tinyint(3) NOT NULL DEFAULT '0' COMMENT '<LABEL_2023> 0 <LABEL_2208> 1 <LABEL_2209> 2 <LABEL_2199>',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2175>',
  PRIMARY KEY (`space_user_id`),
  UNIQUE KEY (`user_id`, `space_id`),
  KEY (`user_id`),
  KEY (`space_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2091>';

-- --------------------------------
-- <LABEL_2210>
-- --------------------------------
DROP TABLE IF EXISTS `mw_document`;
CREATE TABLE `mw_document` (
  `document_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2238> id',
  `parent_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2211> id',
  `space_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2234>id',
  `name` varchar(150) NOT NULL DEFAULT '' COMMENT '<LABEL_2176>',
  `type` tinyint(3) NOT NULL DEFAULT '1' COMMENT '<LABEL_2177> 1 page 2 dir',
  `path` char(30) NOT NULL DEFAULT '0' COMMENT '<LABEL_2009> document_id <LABEL_2092> 0,1,2,...',
  `sequence` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2212>(<LABEL_2090>)',
  `create_user_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2178> id',
  `edit_user_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2072> id',
  `is_delete` tinyint(3) NOT NULL DEFAULT '0' COMMENT '<LABEL_2165> 0 <LABEL_2250> 1 <LABEL_2251>',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2167>',
  PRIMARY KEY (`document_id`),
  KEY (`parent_id`),
  KEY (`space_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2210>';

-- --------------------------------
-- <LABEL_2093>
-- --------------------------------
DROP TABLE IF EXISTS `mw_collection`;
CREATE TABLE `mw_collection` (
  `collection_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2073> id',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2218>id',
  `type` tinyint(3) NOT NULL DEFAULT '1' COMMENT '<LABEL_2179> 1 <LABEL_2238> 2 <LABEL_2234>',
  `resource_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2180> id ',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  PRIMARY KEY (`collection_id`),
  KEY (`user_id`),
  UNIQUE key (`user_id`, `resource_id`, `type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2093>';

-- --------------------------------
-- <LABEL_2094>
-- --------------------------------
DROP TABLE IF EXISTS `mw_follow`;
CREATE TABLE `mw_follow` (
  `follow_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2239> id',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2218>id',
  `type` tinyint(3) NOT NULL DEFAULT '1' COMMENT '<LABEL_2181> 1 <LABEL_2238> 2 <LABEL_2218>',
  `object_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2182> id',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  PRIMARY KEY (`follow_id`),
  KEY (`user_id`),
  KEY (`object_id`),
  UNIQUE key (`user_id`, `object_id`, `type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2094>';

-- --------------------------------
-- <LABEL_2095>
-- --------------------------------
DROP TABLE IF EXISTS `mw_log_document`;
CREATE TABLE `mw_log_document` (
  `log_document_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2144> id',
  `document_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2238>id',
  `space_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2234>id',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2218>id',
  `action` tinyint(3) NOT NULL DEFAULT '1' COMMENT '<LABEL_2230> 1 <LABEL_2240> 2 <LABEL_2241> 3 <LABEL_2242>',
  `comment` varchar(255) NOT NULL DEFAULT '' COMMENT '<LABEL_2183>',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  PRIMARY KEY (`log_document_id`),
  KEY (`document_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2095>';

-- --------------------------------
-- <LABEL_2035>
-- --------------------------------
DROP TABLE IF EXISTS `mw_log`;
CREATE TABLE `mw_log` (
  `log_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2074> id',
  `level` tinyint(3) NOT NULL DEFAULT '6' COMMENT '<LABEL_2184>',
  `path` char(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2185>',
  `get` text NOT NULL COMMENT 'get<LABEL_2243>',
  `post` text NOT NULL COMMENT 'post<LABEL_2243>',
  `message` varchar(255) NOT NULL DEFAULT '' COMMENT '<LABEL_2244>',
  `ip` char(100) NOT NULL DEFAULT '' COMMENT 'ip<LABEL_2245>',
  `user_agent` char(200) NOT NULL DEFAULT '' COMMENT '<LABEL_2186>',
  `referer` char(100) NOT NULL DEFAULT '' COMMENT 'referer',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2218>id',
  `username` char(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2202>',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  PRIMARY KEY (`log_id`),
  KEY (`level`, `username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2035>';

-- --------------------------------
-- <LABEL_2075>
-- --------------------------------
DROP TABLE IF EXISTS `mw_email`;
CREATE TABLE `mw_email` (
  `email_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2222> id',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2036>',
  `sender_address` varchar(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2037>',
  `sender_name` varchar(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2076>',
  `sender_title_prefix` varchar(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2024>',
  `host` char(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2077>',
  `port` int(5) NOT NULL DEFAULT '25' COMMENT '<LABEL_2096>',
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '<LABEL_2202>',
  `password` varchar(50) NOT NULL DEFAULT '' COMMENT '<LABEL_2219>',
  `is_ssl` tinyint(1) NOT NULL DEFAULT '0' COMMENT '<LABEL_2187>ssl， 0 <LABEL_2097> 1 <LABEL_2246>',
  `is_used` tinyint(3) NOT NULL DEFAULT '0' COMMENT '<LABEL_2098>， 0 <LABEL_2097> 1 <LABEL_2246>',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2167>',
  PRIMARY KEY (`email_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2075>';

-- --------------------------------
-- <LABEL_2099>
-- --------------------------------
DROP TABLE IF EXISTS `mw_link`;
CREATE TABLE `mw_link` (
  `link_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2247> id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '<LABEL_2188>',
  `url` varchar(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2189>',
  `sequence` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2212>(<LABEL_2090>)',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2167>',
  PRIMARY KEY (`link_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2099>';

-- --------------------------------
-- <LABEL_2038>
-- --------------------------------
DROP TABLE IF EXISTS `mw_login_auth`;
CREATE TABLE `mw_login_auth` (
  `login_auth_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2100>ID',
  `name` varchar(30) NOT NULL COMMENT '<LABEL_2078>',
  `username_prefix` varchar(30) NOT NULL COMMENT '<LABEL_2101>',
  `url` varchar(200) NOT NULL COMMENT '<LABEL_2190> url',
  `ext_data` varchar(500) NOT NULL DEFAULT '' COMMENT '<LABEL_2191>: token=aaa&key=bbb',
  `is_used` tinyint(1) NOT NULL DEFAULT '0' COMMENT '<LABEL_2098>， 0 <LABEL_2097> 1 <LABEL_2246>',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '<LABEL_2165> 0 <LABEL_2250> 1 <LABEL_2251>',
  `create_time` int(11) NOT NULL COMMENT '<LABEL_2166>',
  `update_time` int(11) NOT NULL COMMENT '<LABEL_2167>',
  PRIMARY KEY (`login_auth_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2038>';

-- --------------------------------
-- <LABEL_2102>
-- --------------------------------
DROP TABLE IF EXISTS `mw_config`;
CREATE TABLE `mw_config` (
  `config_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2103>Id',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2192>',
  `key` char(50) NOT NULL DEFAULT '' COMMENT '<LABEL_2213>',
  `value` text NOT NULL COMMENT '<LABEL_2214>',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2167>',
  PRIMARY KEY (`config_id`),
  unique KEY (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2102>';

-- --------------------------------
-- <LABEL_2079>
-- --------------------------------
DROP TABLE IF EXISTS `mw_contact`;
CREATE TABLE `mw_contact` (
  `contact_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2215> id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '<LABEL_2104>',
  `mobile` char(13) NOT NULL DEFAULT '' COMMENT '<LABEL_2193>',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '<LABEL_2222>',
  `position` varchar(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2105>',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2167>',
  PRIMARY KEY (`contact_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2194>';

-- --------------------------------
-- <LABEL_2106>
-- --------------------------------
DROP TABLE IF EXISTS `mw_attachment`;
CREATE TABLE `mw_attachment` (
  `attachment_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '<LABEL_2248> id',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2178>id',
  `document_id` int(10) NOT NULL DEFAULT '0' COMMENT '<LABEL_2195>id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '<LABEL_2196>',
  `path` varchar(100) NOT NULL DEFAULT '' COMMENT '<LABEL_2197>',
  `source` tinyint(1) NOT NULL DEFAULT '0' COMMENT '<LABEL_2198>， 0 <LABEL_2107> 1 <LABEL_2249>',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2166>',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '<LABEL_2167>',
  PRIMARY KEY (`attachment_id`),
  KEY (`document_id`, `source`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='<LABEL_2106>';
