import { z } from "zod";
import { positiveInt } from "./zod.util";

export const positionSchema = z.object({
  row: positiveInt,
  column: positiveInt,
});

export type Position = z.infer<typeof positionSchema>;
