import { drizzle } from "drizzle-orm/mysql2";
import type { Logger } from "drizzle-orm/logger";
import { env } from "../config/env.js";
import { logger } from "../core/logger.js";

class MyLogger implements Logger {
  logQuery(query: string, params: unknown[]): void {
    logger.debug({ query, params });
  }
}

// 拼接url
const databaseUrl = `mysql://${env.DB_USER}:${env.DB_PASSWORD}@${env.DB_HOST}:${env.DB_PORT}/${env.DB_NAME}`;

const dbClient = drizzle({
  connection: { uri: databaseUrl },
  logger: new MyLogger()
});

export default dbClient;
