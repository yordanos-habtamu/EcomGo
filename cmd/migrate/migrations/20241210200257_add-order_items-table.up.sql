CREATE TABLE IF NOT EXISTS orderItems (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,            -- Unique ID for each order item
  `order_id` INT UNSIGNED NOT NULL,                     -- The ID of the order (foreign key)
  `product_id` INT UNSIGNED NOT NULL,                   -- The ID of the product (foreign key)
  `quantity` INT NOT NULL DEFAULT 1,                     -- Quantity of the product in this order item
  `unit_price` DECIMAL(10, 2) NOT NULL,                  -- Price of a single unit of the product at the time of the order
  `total_price` DECIMAL(10, 2) NOT NULL,                 -- Total price for this product in the order (unit_price * quantity)
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the order item is created
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Timestamp for the last update
  PRIMARY KEY(`id`),                                    -- Primary key for the table
  FOREIGN KEY(`order_id`) REFERENCES orders(`id`),      -- Foreign key to the orders table
  FOREIGN KEY(`product_id`) REFERENCES products(`id`)   -- Foreign key to the products table
);

