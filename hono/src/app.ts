import { Hono } from "hono";
import { env } from "./config/env.js";
import type { AppEnv } from "./types/env.js";
import { reportError, toErrorResponse } from "./middlewares/errorHandler.js";
import { requestId } from "./middlewares/requestId.js";
import { requestLogger } from "./middlewares/requestLogger.js";

export function createApp(): Hono<AppEnv> {
  const app = new Hono<AppEnv>();

  app.use("*", requestId);
  app.use("*", requestLogger);

  app.get("/", (c) => {
    return c.json({
      message: "Hono template is running",
      env: env
    });
  });

  app.get("/health", (c) => {
    return c.json({
      status: "ok"
    });
  });

  app.onError((error, c) => {
    reportError(error, c.get("requestId"));
    const { code, message } = toErrorResponse(error);

    c.status(code as 400 | 401 | 403 | 404 | 500);
    return c.json({
      code,
      message
    });
  });

  app.notFound((c) => {
    return c.json(
      {
        code: 404,
        message: "route not found"
      },
      404
    );
  });

  return app;
}
