CREATE TABLE `user` (
	`id` int AUTO_INCREMENT NOT NULL,
	`nickname` varchar(255),
	`age` int,
	`email` varchar(255),
	`username` varchar(255) NOT NULL,
	`password` varchar(255) NOT NULL,
	CONSTRAINT `user_id` PRIMARY KEY(`id`),
	CONSTRAINT `user_email_unique` UNIQUE(`email`),
	CONSTRAINT `user_username_unique` UNIQUE(`username`)
);
