import type { Context, Next } from "hono";
import { logger } from "../core/logger.js";
import type { AppEnv } from "../types/env.js";

export async function requestLogger(c: Context<AppEnv>, next: Next): Promise<void> {
  const start = Date.now();
  await next();

  logger.info(
    {
      method: c.req.method,
      path: c.req.path,
      status: c.res.status,
      durationMs: Date.now() - start,
      requestId: c.get("requestId")
    },
    "request completed"
  );
}
