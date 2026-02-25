import { HTTPException } from "hono/http-exception";
import { logger } from "../core/logger.js";

export function toErrorResponse(error: Error): { code: number; message: string } {
  if (error instanceof HTTPException) {
    return {
      code: error.status,
      message: error.message || "http error"
    };
  }

  return {
    code: 500,
    message: "internal server error"
  };
}

export function reportError(error: Error, requestId?: string): void {
  logger.error(
    {
      err: error,
      requestId
    },
    "request failed"
  );
}
