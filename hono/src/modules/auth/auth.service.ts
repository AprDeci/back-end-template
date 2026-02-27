import { AppError } from "../../core/error.js";
import { ResponseCode } from "../../core/response.js";
import { hashPassword, verifyPassword } from "../../utils/password.js";
import { generateTokenWithUserInfo } from "../../utils/jwt.js";
import type { LoginInput, RegisterInput } from "./auth.schema.js";
import dbClient from "../../db/drizzle.js";
import { usersTable } from "../../db/schema.js";
import { eq } from "drizzle-orm";

type RegisterResult = {
  id: number;
  username: string;
};

type LoginResult = {
  id: number;
  username: string;
  token: string;
};

export async function register(input: RegisterInput): Promise<RegisterResult> {
  const key = input.username.trim();
  const result = await dbClient
    .select()
    .from(usersTable)
    .where(eq(usersTable.username, key));
  if (result.length > 0) {
    throw new AppError({
      status: 400,
      code: ResponseCode.PARAM_ERROR,
      message: "username already exists"
    });
  }

  const newUser = {
    username: key,
    password: hashPassword(input.password)
  };

  const inserted = await dbClient
    .insert(usersTable)
    .values(newUser)
    .$returningId();

  return {
    id: inserted[0].id,
    username: newUser.username
  };
}

export async function login(input: LoginInput): Promise<LoginResult> {
  const key = input.username.trim();
  const users = await dbClient
    .select()
    .from(usersTable)
    .where(eq(usersTable.username, key));

  if (users.length === 0) {
    throw new AppError({
      status: 400,
      code: ResponseCode.NOT_FOUND,
      message: "user not found"
    });
  }

  const user = users[0];

  if (!verifyPassword(input.password, user.password)) {
    throw new AppError({
      status: 401,
      code: ResponseCode.UNAUTHORIZED,
      message: "password incorrect"
    });
  }

  const token = await generateTokenWithUserInfo({
    userId: user.id,
    username: user.username,
    role: "user"
  });

  return {
    id: user.id,
    username: user.username,
    token
  };
}

export async function logout(): Promise<{ msg: string }> {
  return { msg: "logout success" };
}
