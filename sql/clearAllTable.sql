-- 禁用外键约束
SET FOREIGN_KEY_CHECKS=0;

truncate street_stall.administrators;
truncate street_stall.evaluations;
truncate street_stall.locations;
truncate street_stall.merchants;
truncate street_stall.orders;
truncate street_stall.places;
truncate street_stall.questions;
truncate street_stall.users;
truncate street_stall.visitors;

-- 启动外键约束
SET FOREIGN_KEY_CHECKS=1;