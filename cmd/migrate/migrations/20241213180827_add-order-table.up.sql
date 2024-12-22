CREATE TABLE IF NOT EXISTS orders (
    `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY, -- Auto-incrementing ID
    `user_id` INT UNSIGNED NOT NULL,              -- Foreign key for the user placing the order
    `total` DECIMAL(10, 2) NOT NULL,              -- Total amount for the order
    `status` ENUM('pending','complete','failed') NOT NULL DEFAULT 'pending', -- Order status
    `address` TEXT NOT NULL,                      -- Shipping address
    `created_at` TIMESTAMP NOT NULL DEFAULT NOW(),-- Timestamp for when the order was created
    `billing_address` TEXT NOT NULL,              -- Billing address
    `payment_method` VARCHAR(50) NOT NULL,        -- Payment method
    `payment_status` VARCHAR(50) NOT NULL,        -- Payment status
    `order_date` TIMESTAMP NOT NULL DEFAULT NOW(),-- Date when the order was placed
    `shipment_date` TIMESTAMP,                    -- Date when the order was shipped
    `delivery_date` TIMESTAMP,                    -- Date when the order was delivered
    `tracking_number` INT,                        -- Tracking number for the shipment
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) -- Foreign key constraint
);