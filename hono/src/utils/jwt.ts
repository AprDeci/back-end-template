import { sign, verify } from "hono/jwt";
import { env } from "../config/env.js";

export type UserClaims = {
  user_id: number;
  username: string;
  role: string;
  exp: number;
};

export async function generateTokenWithUserInfo(input: {
  userId: number;
  username: string;
  role: string;
}): Promise<string> {
  const now = Math.floor(Date.now() / 1000);

  return sign(
    {
      user_id: input.userId,
      username: input.username,
      role: input.role,
      exp: now + env.JWT_EXPIRED
    },
    env.JWT_SECRET,
    "HS256"
  );
}

export async function parseTokenWithClaims(token: string): Promise<UserClaims> {
  const payload = await verify(token, env.JWT_SECRET, "HS256");

  if (
    typeof payload.user_id !== "number" ||
    typeof payload.username !== "string" ||
    typeof payload.role !== "string" ||
    typeof payload.exp !== "number"
  ) {
    throw new Error("invalid token payload");
  }

  return {
    user_id: payload.user_id,
    username: payload.username,
    role: payload.role,
    exp: payload.exp
  };
}
