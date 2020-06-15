DROP SCHEMA IF EXISTS app;
CREATE SCHEMA app;
USE app;

DROP TABLE IF EXISTS user;

CREATE TABLE user
(
  id           INT(10),
  name     VARCHAR(40)
);