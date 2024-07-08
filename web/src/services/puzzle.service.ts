import { apiInstance } from ".";
import { BoardDimensions, Tiles, solutionStepsSchema, tilesSchema } from "../models/tiles.model";

const PREFIX = "puzzle" as const;
export const algorithms = ['BFS', 'DFS', 'ASTAR'] as const
export type Algorithms = typeof algorithms[number]


export class PuzzleService {
  static async generatePuzzle(generateBoardDimensions: BoardDimensions) {
    try {
      const response = await apiInstance.post(`/${PREFIX}/generate`, generateBoardDimensions);

      return tilesSchema.parse(response.data);
    } catch (error) {
      throw new Error("לא ניתן לייצר פאזל");
    }
  }

  static async solve(tiles: Tiles, algorithm: Algorithms, signal: AbortSignal) {
    try {
      const response = await apiInstance.post(
        `/${PREFIX}/${algorithm.toLocaleLowerCase()}/solve`,
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
