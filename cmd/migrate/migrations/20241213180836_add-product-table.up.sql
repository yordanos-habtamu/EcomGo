CREATE TABLE IF NOT EXISTS products (
    `id` SERIAL PRIMARY KEY,              -- Auto-incrementing ID
    `name` VARCHAR(255) NOT NULL,         -- Name of the product
    `description` TEXT NOT NULL,          -- Description of the product
    `price` DECIMAL(10, 2) NOT NULL,      -- Price with up to 2 decimal places
    `stock` INT NOT NULL,                 -- Number of items in stock
    `category` VARCHAR(255) NOT NULL,     -- Category of the product
    `img_url` TEXT,                       -- URL for the product image
    `created_at` TIMESTAMP NOT NULL DEFAULT NOW(), -- Record creation timestamp
    `updated_at` TIMESTAMP NOT NULL DEFAULT NOW(), -- Last update timestamp
    `is_active` BOOLEAN NOT NULL DEFAULT TRUE     -- Whether the product is active
);