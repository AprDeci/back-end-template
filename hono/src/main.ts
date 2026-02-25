import { serve } from "@hono/node-server";
import { env } from "./config/env.js";
import { logger } from "./core/logger.js";
import { createApp } from "./app.js";

const app = createApp();

let server: ReturnType<typeof serve>;

try {
  server = serve(
    {
      fetch: app.fetch,
      port: env.PORT
    },
    (info) => {
      logger.info({ port: info.port }, "server started");
    }
  );
} catch (error) {
  const err = error as NodeJS.ErrnoException;

  if (err.code === "EADDRINUSE") {
    logger.error({ port: env.PORT }, "port already in use, set another PORT in .env");
    process.exit(1);
  }

  logger.error({ err }, "server failed to start");
  process.exit(1);
}

const gracefulShutdown = (signal: string): void => {
  logger.info({ signal }, "shutdown signal received");
  server.close((error) => {
    if (error) {
      logger.error({ err: error }, "shutdown failed");
      process.exit(1);
    }

    logger.info("server stopped");
    process.exit(0);
  });
};

process.on("SIGINT", () => gracefulShutdown("SIGINT"));
process.on("SIGTERM", () => gracefulShutdown("SIGTERM"));
