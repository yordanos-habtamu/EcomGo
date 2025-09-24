CREATE TABLE IF NOT EXISTS products (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY, -- Auto-incrementing ID
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT NOT NULL,
    `price` DECIMAL(10, 2) NOT NULL,
    `stock` INT NOT NULL,
    `category` VARCHAR(255) NOT NULL,
    `img_url` TEXT,
    `created_at` TIMESTAMP NOT NULL DEFAULT NOW(),
    `updated_at` TIMESTAMP NOT NULL DEFAULT NOW(),
    `is_active` BOOLEAN NOT NULL DEFAULT TRUE
);