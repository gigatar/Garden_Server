DROP TABLE IF EXISTS `controllers`;
CREATE TABLE `controllers` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT_NULL,
  `serial` VARCHAR(255) NOT NULL,
  `api_key` VARCHAR(32) NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID` (`ID`),
  UNIQUE KEY `serial` (`serial`)
  );

DROP TABLE IF EXISTS `sensor_data`;
CREATE TABLE `sensor_data` (
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `controller_id` int(11) NOT_NULL,
  `moisture` SMALLINT(11) NOT_NULL,
  `temperature` SMALLINT(3) NOT_NULL,
  `humidity` TINYint(3) NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID` (`ID`),
  );

  DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(20) NOT NULL,
  `password` varchar(164) NOT NULL,
  PRIMARY KEY (`ID`),
  UNIQUE KEY `ID` (`ID`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;