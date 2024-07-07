import { BoardDimensions, Tiles } from "../../models/tiles.model";

export const puzzleKeys = {
  generate(boardDimensions: BoardDimensions) {
    return ["puzzle", "generate", boardDimensions] as const;
  },
  solve: (tiles: Tiles) => {
    return ["puzzle", "bfs", "solve", tiles] as const;
  },
};
