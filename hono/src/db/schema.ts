import { int, mysqlTable, serial, varchar } from "drizzle-orm/mysql-core";

export const usersTable = mysqlTable("user", {
  id: int().autoincrement().primaryKey(),
  nickname: varchar({ length: 255 }).notNull(),
  age: int().notNull(),
  email: varchar({ length: 255 }).notNull().unique(),
  username: varchar({ length: 255 }).notNull().unique(),
  password: varchar({ length: 255 }).notNull()
});
