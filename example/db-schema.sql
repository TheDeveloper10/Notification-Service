CREATE DATABASE IF NOT EXISTS TestDatabase;

USE TestDatabase;

CREATE TABLE IF NOT EXISTS `Templates`(
    `Id`        INTEGER       PRIMARY KEY AUTO_INCREMENT,
    `EmailBody` VARCHAR(2048),
    `SMSBody`   VARCHAR(2048),
    `PushBody`  VARCHAR(2048),
	`Language`  VARCHAR(3)    NOT NULL,
	`Type`      VARCHAR(8)    NOT NULL
);

CREATE TABLE IF NOT EXISTS `Notifications`(
    `Id`                   INTEGER       PRIMARY KEY AUTO_INCREMENT,
    `AppId`                VARCHAR(16)   NOT NULL,
    `TemplateId`           INTEGER       NOT NULL,
    `ContactInfo`          VARCHAR(168),
    `Title`                VARCHAR(128)  NOT NULL,
    `Message`              VARCHAR(2048) NOT NULL,
    `SentTime`             INTEGER       NOT NULL DEFAULT(UNIX_TIMESTAMP())
);

CREATE TABLE IF NOT EXISTS `Clients`(
    `Id`           VARCHAR(16)  PRIMARY KEY,
    `Secret`       VARCHAR(128) NOT NULL,
    `Permissions`  INTEGER      NOT NULL,
    `CreationTime` INTEGER      NOT NULL DEFAULT(UNIX_TIMESTAMP())
);
