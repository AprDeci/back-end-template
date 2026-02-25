import { decode, sign, verify } from "hono/jwt";

export const GenerateToken = (payload: any) => {
  return sign(payload, "secret");
};
