CREATE TABLE `creditcard_transactions_status` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `creditcard_transaction_id` INT NOT NULL,
    `status` VARCHAR(20) NOT NULL,
    `created_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    CONSTRAINT `creditcard_transaction_fk` FOREIGN KEY (`creditcard_transaction_id`) REFERENCES `creditcard_transactions`(`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB;
