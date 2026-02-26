import z from "zod";

export const ResponseCode = {
  SUCCESS: 0,
  PARAM_ERROR: 10001,
  NOT_FOUND: 10002,
  DB_ERROR: 10003,
  UNAUTHORIZED: 10004,
  SERVER_ERROR: 10005
} as const;

export type AppResponse<T> = {
  code: number;
  message: string;
  data: T;
};

export const AppResponseSchema = z.object({
  code: z.number(),
  message: z.string(),
  data: z.any()
});

export function success<T>(data: T, message = "ok"): AppResponse<T> {
  return {
    code: ResponseCode.SUCCESS,
    message,
    data
  };
}

export function fail(code: number, message: string): AppResponse<null> {
  return {
    code,
    message,
    data: null
  };
}
