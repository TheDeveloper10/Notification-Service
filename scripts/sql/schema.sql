CREATE DATABASE IF NOT EXISTS TestDatabase;

USE TestDatabase;

CREATE TABLE IF NOT EXISTS `Notifications`(
    `Id`                   INTEGER       PRIMARY KEY AUTO_INCREMENT,
    `AppId`                VARCHAR(16)   NOT NULL,
    `TemplateId`           INTEGER       NOT NULL,
    `ContactType`          VARCHAR(8)    NOT NULL,
    `ContactInfo`          VARCHAR(168),
    `Title`                VARCHAR(128)  NOT NULL,
    `Message`              VARCHAR(2048) NOT NULL,
    `SentTime`             INTEGER       NOT NULL DEFAULT(UNIX_TIMESTAMP())
);

CREATE TABLE IF NOT EXISTS `Templates`(
    `Id`          INTEGER       PRIMARY KEY AUTO_INCREMENT,
    `ContactType` VARCHAR(8)    NOT NULL,
	`Template`    VARCHAR(2048) NOT NULL,
	`Language`    VARCHAR(3)    NOT NULL,
	`Type`        VARCHAR(8)    NOT NULL
);