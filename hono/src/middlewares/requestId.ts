import type { Context, Next } from "hono";
import { randomUUID } from "node:crypto";
import type { AppEnv } from "../types/env.js";

const REQUEST_ID_HEADER = "x-request-id";

export async function requestId(c: Context<AppEnv>, next: Next): Promise<void> {
  const incomingRequestId = c.req.header(REQUEST_ID_HEADER);
  const id = incomingRequestId?.trim() || randomUUID();

  c.set("requestId", id);
  c.header(REQUEST_ID_HEADER, id);

  await next();
}
