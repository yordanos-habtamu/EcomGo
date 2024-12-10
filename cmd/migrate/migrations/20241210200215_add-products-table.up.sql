CREATE TABLE IF NOT EXISTS products (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,           -- Product ID (Primary Key)
  `name` VARCHAR(255) NOT NULL,                        -- Product name
  `description` TEXT,                                  -- Product description
  `price` DECIMAL(10, 2) NOT NULL,                     -- Product price
  `stock_quantity` INT NOT NULL DEFAULT 0,             -- Quantity available in stock
  `category` VARCHAR(255),                             -- Category the product belongs to
  `image_url` VARCHAR(255),                            -- URL to the product image
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp when product is created
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Timestamp when product is last updated
  `is_active` BOOLEAN NOT NULL DEFAULT TRUE,           -- Indicates whether the product is active (available for sale)
  PRIMARY KEY(`id`)                                    -- Primary key for the table
);
