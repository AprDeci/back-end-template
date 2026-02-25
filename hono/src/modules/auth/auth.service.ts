import { AppError } from "../../core/error.js";
import { ResponseCode } from "../../core/response.js";
import { hashPassword, verifyPassword } from "../../utils/password.js";
import { generateTokenWithUserInfo } from "../../utils/jwt.js";
import type { LoginInput, RegisterInput } from "./auth.schema.js";

type AuthUser = {
  id: number;
  username: string;
  passwordHash: string;
};

type RegisterResult = {
  id: number;
  username: string;
};

type LoginResult = {
  id: number;
  username: string;
  token: string;
};

const usersByName = new Map<string, AuthUser>();
let nextId = 1;

export async function register(input: RegisterInput): Promise<RegisterResult> {
  const key = input.username.trim();
  if (usersByName.has(key)) {
    throw new AppError({
      status: 400,
      code: ResponseCode.PARAM_ERROR,
      message: "username already exists"
    });
  }

  const newUser: AuthUser = {
    id: nextId,
    username: key,
    passwordHash: hashPassword(input.password)
  };

  nextId += 1;
  usersByName.set(key, newUser);

  return {
    id: newUser.id,
    username: newUser.username
  };
}

export async function login(input: LoginInput): Promise<LoginResult> {
  const key = input.username.trim();
  const user = usersByName.get(key);

  if (!user) {
    throw new AppError({
      status: 400,
      code: ResponseCode.NOT_FOUND,
      message: "user not found"
    });
  }

  if (!verifyPassword(input.password, user.passwordHash)) {
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
