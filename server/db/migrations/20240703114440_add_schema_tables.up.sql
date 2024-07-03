CREATE TABLE user(
    user_id INT(8) UNIQUE NOT NULL AUTO_INCREMENT,
    user_name VARCHAR(16) NOT NULL,
    email VARCHAR(255) ,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (email)
);

CREATE TABLE server(
    server_id INT(8) AUTO_INCREMENT,
    server_name VARCHAR(16) NOT NULL,
    PRIMARY KEY (server_id)
);

CREATE TABLE joins(
    user_id INT(8),
    server_id INT(8),
    PRIMARY KEY (user_id, server_id),
    FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE,
    FOREIGN KEY (server_id) REFERENCES server(server_id) ON DELETE CASCADE
);

CREATE TABLE owns(
    user_id INT(8),
    server_id INT(8),
    PRIMARY KEY (user_id, server_id),
    FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE,
    FOREIGN KEY (server_id) REFERENCES server(server_id) ON DELETE CASCADE
);

CREATE TABLE channel(
    channel_id INT(8) AUTO_INCREMENT,
    channel_name VARCHAR(16) NOT NULL,
    PRIMARY KEY (channel_id)
);

CREATE TABLE has(
    channel_id INT(8),
    server_id INT(8),
    PRIMARY KEY (channel_id, server_id),
    FOREIGN KEY (server_id) REFERENCES server(server_id) ON DELETE CASCADE,
    FOREIGN KEY (channel_id) REFERENCES channel(channel_id) ON DELETE CASCADE
);

CREATE TABLE chat(
    channel_id INT(8),
    user_id INT(8),
    time TIMESTAMP,
    content VARCHAR(1023) NOT NULL,
    PRIMARY KEY (channel_id, user_id, time),
    FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE RESTRICT,
    FOREIGN KEY (channel_id) REFERENCES channel(channel_id) ON DELETE CASCADE
);