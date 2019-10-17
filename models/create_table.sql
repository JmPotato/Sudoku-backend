CREATE TABLE IF NOT EXISTS `user` (
    `uid` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `type` TINYINT UNSIGNED NOT NULL,
    `username` VARCHAR(20) NOT NULL,
    `authentication` CHAR(32) NOT NULL,
    `score` INT UNSIGNED,
    `submited` INT UNSIGNED,
    `passed` INT UNSIGNED,
    `created_time` TIMESTAMP,
    PRIMARY KEY ( `uid` )
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `puzzle` (
    `pid` INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `type` TINYINT UNSIGNED NOT NULL,
    `content` CHAR(81) NOT NULL,
    `descriptor` CHAR(32) NOT NULL,
    `level` TINYINT UNSIGNED,
    `submited` INT UNSIGNED,
    `passed` INT UNSIGNED,
    `created_time` TIMESTAMP,
    PRIMARY KEY ( `pid` )
) ENGINE=InnoDB DEFAULT CHARSET=utf8;