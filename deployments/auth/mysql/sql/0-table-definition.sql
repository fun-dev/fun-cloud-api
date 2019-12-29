USE auth;
CREATE TABLE users (
    id INT(11) AUTO_INCREMENT NOT NULL,
    icon_url VARCHAR(500),
    google_name VARCHAR(255),
    access_token VARCHAR(255),
    PRIMARY KEY(id),
--     created_at DATETIME  default current_timestamp,
 dt DATETIME_INTERVAL_CODE DEFAULT CURRENT_TIMESTAMP ,
--  ON UPDATE CURRENT_TIMESTAMP,

--     updated_at TIMESTAMP default current_timestamp on update current_timestamp,
 ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP

--     deleted_at TIMESTAMP default current_timestamp on update current_timestamp
DELETE FROM users WHERE timestamp < (NOW() - INTERVAL 10 MINUTE)
);