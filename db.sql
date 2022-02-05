-- Adminer 4.8.1 MySQL 5.5.5-10.6.4-MariaDB-1:10.6.4+maria~focal dump

SET NAMES utf8;
SET time_zone = '+02:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP DATABASE IF EXISTS `gomicro`;
CREATE DATABASE `gomicro` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `gomicro`;

DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `username` longtext DEFAULT NULL UNIQUE,
  `mail` longtext DEFAULT NULL UNIQUE,
  `password` longtext DEFAULT NULL,
  `balance` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_accounts_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `ads`;
CREATE TABLE `ads` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` longtext DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  `price` double DEFAULT NULL,
  `picture` longtext DEFAULT NULL,
  `available` tinyint(1) DEFAULT NULL,
  `poster_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_ads_deleted_at` (`deleted_at`),
  KEY `fk_ads_poster` (`poster_id`),
  CONSTRAINT `fk_ads_poster` FOREIGN KEY (`poster_id`) REFERENCES `accounts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `message` longtext DEFAULT NULL,
  `sender_id` bigint(20) unsigned DEFAULT NULL,
  `transaction_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_messages_deleted_at` (`deleted_at`),
  KEY `fk_messages_sender` (`sender_id`),
  KEY `fk_messages_transaction` (`transaction_id`),
  CONSTRAINT `fk_messages_sender` FOREIGN KEY (`sender_id`) REFERENCES `accounts` (`id`),
  CONSTRAINT `fk_messages_transaction` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `price` double DEFAULT NULL,
  `status` longtext DEFAULT NULL,
  `poster_id` bigint(20) unsigned DEFAULT NULL,
  `buyer_id` bigint(20) unsigned DEFAULT NULL,
  `ad_id` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_transactions_deleted_at` (`deleted_at`),
  KEY `fk_transactions_poster` (`poster_id`),
  KEY `fk_transactions_ad` (`ad_id`),
  KEY `fk_transactions_buyer` (`buyer_id`),
  CONSTRAINT `fk_transactions_ad` FOREIGN KEY (`ad_id`) REFERENCES `ads` (`id`),
  CONSTRAINT `fk_transactions_buyer` FOREIGN KEY (`buyer_id`) REFERENCES `accounts` (`id`),
  CONSTRAINT `fk_transactions_poster` FOREIGN KEY (`poster_id`) REFERENCES `accounts` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


-- 2021-09-29 16:12:51