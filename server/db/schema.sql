
DROP DATABASE IF EXISTS `pathLogic`;
CREATE DATABASE `pathLogic`;

USE `pathLogic`;

/*Table structure for table `site` */

DROP TABLE IF EXISTS `sites`;

CREATE TABLE `site` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE SITES (
  id SERIAL PRIMARY KEY,
  name text NOT NULL
);
CREATE INDEX sites_name_index ON sites(id);
INSERT INTO SITES (name) VALUES ('poke bowl wicker park');
INSERT INTO SITES (name) VALUES ('poke bowl river north');

DROP TABLE IF EXISTS `inventory`;
CREATE TABLE `inventory` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `site_id` bigint NOT NULL AUTO_INCREMENT,
  `item_id` bigint NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE INDEX inventory_site_id_index ON inventory(site_id) USING BTREE;
CREATE INDEX inventory_item_id_index ON inventory(item_id) USING BTREE;

DROP TABLE IF EXISTS `items`;

CREATE TABLE `items` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `type` varchar NOT NULL,
  `name` varchar NOT NULL,
  `description` varchar NOT NULL,
  `history` varchar NOT NULL,
  `status` varchar NOT NULL,
  `category_id` varchar NOT NULL,
  `price_id` varchar NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE INDEX items_type_index ON items(type) USING BTREE;

DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE INDEX category_name_index ON category(name) USING BTREE;

DROP TABLE IF EXISTS `price`;
CREATE TABLE `price` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `price` FLOAT NOT NULL,
  `currency` varchar NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
CREATE INDEX price_price_index ON price(price) USING BTREE;

ALTER TABLE `inventory` ADD CONSTRAINT `inventory_fk0` FOREIGN KEY (`site_id`) REFERENCES `sites`(`id`);

ALTER TABLE `inventory` ADD CONSTRAINT `inventory_fk1` FOREIGN KEY (`item_id`) REFERENCES `items`(`id`);

ALTER TABLE `items` ADD CONSTRAINT `item_fk0` FOREIGN KEY (`category_id`) REFERENCES `category`(`id`);

ALTER TABLE `items` ADD CONSTRAINT `item_fk1` FOREIGN KEY (`price_id`) REFERENCES `price`(`id`);
