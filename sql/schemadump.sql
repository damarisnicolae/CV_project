DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id                INT IDENTITY(1,1) PRIMARY KEY,
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
  summary           VARCHAR(256) NULL,
  workexperience    VARCHAR(128) NULL,
  education         VARCHAR(128) NULL,
  skills            VARCHAR(128) NULL,
  languages         VARCHAR(128) NULL
);

CREATE TABLE template (
  id            INT IDENTITY(1,1) PRIMARY KEY,
  path          VARCHAR(128) NOT NULL
);

INSERT INTO users
  (jobtitle, firstname, lastname, email, phone, address, city, country, postalcode, dateofbirth, nationality)
VALUES
  ('Inginer Software', 'Anne', 'Ungurean', 'anne_ungurean@yahoo.com', '0727999999', 'Str. Lalelelor, nr.17', 'București', 'Romania', '010001', '1990-11-11', 'română');

INSERT INTO template
  (path)
VALUES
  ('./template1.html'),
  ('./template2.html'),
  ('./template3.html');
