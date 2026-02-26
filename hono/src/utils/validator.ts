import type { ValidationTargets } from "hono";
import type { z } from "zod";
import { validator as zValidator } from "hono-openapi";
import { AppError } from "../core/error.js";
import { ResponseCode } from "../core/response.js";

type ValidatorTarget = Extract<keyof ValidationTargets, "json" | "query" | "param">;

export const zv = <T extends z.ZodTypeAny>(
  target: ValidatorTarget,
  schema: T
) =>
  zValidator(target, schema, (result) => {
    if (!result.success) {
      const firstIssue = result.error[0];
      throw new AppError({
        status: 400,
        code: ResponseCode.PARAM_ERROR,
        message: firstIssue?.message ?? "param error"
      });
    }
  });
