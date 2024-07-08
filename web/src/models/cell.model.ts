import { z } from "zod";

export const BLANK_TILE = -1 as const;

export const cellValueSchema = z.number().int().positive().or(z.literal(BLANK_TILE));
export type CellValue = z.infer<typeof cellValueSchema>;
