DROP SCHEMA IF EXISTS sample;
-- CREATE SCHEMA sample;
-- USE sample;
USE auth;
DROP TABLE IF EXISTS users;

CREATE TABLE users
(
    id INT(11) NOT NULL
    AUTO_INCREMENT,
    -- NOT NULL IDENTITY(1, 1),
    -- NOT NULL AUTO_INCREMENT, このコードが使用できなかったため代わりにIDENTITYを使用した
    icon_url VARCHAR
    (500),
    google_name VARCHAR
    (255),
    access_token VARCHAR
    (255),
    PRIMARY KEY
    (id)
    --     created_at DATETIME  default current_timestamp,
    --     updated_at TIMESTAMP default current_timestamp on update current_timestamp,
    --     deleted_at TIMESTAMP default current_timestamp on update current_timestamp
);

INSERT INTO users (id,icon_url,google_name,access_token) VALUES (1018000,'aa','aa','aa');
INSERT INTO users (id,icon_url,google_name,access_token) VALUES (1018085,'taketo','wakamatsu','takematsu');
--     INSERT INTO users
--         (id,icon_url,google_name,access_token)
--     VALUES
--         (1018000, "https://lh3.googleusercontent.com/a-/AAuE7mDbifKYLqkI_JjBBneLJRjgLBm5R3JDJClyDWHV=s96-c", "戸澤涼", "eyJhbGciOiJSUzI1NiIsImtpZCI6IjU3YjE5MjhmMmY2MzMyOWYyZTkyZjRmMjc4Zjk0ZWUxMDM4YzkyM2MiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwiYXpwIjoiNDMxMjk0OTY4MjgtaDRlbWlnZW1qaTFhbHE0OXYxb3FmM3ZpZmNhOXNpNjMuYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiI0MzEyOTQ5NjgyOC1oNGVtaWdlbWppMWFscTQ5djFvcWYzdmlmY2E5c2k2My5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsInN1YiI6IjExNDg3OTIzMjE4MTIwNDM0Mzc3OSIsImVtYWlsIjoic2l0aXRvdTcwQGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJhdF9oYXNoIjoiLUZGcTk2MmVaQ1U5alV5eXZPSzdHQSIsIm5hbWUiOiLpm6vnn7PljZPogLYiLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EtL0FBdUU3bUJNQkZqTFl5QlJ3RlVUQUVNTnRBeGdaMWI0czNGdGFmR2tCOC03a3c9czk2LWMiLCJnaXZlbl9uYW1lIjoi5Y2T6IC2IiwiZmFtaWx5X25hbWUiOiLpm6vnn7MiLCJsb2NhbGUiOiJqYSIsImlhdCI6MTU3NzA4MTU4MiwiZXhwIjoxNTc3MDg1MTgyLCJqdGkiOiIyN2EyMTVmMDM3ZTFlZDg2YjVhOTkxY2U1NDRjMjY5NzgwNTdhN2M4In0.RTHVPEvlOL7GaqT0x0bOO1grfyYfDt1ohXVyJdE7URXSQFrXSENUuYS9GnSsmiWgIKBO-7W2T7QQsTgqeS8T_39GpSyP4YAx03XVp0izShT0MRar2QaT26g983A_dI2yl-xpivIfKHUvQZcYIOPcsgG6w1ubMA8kldXoHlrzYVywAF9LIS2Uh8p3jDRK4S_jGAXhDSKKEUD54Xc2RIWDl5sWA035MYf7GHSOhFWJd7Fjeml-DkXOsZ2O9Ll0pUugZHae6OcYAc_tVp0nd3mUc19Hso22JVPNU4VyvyOAGGJGIQGX5GsTXgpYnoZWMhLal7UyMy6reAH9EmsdPV_1ow");
