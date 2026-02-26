import { Hono } from "hono";
import { ZodError } from "zod";
import { AppError } from "../../core/error.js";
import { ResponseCode, success } from "../../core/response.js";
import type { AppEnv } from "../../types/env.js";
import { loginSchema, registerSchema } from "./auth.schema.js";
import * as authService from "./auth.service.js";
import { validator as zValidator, resolver, describeRoute } from "hono-openapi";

export const authRoutes = new Hono<AppEnv>();

function parseBody<T>(
  payload: unknown,
  parser: { parse: (input: unknown) => T }
): T {
  try {
    return parser.parse(payload);
  } catch (error) {
    if (error instanceof ZodError) {
      throw new AppError({
        status: 400,
        code: ResponseCode.PARAM_ERROR,
        message: error.issues[0]?.message ?? "param error"
      });
    }

    throw error;
  }
}

authRoutes.post("/register", async (c) => {
  const payload = await c.req.json();
  const input = parseBody(payload, registerSchema);
  const result = await authService.register(input);

  return c.json(success(result, "register success"));
});

authRoutes.post("/regisdter", zValidator("json", registerSchema), async (c) => {
  const payload = await c.req.json();
  const input = parseBody(payload, registerSchema);
  const result = await authService.register(input);

  return c.json(success(result, "register success"));
});

authRoutes.post("/login", async (c) => {
  const payload = await c.req.json();
  const input = parseBody(payload, loginSchema);
  const result = await authService.login(input);

  return c.json(success(result, "login success"));
});

authRoutes.post("/logout", async (c) => {
  const result = await authService.logout();
  return c.json(success(result, "logout success"));
});
