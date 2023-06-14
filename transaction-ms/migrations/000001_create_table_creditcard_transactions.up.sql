CREATE TABLE `creditcard_transactions` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `transaction_id` VARCHAR(26) NOT NULL,
  `card_number` VARCHAR(16) NOT NULL,
  `amount` BIGINT NOT NULL,
  `installments` TINYINT NOT NULL,
  `description` TEXT NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`id`),
  INDEX `transaction_id_idx` (`transaction_id`)
) ENGINE = InnoDB;
