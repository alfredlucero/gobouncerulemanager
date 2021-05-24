-- Steps to load this up assuming you are in the project's root directory and have a local MySQL installed.
-- mysql -u <user> -p
-- Type in password to enter MySQL CLI...
-- source db/db.sql
-- SELECT * FROM bounce_rule;
DROP DATABASE bouncerulemanager;

CREATE DATABASE bouncerulemanager CHARACTER SET utf8;

USE bouncerulemanager;

CREATE TABLE bounce_rule (
  id SMALLINT NOT NULL AUTO_INCREMENT,
  response_code SMALLINT NOT NULL DEFAULT 0,
  enhanced_code VARCHAR(16) NOT NULL,
  regex VARCHAR(255) NOT NULL,
  priority TINYINT NOT NULL DEFAULT 0,
  description VARCHAR(255) NOT NULL,
  bounce_action VARCHAR(255) NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE bounce_rule_change (
  id SMALLINT NOT NULL AUTO_INCREMENT,
  action VARCHAR(16) NOT NULL,
  bounce_rule_id SMALLINT NOT NULL,
  response_code SMALLINT NOT NULL DEFAULT 0,
  enhanced_code VARCHAR(16) NOT NULL,
  regex VARCHAR(255) NOT NULL,
  priority TINYINT NOT NULL DEFAULT 0,
  description VARCHAR(255) NOT NULL,
  bounce_action VARCHAR(255) NOT NULL,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

-- FOREIGN KEY (bounce_rule_id) REFERENCES bounce_rule(id) ON DELETE CASCADE

-- SHOW TABLES;
-- DESCRIBE bounce_rule;
-- DESCRIBE bounce_rule_change;

-- Use triggers to update the bounce_rule_change table

DELIMITER //
CREATE TRIGGER add_bounce_rule_created_change AFTER INSERT ON bounce_rule
FOR EACH ROW
BEGIN
  INSERT INTO bounce_rule_change (action, bounce_rule_id, response_code, enhanced_code, regex, priority, description, bounce_action)
  VALUES('created', NEW.id, NEW.response_code, NEW.enhanced_code, NEW.regex, NEW.priority, NEW.description, NEW.bounce_action);
END//
DELIMITER ;

DELIMITER //
CREATE TRIGGER add_bounce_rule_updated_change AFTER UPDATE ON bounce_rule
FOR EACH ROW
BEGIN
  INSERT INTO bounce_rule_change (action, bounce_rule_id, response_code, enhanced_code, regex, priority, description, bounce_action)
  VALUES('updated', NEW.id, NEW.response_code, NEW.enhanced_code, NEW.regex, NEW.priority, NEW.description, NEW.bounce_action);
END//
DELIMITER ;

DELIMITER //
CREATE TRIGGER add_bounce_rule_deleted_change AFTER DELETE ON bounce_rule
FOR EACH ROW
BEGIN
  INSERT INTO bounce_rule_change (action, bounce_rule_id, response_code, enhanced_code, regex, priority, description, bounce_action)
  VALUES('deleted', OLD.id, OLD.response_code, OLD.enhanced_code, OLD.regex, OLD.priority, OLD.description, OLD.bounce_action);
END//
DELIMITER ;

-- SHOW TRIGGERS;

INSERT INTO bounce_rule (response_code, enhanced_code, regex, priority, description, bounce_action)
VALUES (500, '4.7.1', 'some 500 4.7.1 regex', 0, 'some description about 500 4.7.1', 'suppress');
INSERT INTO bounce_rule (response_code, enhanced_code, regex, priority, description, bounce_action)
VALUES (450, '4.7.2', 'some 450 4.7.2 regex', 0, 'some description about 450 4.7.2', 'retry');
INSERT INTO bounce_rule (response_code, enhanced_code, regex, priority, description, bounce_action)
VALUES (501, '4.7.3', 'some 501 4.7.3 regex', 0, 'some description about 501 4.7.3', 'no_action');
INSERT INTO bounce_rule (response_code, enhanced_code, regex, priority, description, bounce_action)
VALUES (475, '5.0.1', 'some 475 5.0.1 regex', 0, 'some description about 475 5.0.1', 'suppress');

UPDATE bounce_rule 
SET response_code = 502,
    enhanced_code = '4.7.3',
    regex = 'some 502 4.7.3 regex',
    priority = 1,
    description = 'some description about 502 4.7.3',
    bounce_action = 'suppress'
WHERE id = 3;

-- SELECT * FROM bounce_rule;
DELETE FROM bounce_rule WHERE id = 4;

-- INSERT INTO bounce_rule_change (action, bounce_rule_id, response_code, enhanced_code, regex, priority, description, bounce_action)
-- VALUES ('created', 1, 500, '4.7.1', 'some 500 4.7.1 regex', 0, 'some description about 500 4.7.1', 'suppress');
-- INSERT INTO bounce_rule_change (bounce_rule_id, response_code, enhanced_code, regex, priority, description, bounce_action)
-- VALUES ('created', 2, 450, '4.7.2', 'some 450 4.7.2 regex', 0, 'some description about 450 4.7.2', 'retry');
-- INSERT INTO bounce_rule_change (bounce_rule_id, response_code, enhanced_code, regex, priority, description, bounce_action)
-- VALUES ('created', 3, 501, '4.7.3', 'some 501 4.7.3 regex', 0, 'some description about 501 4.7.3', 'no_action');
-- INSERT INTO bounce_rule_change (bounce_rule_id, response_code, enhanced_code, regex, priority, description, bounce_action)
-- VALUES ('created', 4, 475, '5.0.1', 'some 475 5.0.1 regex', 0, 'some description about 475 5.0.1', 'suppress');


CREATE TABLE throughput_rule (
  id INT(10) NOT NULL AUTO_INCREMENT,
  mx_domain VARCHAR(255) NOT NULL UNIQUE,
  max_connections INT(11) NOT NULL,
  messages_per_connection INT(11) NOT NULL,
  connection_ttl_millis INT(11) NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE throughput_rule_change (
  id INT NOT NULL AUTO_INCREMENT,
  action VARCHAR(16) NOT NULL,
  throughput_rule_id INT(10) NOT NULL,
  mx_domain VARCHAR(255) NOT NULL,
  max_connections INT(11) NOT NULL,
  messages_per_connection INT(11) NOT NULL,
  connection_ttl_millis INT(11) NOT NULL,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (throughput_rule_id) REFERENCES throughput_rule(id) ON DELETE CASCADE
);

-- SHOW TABLES;
-- DESCRIBE throughput_rule;
-- DESCRIBE throughput_rule_change;

-- Ensure both throughput_rule and throughput_rule_change inserts happen together through transactions

START TRANSACTION;
INSERT INTO throughput_rule (mx_domain, max_connections, messages_per_connection, connection_ttl_millis)
  VALUES('example.com', 36, 50, 0);
INSERT INTO throughput_rule_change (action, throughput_rule_id, mx_domain, max_connections, messages_per_connection, connection_ttl_millis) 
  VALUES('created', LAST_INSERT_ID(), 'example.com', 36, 50, 0);
COMMIT;

START TRANSACTION;
INSERT INTO throughput_rule (mx_domain, max_connections, messages_per_connection, connection_ttl_millis)
  VALUES('somemx.net', 100, 100, 15);
INSERT INTO throughput_rule_change (action, throughput_rule_id, mx_domain, max_connections, messages_per_connection, connection_ttl_millis) 
  VALUES('created', LAST_INSERT_ID(), 'somemx.net', 100, 100, 15);
COMMIT;
