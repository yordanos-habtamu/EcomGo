CREATE TABLE IF NOT EXISTS users (
    `id` SERIAL PRIMARY KEY,            -- Auto-incrementing ID
    `first_name` VARCHAR(255) NOT NULL, -- First name of the user
    `last_name` VARCHAR(255) NOT NULL,  -- Last name of the user
    `sex` ENUM('male', 'female') NOT NULL,         -- Gender/Sex
    `email` VARCHAR(255) UNIQUE NOT NULL, -- Email (must be unique)
    `DoB` DATE NOT NULL,                -- Date of Birth
    `password` VARCHAR(255) NOT NULL,   -- Hashed password
    `created_at` TIMESTAMP NOT NULL DEFAULT NOW(), -- Record creation time
    `role` ENUM('admin', 'customer', 'seller') NOT NULL DEFAULT 'customer'         -- User role (e.g., admin, user)
);

