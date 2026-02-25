import { HTTPException } from "hono/http-exception";
import { AppError } from "../core/error.js";
import { logger } from "../core/logger.js";
import { ResponseCode } from "../core/response.js";

export function toErrorResponse(error: Error): {
  status: number;
  code: number;
  message: string;
} {
  if (error instanceof AppError) {
    return {
      status: error.status,
      code: error.code,
      message: error.message
    };
  }

  if (error instanceof HTTPException) {
    return {
      status: error.status,
      code: error.status === 401 ? ResponseCode.UNAUTHORIZED : ResponseCode.SERVER_ERROR,
      message: error.message || "http error"
    };
  }

  return {
    status: 500,
    code: ResponseCode.SERVER_ERROR,
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
