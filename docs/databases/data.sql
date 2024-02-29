-- --------------------------------------
-- MM-Wiki <LABEL_2108>
-- warning：<LABEL_2000>
-- author: phachon
-- --------------------------------------


-- --------------------------------------
-- <LABEL_2109>（root）, password：123456<LABEL_2004>，<LABEL_2003>
-- --------------------------------------
-- INSERT INTO `mw_user` (`user_id`, `username`, `password`, `given_name`, `email`,  `mobile`, `role_id`, `is_delete`, `create_time`, `update_time`)
-- VALUES ('1', 'root', 'e10adc3949ba59abbe56e057f20f883e', 'root', 'root@123456.com', '1102222', '1', '0', unix_timestamp(now()), unix_timestamp(now()));


-- --------------------------------------
-- <LABEL_2110> 1 <LABEL_2080>，2 <LABEL_2199>，3 <LABEL_2111>
-- --------------------------------------
INSERT INTO `mw_role` (`role_id`, `name`, `type`, `is_delete`, `create_time`, `update_time`) VALUES ('1', '<LABEL_2080>', '1', '0', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_role` (`role_id`, `name`, `type`, `is_delete`, `create_time`, `update_time`) VALUES ('2', '<LABEL_2199>', '1', '0', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_role` (`role_id`, `name`, `type`, `is_delete`, `create_time`, `update_time`) VALUES ('3', '<LABEL_2111>', '1', '0', unix_timestamp(now()), unix_timestamp(now()));


-- --------------------------------------
-- <LABEL_2112>
-- --------------------------------------
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (1, '<LABEL_2113>', 0, 'menu', '', '', 'glyphicon-leaf', '', 1, 1, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (2, '<LABEL_2114>', 1, 'controller', 'profile', 'info', 'glyphicon-list', '', 1, 11, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (3, '<LABEL_2115>', 1, 'controller', 'profile', 'edit', 'glyphicon-list', '', 0, 12, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (4, '<LABEL_2039>', 1, 'controller', 'profile', 'modify', 'glyphicon-list', '', 0, 13, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (5, '<LABEL_2040>', 1, 'controller', 'profile', 'followUser', 'glyphicon-list', '', 0, 14, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (6, '<LABEL_2041>', 1, 'controller', 'profile', 'followDoc', 'glyphicon-list', '', 0, 15, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (7, '<LABEL_2116>', 1, 'controller', 'profile', 'activity', 'glyphicon-list', '', 1, 16, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (8, '<LABEL_2117>', 1, 'controller', 'profile', 'password', 'glyphicon-list', '', 1, 17, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (9, '<LABEL_2042>', 1, 'controller', 'profile', 'savePass', 'glyphicon-list', '', 0, 18, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (10, '<LABEL_2118>', 1, 'menu', '', '', 'glyphicon-user', '', 1, 2, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (11, '<LABEL_2119>', 10, 'controller', 'user', 'add', 'glyphicon-list', '', 1, 21, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (12, '<LABEL_2043>', 10, 'controller', 'user', 'save', 'glyphicon-list', '', 0, 22, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (13, '<LABEL_2120>', 10, 'controller', 'user', 'list', 'glyphicon-list', '', 1, 23, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (14, '<LABEL_2121>', 10, 'controller', 'user', 'edit', 'glyphicon-list', '', 0, 24, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (15, '<LABEL_2044>', 10, 'controller', 'user', 'modify', 'glyphicon-list', '', 0, 25, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (16, '<LABEL_2122>', 10, 'controller', 'user', 'forbidden', 'glyphicon-list', '', 0, 26, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (17, '<LABEL_2123>', 10, 'controller', 'user', 'recover', 'glyphicon-list', '', 0, 27, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (18, '<LABEL_2124>', 10, 'controller', 'user', 'info', 'glyphicon-list', '', 0, 28, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (19, '<LABEL_2125>', 1, 'menu', '', '', 'glyphicon-gift', '', 1, 3, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (20, '<LABEL_2126>', 19, 'controller', 'role', 'add', 'glyphicon-list', '', 1, 31, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (21, '<LABEL_2045>', 19, 'controller', 'role', 'save', 'glyphicon-list', '', 0, 32, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (22, '<LABEL_2127>', 19, 'controller', 'role', 'list', 'glyphicon-list', '', 1, 33, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (23, '<LABEL_2128>', 19, 'controller', 'role', 'edit', 'glyphicon-list', '', 0, 34, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (24, '<LABEL_2046>', 19, 'controller', 'role', 'modify', 'glyphicon-list', '', 0, 35, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (25, '<LABEL_2047>', 19, 'controller', 'role', 'user', 'glyphicon-list', '', 0, 36, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (26, '<LABEL_2129>', 19, 'controller', 'role', 'privilege', 'glyphicon-list', '', 0, 37, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (27, '<LABEL_2048>', 19, 'controller', 'role', 'grantPrivilege', 'glyphicon-list', '', 0, 38, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (28, '<LABEL_2130>', 19, 'controller', 'role', 'delete', 'glyphicon-list', '', 0, 29, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (29, '<LABEL_2049>', 19, 'controller', 'role', 'resetUser', 'glyphicon-list', '', 0, 310, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (30, '<LABEL_2131>', 1, 'menu', '', '', 'glyphicon-lock', '', 1, 4, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (31, '<LABEL_2132>', 30, 'controller', 'privilege', 'add', 'glyphicon-list', '', 1, 41, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (32, '<LABEL_2050>', 30, 'controller', 'privilege', 'save', 'glyphicon-list', '', 0, 42, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (33, '<LABEL_2133>', 30, 'controller', 'privilege', 'list', 'glyphicon-list', '', 1, 43, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (34, '<LABEL_2134>', 30, 'controller', 'privilege', 'edit', 'glyphicon-list', '', 0, 44, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (35, '<LABEL_2051>', 30, 'controller', 'privilege', 'modify', 'glyphicon-list', '', 0, 45, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (36, '<LABEL_2135>', 30, 'controller', 'privilege', 'delete', 'glyphicon-list', '', 0, 46, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (37, '<LABEL_2136>', 1, 'menu', '', '', 'glyphicon-th-large', '', 1, 5, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (38, '<LABEL_2137>', 37, 'controller', 'space', 'add', 'glyphicon-list', '', 1, 51, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (39, '<LABEL_2052>', 37, 'controller', 'space', 'save', 'glyphicon-list', '', 0, 52, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (40, '<LABEL_2138>', 37, 'controller', 'space', 'list', 'glyphicon-list', '', 1, 53, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (41, '<LABEL_2139>', 37, 'controller', 'space', 'edit', 'glyphicon-list', '', 0, 54, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (42, '<LABEL_2053>', 37, 'controller', 'space', 'modify', 'glyphicon-list', '', 0, 55, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (43, '<LABEL_2054>', 37, 'controller', 'space', 'member', 'glyphicon-list', '', 0, 56, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (44, '<LABEL_2055>', 37, 'controller', 'space_user', 'save', 'glyphicon-list', '', 0, 57, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (45, '<LABEL_2056>', 37, 'controller', 'space_user', 'remove', 'glyphicon-list', '', 0, 58, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (46, '<LABEL_2012>', 37, 'controller', 'space_user', 'modify', 'glyphicon-list', '', 0, 59, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (47, '<LABEL_2140>', 37, 'controller', 'space', 'delete', 'glyphicon-list', '', 0, 510, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (48, '<LABEL_2141>', 37, 'controller', 'space', 'download', 'glyphicon-list', '', 0, 512, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (49, '<LABEL_2142>', 1, 'menu', '', '', 'glyphicon-list-alt', '', 1, 6, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (50, '<LABEL_2143>', 49, 'controller', 'log', 'system', 'glyphicon-list', '', 1, 61, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (51, '<LABEL_2057>', 49, 'controller', 'log', 'info', 'glyphicon-list', '', 0, 62, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (52, '<LABEL_2144>', 49, 'controller', 'log', 'document', 'glyphicon-list', '', 1, 63, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (53, '<LABEL_2145>', 1, 'menu', '', '', 'glyphicon-cog', '', 1, 7, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (54, '<LABEL_2146>', 53, 'controller', 'config', 'global', 'glyphicon-list', '', 1, 71, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (55, '<LABEL_2058>', 53, 'controller', 'config', 'modify', 'glyphicon-list', '', 0, 72, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (56, '<LABEL_2147>', 53, 'controller', 'email', 'list', 'glyphicon-list', '', 1, 73, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (57, '<LABEL_2025>', 53, 'controller', 'email', 'add', 'glyphicon-list', '', 0, 74, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (58, '<LABEL_2010>', 53, 'controller', 'email', 'save', 'glyphicon-list', '', 0, 75, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (59, '<LABEL_2026>', 53, 'controller', 'email', 'edit', 'glyphicon-list', '', 0, 76, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (60, '<LABEL_2011>', 53, 'controller', 'email', 'modify', 'glyphicon-list', '', 0, 77, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (61, '<LABEL_2027>', 53, 'controller', 'email', 'used', 'glyphicon-list', '', 0, 78, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (62, '<LABEL_2028>', 53, 'controller', 'email', 'delete', 'glyphicon-list', '', 0, 79, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (63, '<LABEL_2148>', 53, 'controller', 'auth', 'list', 'glyphicon-list', '', 1, 81, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (64, '<LABEL_2059>', 53, 'controller', 'auth', 'add', 'glyphicon-list', '', 0, 82, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (65, '<LABEL_2013>', 53, 'controller', 'auth', 'save', 'glyphicon-list', '', 0, 83, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (66, '<LABEL_2060>', 53, 'controller', 'auth', 'edit', 'glyphicon-list', '', 0, 84, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (67, '<LABEL_2014>', 53, 'controller', 'auth', 'modify', 'glyphicon-list', '', 0, 85, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (68, '<LABEL_2061>', 53, 'controller', 'auth', 'delete', 'glyphicon-list', '', 0, 86, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (69, '<LABEL_2062>', 53, 'controller', 'auth', 'used', 'glyphicon-list', '', 0, 87, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (70, '<LABEL_2063>', 53, 'controller', 'auth', 'doc', 'glyphicon-list', '', 0, 88, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (71, '<LABEL_2149>', 1, 'menu', '', '', 'glyphicon-link', '', 1, 8, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (72, '<LABEL_2150>', 71, 'controller', 'link', 'list', 'glyphicon-list', '', 1, 81, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (73, '<LABEL_2151>', 71, 'controller', 'link', 'add', 'glyphicon-list', '', 0, 82, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (74, '<LABEL_2064>', 71, 'controller', 'link', 'save', 'glyphicon-list', '', 0, 83, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (75, '<LABEL_2152>', 71, 'controller', 'link', 'edit', 'glyphicon-list', '', 0, 84, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (76, '<LABEL_2065>', 71, 'controller', 'link', 'modify', 'glyphicon-list', '', 0, 85, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (77, '<LABEL_2153>', 71, 'controller', 'link', 'delete', 'glyphicon-list', '', 0, 86, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (78, '<LABEL_2081>', 71, 'controller', 'contact', 'list', 'glyphicon-list', '', 1, 91, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (79, '<LABEL_2082>', 71, 'controller', 'contact', 'add', 'glyphicon-list', '', 0, 92, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (80, '<LABEL_2029>', 71, 'controller', 'contact', 'save', 'glyphicon-list', '', 0, 93, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (81, '<LABEL_2083>', 71, 'controller', 'contact', 'edit', 'glyphicon-list', '', 0, 94, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (82, '<LABEL_2030>', 71, 'controller', 'contact', 'modify', 'glyphicon-list', '', 0, 95, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (83, '<LABEL_2084>', 71, 'controller', 'contact', 'delete', 'glyphicon-list', '', 0, 96, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (84, '<LABEL_2154>', 1, 'menu', '', '', 'glyphicon-signal', '', 1, 9, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (85, '<LABEL_2155>', 84, 'controller', 'static', 'default', 'glyphicon-list', '', 1, 91, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (86, 'ajax<LABEL_2015>', 84, 'controller', 'static', 'spaceDocsRank', 'glyphicon-list', '', 0, 92, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (87, 'ajax<LABEL_2016>', 84, 'controller', 'static', 'collectDocRank', 'glyphicon-list', '', 0, 93, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (88, 'ajax<LABEL_2006>', 84, 'controller', 'static', 'docCountByTime', 'glyphicon-list', '', 0, 94, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (89, '<LABEL_2156>', 84, 'controller', 'static', 'monitor', 'glyphicon-list', '', 1, 95, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (90, 'ajax<LABEL_2031>', 84, 'controller', 'static', 'serverStatus', 'glyphicon-list', '', 0, 96, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (91, 'ajax<LABEL_2032>', 84, 'controller', 'static', 'serverTime', 'glyphicon-list', '', 0, 97, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (92, '<LABEL_2033>', 53, 'controller', 'email', 'test', 'glyphicon-list', '', 0, 80, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (93, '<LABEL_2085>', 71, 'controller', 'contact', 'import', 'glyphicon-list', '', 0, 97, unix_timestamp(now()), unix_timestamp(now()));



-- ---------------------------------------------
-- <LABEL_2007>
-- ---------------------------------------------
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (1, 3, 1, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (2, 3, 2, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (3, 3, 3, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (4, 3, 4, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (5, 3, 5, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (6, 3, 6, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (7, 3, 7, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (8, 3, 8, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (9, 3, 9, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (10, 2, 1, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (11, 2, 2, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (12, 2, 3, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (13, 2, 4, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (14, 2, 5, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (15, 2, 6, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (16, 2, 7, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (17, 2, 8, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (18, 2, 9, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (19, 2, 37, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (20, 2, 38, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (21, 2, 39, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (22, 2, 40, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (23, 2, 41, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (24, 2, 42, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (25, 2, 43, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (26, 2, 44, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (27, 2, 45, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (28, 2, 46, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (29, 2, 47, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (30, 2, 48, unix_timestamp(now()));


-- -------------------------------------------
-- <LABEL_2157>
-- -------------------------------------------
INSERT INTO `mw_config` VALUES ('1', '<LABEL_2158>', 'main_title', '<LABEL_2008>，<LABEL_2216>：<LABEL_2159> XXXX <LABEL_2160> wiki <LABEL_2217>！', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_config` VALUES ('2', '<LABEL_2161>', 'main_description', '<LABEL_2034>：<LABEL_2017>，<LABEL_2001> root@xxx.com！', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_config` VALUES ('3', '<LABEL_2018>', 'auto_follow_doc_open', '', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_config` VALUES ('4', '<LABEL_2019>', 'send_email_open', '', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_config` VALUES ('5', '<LABEL_2020>', 'sso_open', '', unix_timestamp(now()), unix_timestamp(now()));
-- INSERT INTO `mw_config` (name, key, value, create_time, update_time) VALUES ('<LABEL_2086>', 'system_version', 'v0.0.0', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_config` (name, `key`, value, create_time, update_time) VALUES ('<LABEL_2066>', 'fulltext_search_open', '1', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_config` (name, `key`, value, create_time, update_time) VALUES ('<LABEL_2067>', 'doc_search_timer', '3600', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_config` (name, `key`, value, create_time, update_time) VALUES ('<LABEL_2162>', 'system_name', 'Markdown Mini Wiki', unix_timestamp(now()), unix_timestamp(now()));
