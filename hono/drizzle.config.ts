import { defineConfig } from "drizzle-kit";

export default defineConfig({
  dialect: "mysql", // 'mysql' | 'sqlite' | 'turso'
  schema: "./src/db/schema.ts",
  out: "./drizzle",
  dbCredentials: {
    host: process.env.DB_HOST!,
    port: process.env.DB_PORT ? Number(process.env.DB_PORT) : 3306,
    user: process.env.DB_USER!,
    password: process.env.DB_PASSWORD!,
    database: process.env.DB_NAME!
  }
});
