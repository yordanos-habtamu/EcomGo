CREATE TABLE IF NOT EXISTS `order_items` (
    `id` INT UNSIGNED PRIMARY KEY NOT NULL,                      -- Auto-incrementing ID
    `product_id` INT UNSIGNED NOT NULL,                    -- Foreign key referencing the product
    `order_id` INT  UNSIGNED NOT NULL,                      -- Foreign key referencing the order
    `quantity` INT NOT NULL DEFAULT 1,            -- Quantity of the product in the order
    `created_at` TIMESTAMP NOT NULL DEFAULT NOW(),-- Timestamp for when the record was created
    `updated_at` TIMESTAMP NOT NULL DEFAULT NOW(),-- Timestamp for when the record was last updated
    CONSTRAINT `fk_product` FOREIGN KEY (`product_id`) REFERENCES products(`id`) ON DELETE CASCADE,
    CONSTRAINT `fk_order` FOREIGN KEY (`order_id`) REFERENCES orders(`id`) ON DELETE CASCADE
);
