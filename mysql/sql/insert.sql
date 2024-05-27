INSERT INTO hello_worlds (lang, message) VALUES ('en', 'Hello World');
INSERT INTO hello_worlds (lang, message) VALUES ('ja', 'こんにちは 世界');

INSERT INTO users (name, password) VALUES ('taro', 'password');
INSERT INTO users (name, password) VALUES ('hanako', 'PASSWORD');

INSERT INTO posts (user_id, title, body) VALUES (1, 'test1', '質問1\n改行');
INSERT INTO posts (user_id, title, body) VALUES (1, 'test2', '質問2\n改行');

INSERT INTO posts (name, owner, icon) VALUES ("server1", 'manager', 'hoge');

INSERT INTO channels (server_id, name) VALUES (1, 'channel1');
INSERT INTO channels (server_id, name) VALUES (1, 'channel2');