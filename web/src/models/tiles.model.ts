import { z } from "zod";
import { cellValueSchema } from "./cell.model";
import { positiveInt } from "./zod.util";

export const tilesSchema = z.array(z.array(cellValueSchema).min(2).max(8)).min(2).max(8);
export type Tiles = z.infer<typeof tilesSchema>;

export const solutionStepsSchema = z.array(tilesSchema);
export type SolutionSteps = z.infer<typeof solutionStepsSchema>;

export const boardDimensionsSchema = z.object({
  rowCount: positiveInt,
  columnCount: positiveInt,
});
export type BoardDimensions = z.infer<typeof boardDimensionsSchema>;
