import { BoardDimensions, Tiles } from "../../models/tiles.model";
import { Algorithms } from "../../services/puzzle.service";

export const puzzleKeys = {
  generate(boardDimensions: BoardDimensions) {
    return ["puzzle", "generate", boardDimensions] as const;
  },
  solve: (tiles: Tiles, algorithm: Algorithms) => {
    return ["puzzle", algorithm, "solve", tiles] as const;
  },
};
