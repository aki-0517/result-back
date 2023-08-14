CREATE TABLE `players` (
    `player_id` varchar(26) COLLATE utf8mb4_bin NOT NULL,
    `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `sail_no` tinyint DEFAULT '0',
    `sex` varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `icon_url` varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `point` tinyint DEFAULT '0',
    `comment` varchar(255) COLLATE utf8mb4_bin
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `admins` (
    `admin_id` varchar(26) COLLATE utf8mb4_bin NOT NULL,
    `name` varchar(255) COLLATE utf8mb4_bin NOT NULL,
    `hashed_password` varchar(255) COLLATE utf8mb4_bin DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

ALTER TABLE `players`
ADD PRIMARY KEY (`player_id`);

ALTER TABLE `admins`
ADD PRIMARY KEY (`admin_id`);