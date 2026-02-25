import dotenv from "dotenv";
import { z } from "zod";

dotenv.config({ quiet: true });

const envSchema = z.object({
  NODE_ENV: z
    .enum(["development", "test", "production"])
    .default("development"),
  PORT: z.coerce.number().int().positive().default(8080),
  LOG_LEVEL: z
    .enum(["fatal", "error", "warn", "info", "debug", "trace", "silent"])
    .default("info"),
  DB_HOST: z.string().default("127.0.0.1"),
  DB_PORT: z.coerce.number().int().positive().default(5432),
  DB_USER: z.string().default("mysql"),
  DB_PASSWORD: z.string().default("password"),
  DB_NAME: z.string().default("hono_template"),
  JWT_SECRET: z.string().default("secret"),
  JWT_EXPIRED: z.coerce.number().int().positive().default(86400)
});

export const env = envSchema.parse(process.env);
