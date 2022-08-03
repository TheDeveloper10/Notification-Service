CREATE TABLE `Notifications`(
    `Id`          INTEGER       PRIMARY KEY AUTO_INCREMENT,
    `Title`       VARCHAR(128)  NOT NULL,
    `ContactType` TINYINT       NOT NULL,
    `ContactInfo` VARCHAR(128)  NOT NULL,
    `Message`     VARCHAR(2048) NOT NULL,
    `UserId`      VARCHAR(64)   NOT NULL,
    `AppId`       VARCHAR(64)   NOT NULL,
    `SentTime`    INTEGER       NOT NULL DEFAULT(UNIX_TIMESTAMP())
);

CREATE TABLE `Templates`(
    `Id`          INTEGER       PRIMARY KEY AUTO_INCREMENT,
    `ContactType` TINYINT       NOT NULL,
	`Template`    VARCHAR(2048) NOT NULL
);