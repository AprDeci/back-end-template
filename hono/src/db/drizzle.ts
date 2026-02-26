import { drizzle } from "drizzle-orm/mysql2";
import { env } from "../config/env.js";

// 拼接url
const databaseUrl = `mysql://${env.DB_USER}:${env.DB_PASSWORD}@${env.DB_HOST}:${env.DB_PORT}/${env.DB_NAME}`;

const dbClient = drizzle({ connection: { uri: databaseUrl } });

export default dbClient;
