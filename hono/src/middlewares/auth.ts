import type { Context, Next } from "hono";
import type { AppEnv } from "../types/env.js";
import { AppError } from "../core/error.js";
import { ResponseCode } from "../core/response.js";
import { parseTokenWithClaims } from "../utils/jwt.js";

const AUTH_WHITELIST = new Set<string>([
  "/",
  "/health",
  "/api/auth/login",
  "/api/auth/register",
  "/docs"
]);

function isWhitelistPath(path: string): boolean {
  if (AUTH_WHITELIST.has(path)) {
    return true;
  }

  return path.startsWith("/docs/");
}

function extractBearerToken(authHeader: string | undefined): string {
  const raw = authHeader?.trim() ?? "";

  if (!raw) {
    return "";
  }

  if (raw.toLowerCase().startsWith("bearer ")) {
    return raw.slice(7).trim();
  }

  return raw;
}

export async function auth(c: Context<AppEnv>, next: Next): Promise<void> {
  if (isWhitelistPath(c.req.path)) {
    await next();
    return;
  }

  const token = extractBearerToken(c.req.header("Authorization"));

  if (!token) {
    throw new AppError({
      status: 401,
      code: ResponseCode.UNAUTHORIZED,
      message: "unauthorized"
    });
  }

  try {
    const claims = await parseTokenWithClaims(token);
    c.set("userId", claims.user_id);
    c.set("username", claims.username);
    c.set("role", claims.role);
  } catch {
    throw new AppError({
      status: 401,
      code: ResponseCode.UNAUTHORIZED,
      message: "unauthorized"
    });
  }

  await next();
}
