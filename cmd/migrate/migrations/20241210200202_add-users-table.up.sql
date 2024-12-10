CREATE TABLE IF NOT EXISTS users (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `firstName` VARCHAR(255) NOT NULL,
  `lastName` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `DoB` Date NOT NULL,
  `sex` ENUM('male', 'female') NOT NULL,
  `role` ENUM('admin', 'customer', 'seller') NOT NULL DEFAULT 'customer',
  `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(`id`),
  UNIQUE KEY(`email`)
);
