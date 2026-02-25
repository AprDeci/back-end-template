export class AppError extends Error {
  status: number;
  code: number;

  constructor(options: { message: string; code: number; status?: number }) {
    super(options.message);
    this.name = "AppError";
    this.code = options.code;
    this.status = options.status ?? 500;
  }
}
