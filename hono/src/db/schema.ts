import { int, mysqlTable, timestamp, varchar } from "drizzle-orm/mysql-core";

const timestamps = {
  updated_at: timestamp().onUpdateNow(),
  created_at: timestamp().defaultNow().notNull(),
  deleted_at: timestamp()
};

export const usersTable = mysqlTable("user", {
  id: int().autoincrement().primaryKey(),
  nickname: varchar({ length: 255 }),
  age: int(),
  email: varchar({ length: 255 }).unique(),
  username: varchar({ length: 255 }).notNull().unique(),
  password: varchar({ length: 255 }).notNull(),
  ...timestamps
});
