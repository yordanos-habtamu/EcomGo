CREATE TABLE IF NOT EXISTS permissions (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,          -- Unique ID for each permission
  `name` VARCHAR(255) NOT NULL,                       -- Name of the permission (e.g., 'view_dashboard', 'edit_profile')
  `description` TEXT,                                 -- Description of what the permission allows
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp when permission is created
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Timestamp for last update
  PRIMARY KEY(`id`)                                   -- Primary key for the table
);