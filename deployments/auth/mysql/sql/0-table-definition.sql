USE auth;
CREATE TABLE users (
    id INT(11) AUTO_INCREMENT NOT NULL,
    icon_url VARCHAR(500),
    google_name VARCHAR(255),
    access_token VARCHAR(255),
    PRIMARY KEY(id)
--     created_at DATETIME  default current_timestamp,
--     updated_at TIMESTAMP default current_timestamp on update current_timestamp,
--     deleted_at TIMESTAMP default current_timestamp on update current_timestamp
);