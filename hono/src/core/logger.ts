import pino from "pino";
import pinoPretty from "pino-pretty";

import { env } from "../config/env.js";

export const logger = pino({
  level: env.LOG_LEVEL,
  base: undefined,
  timestamp: pino.stdTimeFunctions.isoTime,
  transport: {
    target: "pino-pretty"
  }
});
