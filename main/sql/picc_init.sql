insert into picc_user(id,  tel, name,passwd, prov, city, county, post_code ) values('130123198702270036', '13731143001', '张凯', '5e68697e5730726c64240d6ef7d186a0ed1641634282bcc1df28','河北', '石家庄','正定县','130123');
insert into picc_user(id,  tel, name,passwd, prov, city, county, post_code ) values('130123198702270035', '13731143000', '张凯pt', '5e68697e5730726c64240d6ef7d186a0ed1641634282bcc1df28','河北', '石家庄','正定县','130123');

insert into picc_role(id, name) values(1,'超级管理员');
insert into picc_role(id, name) values(2,'承保角色');
insert into picc_role(id, name) values(3,'核保角色');
insert into picc_role(id, name) values(4,'定损角色');

insert into picc_entity(id, name, alias) values(1,'ad', '用户管理');
insert into picc_entity(id, name, alias) values(2,'su', '标的管理');
insert into picc_entity(id, name, alias) values(3,'ud', '承保');
insert into picc_entity(id, name, alias) values(4,'ve', '核保');
insert into picc_entity(id, name, alias) values(5,'lo', '定损');
insert into picc_entity(id, name, alias) values(6,'pm', '保单管理');
insert into picc_entity(id, name, alias) values(7,'st', '统计');
insert into picc_entity(id, name, alias) values(8,'ll', '大面积定损');

insert into picc_user_role(user_id, role_id) values('130123198702270036', 1);
insert into picc_user_role(user_id, role_id) values('130123198702270035', 2);
insert into picc_user_role(user_id, role_id) values('130123198702270035', 3);

insert into picc_role_privilege(role_id, entity_id, can_add, can_read, can_update, can_del, can_assign) values(1, 1, 1,1,1,1,1);
insert into picc_role_privilege(role_id, entity_id, can_add, can_read, can_update, can_del, can_assign) values(1, 2, 1,1,1,1,1);
insert into picc_role_privilege(role_id, entity_id, can_add, can_read, can_update, can_del, can_assign) values(1, 3, 1,1,1,1,1);
insert into picc_role_privilege(role_id, entity_id, can_add, can_read, can_update, can_del, can_assign) values(1, 4, 1,1,1,1,1);
insert into picc_role_privilege(role_id, entity_id, can_add, can_read, can_update, can_del, can_assign) values(1, 5, 1,1,1,1,1);
insert into picc_role_privilege(role_id, entity_id, can_add, can_read, can_update, can_del, can_assign) values(1, 6, 1,1,1,1,1);
insert into picc_role_privilege(role_id, entity_id, can_add, can_read, can_update, can_del, can_assign) values(1, 7, 1,1,1,1,1);
insert into picc_role_privilege(role_id, entity_id, can_add, can_read, can_update, can_del, can_assign) values(1, 8, 1,1,1,1,1);
insert into picc_role_privilege(role_id, entity_id, can_add, can_read, can_update, can_del, can_assign) values(2, 3, 1,1,1,1,1);
insert into picc_role_privilege(role_id, entity_id, can_add, can_read, can_update, can_del, can_assign) values(2, 6, 1,1,1,1,1);
insert into picc_role_privilege(role_id, entity_id, can_add, can_read, can_update, can_del, can_assign) values(3, 4, 1,1,1,1,1);
insert into picc_role_privilege(role_id, entity_id, can_add, can_read, can_update, can_del, can_assign) values(3, 4, 1,1,1,1,1);
insert into picc_role_privilege(role_id, entity_id, can_add, can_read, can_update, can_del, can_assign) values(4, 5, 1,1,1,1,1);
insert into picc_role_privilege(role_id, entity_id, can_add, can_read, can_update, can_del, can_assign) values(4, 6, 1,1,1,1,1);


