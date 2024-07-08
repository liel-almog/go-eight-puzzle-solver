import { apiInstance } from ".";
import { BoardDimensions, Tiles, solutionStepsSchema, tilesSchema } from "../models/tiles.model";

const PREFIX = "puzzle" as const;

export class PuzzleService {
  static async generatePuzzle(generateBoardDimensions: BoardDimensions) {
    try {
      const response = await apiInstance.post(`/${PREFIX}/generate`, generateBoardDimensions);

      return tilesSchema.parse(response.data);
    } catch (error) {
      throw new Error("לא ניתן לייצר פאזל");
    }
  }

  static async bfsSolve(tiles: Tiles, signal: AbortSignal) {
    try {
      const response = await apiInstance.post(
        `/${PREFIX}/bfs/solve`,
        { tiles },
        {
          signal,
        }
      );

      return solutionStepsSchema.parse(response.data);
    } catch (error) {
      throw new Error("לא ניתן לפתור את הפאזל");
    }
  }
}
