-- Active: 1690423316021@@127.0.0.1@3307@bootcamp
use bootcamp;

CREATE TABLE `user` (
  `user_id` char(36) PRIMARY KEY NOT NULL,
  `username` varchar(50) NOT NULL,
  `email` varchar(100) NOT NULL,
  `role` ENUM ('regular', 'admin') NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `updated_by` char(36) NULL DEFAULT NULL,
  `deleted_by` char(36) NULL DEFAULT NULL
);


CREATE TABLE `brand` (
  `brand_id` char(36) PRIMARY KEY NOT NULL,
  `brand_name` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `updated_by` char(36) NULL DEFAULT NULL,
  `deleted_by` char(36) NULL DEFAULT NULL,
  `created_by` char(36) NULL DEFAULT NULL
);


CREATE TABLE `product` (
  `product_id` char(36) PRIMARY KEY NOT NULL,
  `user_id` char(36) NOT NULL,
  `brand_id` char(36) NOT NULL,
  `product_name` varchar(200) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `updated_by` char(36) NULL DEFAULT NULL,
  `deleted_by` char(36) NULL DEFAULT NULL,
  `created_by` char(36) NULL DEFAULT NULL
);

CREATE TABLE `variant` (
  `variant_id` char(36) PRIMARY KEY NOT NULL,
  `product_id` char(36) NOT NULL,
  `variant_name` varchar(100) NOT NULL,
  `price` decimal(10,2) NOT NULL,
  `status` ENUM ('ready', 'out_of_stock', 'limited') NOT NULL,
  `quantity` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `updated_by` char(36) NULL DEFAULT NULL,
  `deleted_by` char(36) NULL DEFAULT NULL,
  `created_by` char(36) NULL DEFAULT NULL
);

CREATE TABLE `image` (
  `image_id` char(36) PRIMARY KEY NOT NULL,
  `variant_id` char(36) NOT NULL,
  `image_url` varchar(200) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `updated_by` char(36) NULL DEFAULT NULL,
  `deleted_by` char(36) NULL DEFAULT NULL,
  `created_by` char(36) NULL DEFAULT NULL
);

CREATE TABLE `warehouse` (
  `warehouse_id` int AUTO_INCREMENT PRIMARY KEY NOT NULL,
  `warehouse_name` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `updated_by` char(36) NULL DEFAULT NULL,
  `deleted_by` char(36) NULL DEFAULT NULL,
  `created_by` char(36) NULL DEFAULT NULL
);

CREATE Table variant_location (
	variant_location_id int AUTO_INCREMENT PRIMARY KEY NOT NULL,
	warehouse_id char(36) NOT NULL,
	variant_id char(36) NOT NULL,
	variant_quantity int NOT NULL
);
  
CREATE INDEX idx_username_and_email ON `user` (`username`,`email`);
CREATE INDEX idx_user_role ON `user` (`role`);
CREATE INDEX `idx_variant_images` ON `image` (`variant_id`, `image_id`, `image_url`);
ALTER TABLE `product` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`);
ALTER TABLE `product` ADD FOREIGN KEY (`brand_id`) REFERENCES `brand` (`brand_id`);
ALTER TABLE `variant` ADD FOREIGN KEY (`product_id`) REFERENCES `product` (`product_id`);
ALTER TABLE `image` ADD FOREIGN KEY (`variant_id`) REFERENCES `variant` (`variant_id`);
ALTER TABLE variant_location ADD FOREIGN KEY (`warehouse_id`) REFERENCES `warehouse` (`warehouse_id`);
ALTER TABLE variant_location ADD FOREIGN KEY (`variant_id`) REFERENCES `variant` (`variant_id`);
ALTER TABLE `user` ADD CONSTRAINT `user_email_and_username` UNIQUE (`username`,`email`);

CREATE TRIGGER variant_update
AFTER UPDATE ON variant
FOR EACH ROW
	UPDATE product
    SET updated_at = CURRENT_TIMESTAMP
    WHERE product.product_id = NEW.product_id;
DELIMITER //
CREATE TRIGGER image_update
AFTER UPDATE ON `image`
FOR EACH ROW
  UPDATE variant
    SET updated_at = CURRENT_TIMESTAMP
    WHERE variant.variant_id = NEW.variant_id;
DELIMITER //
CREATE TRIGGER user_update
AFTER UPDATE ON `user`
FOR EACH ROW
  UPDATE `user`
    SET updated_at = CURRENT_TIMESTAMP
    WHERE user_id = NEW.user_id;
DELIMITER //
CREATE TRIGGER product_update
AFTER UPDATE ON `product`
FOR EACH ROW
  UPDATE `product`
    SET updated_at = CURRENT_TIMESTAMP
    WHERE product_id = NEW.product_id;
DELIMITER

DELIMITER //

CREATE TRIGGER tr_insert_variant_location
AFTER INSERT ON variant
FOR EACH ROW
BEGIN
    DECLARE totalWarehouses INT;
    DECLARE selectedWarehouse INT;
    SET totalWarehouses = (SELECT COUNT(*) FROM warehouse); -- Replace 'warehouse' with your actual warehouse table name.
    SET selectedWarehouse = FLOOR(1 + RAND() * totalWarehouses);

    INSERT INTO variant_location (variant_id, warehouse_id,variant_quantity)
    VALUES (NEW.variant_id, selectedWarehouse,NEW.quantity);
END;
//

DELIMITER ;