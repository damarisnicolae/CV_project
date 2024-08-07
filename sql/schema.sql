DROP TABLE IF EXISTS user;
DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id                INT AUTO_INCREMENT  PRIMARY KEY,
  jobtitle          VARCHAR(128) NOT NULL,
  firstname         VARCHAR(128) NOT NULL,
  lastname          VARCHAR(128) NOT NULL,
  email             VARCHAR(128) NOT NULL,
  phone             VARCHAR(128) NOT NULL,
  address           VARCHAR(128) NOT NULL,
  city              VARCHAR(128) NOT NULL,
  country           VARCHAR(128) NOT NULL,
  postalcode        VARCHAR(128) NOT NULL,
  dateofbirth       DATE NOT NULL,
  nationality       VARCHAR(128) NOT NULL,
  summary           VARCHAR(256),
  workexperience    VARCHAR(128),
  education         VARCHAR(128),
  skills            VARCHAR(128),
  languages         VARCHAR(128)
);

CREATE TABLE template (
  id            INT AUTO_INCREMENT  PRIMARY KEY,
  path          VARCHAR(128) NOT NULL
);
