INSERT INTO users (name, password) VALUES ('taro', 'password');
INSERT INTO users (name, password) VALUES ('hanako', 'PASSWORD');

INSERT INTO servers (name, owner, icon) VALUES ("server1", 'manager', 'hoge.png');

INSERT INTO channels (server_id, name) VALUES (1, 'channel1');
INSERT INTO channels (server_id, name) VALUES (1, 'channel2');

INSERT INTO posts (user_id, channel_id, title, body) VALUES (1,1, 'test1', '質問1\n改行');
INSERT INTO posts (user_id, channel_id, title, body) VALUES (1,1, 'test2', '質問2\n改行');