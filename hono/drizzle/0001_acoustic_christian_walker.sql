ALTER TABLE `user` ADD `updated_at` timestamp;--> statement-breakpoint
ALTER TABLE `user` ADD `created_at` timestamp DEFAULT (now()) NOT NULL;--> statement-breakpoint
ALTER TABLE `user` ADD `deleted_at` timestamp;