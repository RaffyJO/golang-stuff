ALTER TABLE `users`
CHANGE COLUMN name username VARCHAR(255) NOT NULL;

ALTER TABLE `users`
ADD CONSTRAINT `users_username_unique` UNIQUE (`username`);