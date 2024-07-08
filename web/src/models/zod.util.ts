import { z } from "zod";

export const positiveInt = z.number().int().positive();
