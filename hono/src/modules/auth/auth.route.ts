import { Hono } from "hono";
import { AppResponseSchema, success } from "../../core/response.js";
import type { AppEnv } from "../../types/env.js";
import { loginSchema, registerSchema } from "./auth.schema.js";
import * as authService from "./auth.service.js";
import { zv } from "../../utils/validator.js";
import { createRouteDescriber } from "../../utils/openapi.js";

export const authRoutes = new Hono<AppEnv>();

const describeAuthRoute = createRouteDescriber(["auth"], AppResponseSchema);

authRoutes.post(
  "/register",
  describeAuthRoute("Register"),
  zv("json", registerSchema),
  async (c) => {
    const input = await c.req.valid("json");
    const result = await authService.register(input);
    return c.json(success(result, "register success"));
  }
);

authRoutes.post(
  "/login",
  describeAuthRoute("Login"),
  zv("json", loginSchema),
  async (c) => {
    const input = await c.req.valid("json");
    const result = await authService.login(input);
    return c.json(success(result, "login success"));
  }
);

authRoutes.post("/logout", describeAuthRoute("Logout"), async (c) => {
  const result = await authService.logout();
  return c.json(success(result, "logout success"));
});
