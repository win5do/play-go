DROP TABLE IF EXISTS not_null;
CREATE TABLE not_null
(
  id     INT PRIMARY KEY AUTO_INCREMENT,
  name   VARCHAR(50) NOT NULL,
  number INT(10)
);
CREATE INDEX allow_null_key_index
  ON not_null (name);

DROP TABLE IF EXISTS allow_null;
CREATE TABLE allow_null
(
  id     INT PRIMARY KEY AUTO_INCREMENT,
  name   VARCHAR(50),
  number INT(10)
);
CREATE INDEX allow_null_name_index
  ON allow_null (name);