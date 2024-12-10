CREATE TABLE IF NOT EXISTS orders (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,             -- Unique ID for the order
  `user_id` INT UNSIGNED NOT NULL,                       -- ID of the user placing the order (foreign key)
  `status` ENUM('pending', 'processing', 'completed', 'shipped', 'cancelled') NOT NULL DEFAULT 'pending',  -- Order status
  `total_price` DECIMAL(10, 2) NOT NULL,                  -- Total price of the order
  `shipping_address` TEXT NOT NULL,                       -- Shipping address for the order
  `billing_address` TEXT NOT NULL,                        -- Billing address for the order
  `payment_method` ENUM('credit_card', 'paypal', 'bank_transfer') NOT NULL, -- Payment method
  `payment_status` ENUM('pending', 'paid', 'failed') NOT NULL DEFAULT 'pending',  -- Payment status
  `order_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- Date when the order was placed
  `shipped_date` TIMESTAMP DEFAULT NULL,                  -- Date when the order was shipped
  `delivery_date` TIMESTAMP DEFAULT NULL,                 -- Expected delivery date
  `tracking_number` VARCHAR(255),                         -- Tracking number for the shipment
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,  -- Timestamp when the order was created
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Timestamp when the order was last updated
  PRIMARY KEY(`id`),                                     -- Primary key for the table
  FOREIGN KEY(`user_id`) REFERENCES users(`id`)          -- Foreign key to the users table
);
