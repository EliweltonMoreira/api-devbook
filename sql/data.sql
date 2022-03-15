INSERT INTO users (name, nick, email, password)
VALUES
('User 1', 'user_1', 'user1@gmail.com', '$2a$10$oJAm44gxd9QiEKDO41rgvuubFNijmv5ze2vtgG3Vq3wdKANUNDORy'),
('User 2', 'user_2', 'user2@gmail.com', '$2a$10$oJAm44gxd9QiEKDO41rgvuubFNijmv5ze2vtgG3Vq3wdKANUNDORy'),
('User 3', 'user_3', 'user3@gmail.com', '$2a$10$oJAm44gxd9QiEKDO41rgvuubFNijmv5ze2vtgG3Vq3wdKANUNDORy');

INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(3, 1),
(1, 3);