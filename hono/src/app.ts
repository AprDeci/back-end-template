import { Hono } from "hono";

import { env } from "./config/env.js";
import { fail, ResponseCode, success } from "./core/response.js";
import { authRoutes } from "./modules/auth/auth.route.js";
import type { AppEnv } from "./types/env.js";
import { auth } from "./middlewares/auth.js";
import { reportError, toErrorResponse } from "./middlewares/errorHandler.js";
import { requestId } from "./middlewares/requestId.js";
import { requestLogger } from "./middlewares/requestLogger.js";
import { openAPIRouteHandler } from "hono-openapi";
import { Scalar } from "@scalar/hono-api-reference";

export function createApp(): Hono<AppEnv> {
  const app = new Hono<AppEnv>();

  app.use("*", requestId);
  app.use("*", requestLogger);
  app.use("*", auth);

  app.get("/", (c) => {
    return c.json(
      success(
        {
          service: "hono-template",
          env: env.NODE_ENV
        },
        "service is running"
      )
    );
  });

  app.get(
    "/openapi",
    openAPIRouteHandler(app, {
      documentation: {
        info: {
          title: "Hono API",
          version: "1.0.0",
          description: "Greeting API"
        },
        servers: [
          { url: `http://localhost:${env.PORT}`, description: "Local Server" }
        ]
      }
    })
  );

  app.get("/scalar", Scalar({ url: "/openapi" }));

  app.route("/api/auth", authRoutes);

  app.get("/api/private/ping", (c) => {
    return c.json(
      success(
        {
          userId: c.get("userId"),
          username: c.get("username"),
          role: c.get("role")
        },
        "private ping success"
      )
    );
  });

  app.onError((error, c) => {
    reportError(error, c.get("requestId"));
    const { status, code, message } = toErrorResponse(error);

    c.status(status as 400 | 401 | 403 | 404 | 500);
    return c.json(fail(code, message));
  });

  app.notFound((c) => {
    c.status(404);
    return c.json(fail(ResponseCode.NOT_FOUND, "route not found"));
  });

  return app;
}
