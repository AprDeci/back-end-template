import { randomBytes, scryptSync, timingSafeEqual } from "node:crypto";

const KEY_LENGTH = 64;

export function hashPassword(plainText: string): string {
  const salt = randomBytes(16).toString("hex");
  const hash = scryptSync(plainText, salt, KEY_LENGTH).toString("hex");
  return `${salt}:${hash}`;
}

export function verifyPassword(plainText: string, stored: string): boolean {
  const parts = stored.split(":");
  if (parts.length !== 2) {
    return false;
  }

  const [salt, hash] = parts;
  const derived = scryptSync(plainText, salt, KEY_LENGTH);
  const storedHash = Buffer.from(hash, "hex");

  if (derived.length !== storedHash.length) {
    return false;
  }

  return timingSafeEqual(derived, storedHash);
}
