CREATE TABLE `records` (
  `id` int PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `release_date` varchar(50) NOT NULL,
  `image` varchar(255),
  `status` int NOT NULL DEFAULT 0,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `labels` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `label_catalogue_number` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `label_id` int,
  `catalogue_number` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `label_catalogue_number_to_record` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `label_catalogue_number_id` int,
  `record_id` int,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `record_labels` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `label_id` int,
  `record_id` int,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `genres` (
  `genre_id` int PRIMARY KEY AUTO_INCREMENT,
  `genre_name` varchar(255),
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `record_genres` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `record_id` int,
  `genre_id` int,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `artists` (
  `artist_id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `record_artists` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `artist_id` int,
  `record_id` int,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `record_images` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `record_id` int,
  `image` varchar(255),
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `tracklist` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `track` varchar(255) NOT NULL,
  `track_title` varchar(255) NOT NULL,
  `track_duration` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `record_tracklists` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `record_id` int,
  `tracklist_id` int,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `record_tracklists_artists` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `tracklist_id` int,
  `artist_id` int,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `format` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `quantity` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `record_format` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `format_id` int,
  `record_id` int,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `size` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `record_size` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `size_id` int,
  `record_id` int,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `speed` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `record_speed` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `speed_id` int,
  `record_id` int,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `description` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `record_description` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `description_id` int,
  `record_id` int,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `channels` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE TABLE `record_channels` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `channels_id` int,
  `record_id` int,
  `created_at` datetime NOT NULL DEFAULT (now()),
  `updated_at` datetime NOT NULL DEFAULT (now())
);

CREATE INDEX `records_index_0` ON `records` (`title`);

CREATE INDEX `labels_index_1` ON `labels` (`name`);

CREATE INDEX `label_catalogue_number_index_2` ON `label_catalogue_number` (`catalogue_number`);

CREATE INDEX `record_labels_index_3` ON `record_labels` (`record_id`);

CREATE INDEX `record_labels_index_4` ON `record_labels` (`label_id`);

CREATE INDEX `record_labels_index_5` ON `record_labels` (`record_id`, `label_id`);

CREATE INDEX `genres_index_6` ON `genres` (`genre_name`);

CREATE INDEX `record_genres_index_7` ON `record_genres` (`record_id`);

CREATE INDEX `record_genres_index_8` ON `record_genres` (`genre_id`);

CREATE INDEX `record_genres_index_9` ON `record_genres` (`record_id`, `genre_id`);

CREATE INDEX `artists_index_10` ON `artists` (`name`);

CREATE INDEX `record_artists_index_11` ON `record_artists` (`artist_id`);

CREATE INDEX `record_artists_index_12` ON `record_artists` (`record_id`);

CREATE INDEX `record_artists_index_13` ON `record_artists` (`artist_id`, `record_id`);

CREATE INDEX `record_images_index_14` ON `record_images` (`image`);

CREATE INDEX `tracklist_index_15` ON `tracklist` (`track`);

CREATE INDEX `record_tracklists_index_16` ON `record_tracklists` (`record_id`);

CREATE INDEX `record_tracklists_index_17` ON `record_tracklists` (`tracklist_id`);

CREATE INDEX `record_tracklists_index_18` ON `record_tracklists` (`record_id`, `tracklist_id`);

CREATE INDEX `record_tracklists_artists_index_19` ON `record_tracklists_artists` (`tracklist_id`);

CREATE INDEX `record_tracklists_artists_index_20` ON `record_tracklists_artists` (`artist_id`);

CREATE INDEX `record_tracklists_artists_index_21` ON `record_tracklists_artists` (`tracklist_id`, `artist_id`);

CREATE INDEX `format_index_22` ON `format` (`name`);

CREATE INDEX `record_format_index_23` ON `record_format` (`format_id`);

CREATE INDEX `record_format_index_24` ON `record_format` (`record_id`);

CREATE INDEX `record_format_index_25` ON `record_format` (`format_id`, `record_id`);

CREATE INDEX `size_index_26` ON `size` (`name`);

CREATE INDEX `record_size_index_27` ON `record_size` (`size_id`);

CREATE INDEX `record_size_index_28` ON `record_size` (`record_id`);

CREATE INDEX `record_size_index_29` ON `record_size` (`size_id`, `record_id`);

CREATE INDEX `speed_index_30` ON `speed` (`name`);

CREATE INDEX `record_speed_index_31` ON `record_speed` (`speed_id`);

CREATE INDEX `record_speed_index_32` ON `record_speed` (`record_id`);

CREATE INDEX `record_speed_index_33` ON `record_speed` (`speed_id`, `record_id`);

CREATE INDEX `description_index_34` ON `description` (`name`);

CREATE INDEX `record_description_index_35` ON `record_description` (`description_id`);

CREATE INDEX `record_description_index_36` ON `record_description` (`record_id`);

CREATE INDEX `record_description_index_37` ON `record_description` (`description_id`, `record_id`);

CREATE INDEX `channels_index_38` ON `channels` (`name`);

CREATE INDEX `record_channels_index_39` ON `record_channels` (`channels_id`);

CREATE INDEX `record_channels_index_40` ON `record_channels` (`record_id`);

CREATE INDEX `record_channels_index_41` ON `record_channels` (`channels_id`, `record_id`);

ALTER TABLE `label_catalogue_number` ADD FOREIGN KEY (`label_id`) REFERENCES `labels` (`id`);

ALTER TABLE `label_catalogue_number_to_record` ADD FOREIGN KEY (`label_catalogue_number_id`) REFERENCES `label_catalogue_number` (`id`);

ALTER TABLE `label_catalogue_number_to_record` ADD FOREIGN KEY (`record_id`) REFERENCES `records` (`id`);

ALTER TABLE `record_labels` ADD FOREIGN KEY (`label_id`) REFERENCES `labels` (`id`);

ALTER TABLE `record_labels` ADD FOREIGN KEY (`record_id`) REFERENCES `records` (`id`);

ALTER TABLE `record_genres` ADD FOREIGN KEY (`record_id`) REFERENCES `records` (`id`);

ALTER TABLE `record_genres` ADD FOREIGN KEY (`genre_id`) REFERENCES `genres` (`genre_id`);

ALTER TABLE `record_artists` ADD FOREIGN KEY (`artist_id`) REFERENCES `artists` (`artist_id`);

ALTER TABLE `record_artists` ADD FOREIGN KEY (`record_id`) REFERENCES `records` (`id`);

ALTER TABLE `record_images` ADD FOREIGN KEY (`record_id`) REFERENCES `records` (`id`);

ALTER TABLE `record_tracklists` ADD FOREIGN KEY (`record_id`) REFERENCES `records` (`id`);

ALTER TABLE `record_tracklists` ADD FOREIGN KEY (`tracklist_id`) REFERENCES `tracklist` (`id`);

ALTER TABLE `record_tracklists_artists` ADD FOREIGN KEY (`tracklist_id`) REFERENCES `tracklist` (`id`);

ALTER TABLE `record_tracklists_artists` ADD FOREIGN KEY (`artist_id`) REFERENCES `artists` (`artist_id`);

ALTER TABLE `record_format` ADD FOREIGN KEY (`format_id`) REFERENCES `format` (`id`);

ALTER TABLE `record_format` ADD FOREIGN KEY (`record_id`) REFERENCES `records` (`id`);

ALTER TABLE `record_size` ADD FOREIGN KEY (`size_id`) REFERENCES `size` (`id`);

ALTER TABLE `record_size` ADD FOREIGN KEY (`record_id`) REFERENCES `records` (`id`);

ALTER TABLE `record_speed` ADD FOREIGN KEY (`speed_id`) REFERENCES `speed` (`id`);

ALTER TABLE `record_speed` ADD FOREIGN KEY (`record_id`) REFERENCES `records` (`id`);

ALTER TABLE `record_description` ADD FOREIGN KEY (`description_id`) REFERENCES `description` (`id`);

ALTER TABLE `record_description` ADD FOREIGN KEY (`record_id`) REFERENCES `records` (`id`);

ALTER TABLE `record_channels` ADD FOREIGN KEY (`channels_id`) REFERENCES `description` (`id`);

ALTER TABLE `record_channels` ADD FOREIGN KEY (`record_id`) REFERENCES `records` (`id`);
