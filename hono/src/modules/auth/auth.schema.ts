import { z } from "zod";
export const registerSchema = z.object({
  username: z.string().min(1, "username is required"),
  password: z.string().min(1, "password is required")
});

export const loginSchema = z.object({
  username: z.string().min(1, "username is required"),
  password: z.string().min(1, "password is required")
});

export type RegisterInput = z.infer<typeof registerSchema>;
export type LoginInput = z.infer<typeof loginSchema>;
