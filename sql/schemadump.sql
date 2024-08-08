DROP TABLE IF EXISTS user;
CREATE TABLE user (
  id                INT AUTO_INCREMENT  PRIMARY KEY,
  jobtitle          VARCHAR(128) NOT NULL,
  firstname         VARCHAR(128) NOT NULL,
  lastname          VARCHAR(128) NOT NULL,
  email             VARCHAR(128) NOT NULL,
  phone             VARCHAR(128) NOT NULL,
  address           VARCHAR(128) NOT NULL, -- Corrected spelling
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

INSERT INTO user
  (jobtitle, firstname, lastname, email, phone, address, city, country, postalcode, dateofbirth, nationality)
VALUES
  ('Inginer Software', 'Anne', 'Ungurean', 'anne_ungurean@yahoo.com', '0727999999', 'Str. Lalelelor, nr.17', 'București', 'Romania', '010001', '1990-11-11', 'română');

INSERT INTO template
  (path)
VALUES
  ('./template1.html'),
  ('./template2.html'),
  ('./template3.html');
