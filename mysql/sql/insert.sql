INSERT INTO users (name, password) VALUES ('taro', 'password');
INSERT INTO users (name, password) VALUES ('hanako', 'PASSWORD');

INSERT INTO servers (name, owner_id, icon) VALUES ("server1", 1, 'hoge.png');

INSERT INTO channels (server_id, name) VALUES (1, 'channel1');
INSERT INTO channels (server_id, name) VALUES (1, 'channel2');

INSERT INTO posts (user_id, channel_id, content) VALUES (1, 1,  '質問1\n改行');
INSERT INTO posts (user_id, channel_id, content) VALUES (1, 1, '質問2\n改行');
