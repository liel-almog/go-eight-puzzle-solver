import { useQuery } from "@tanstack/react-query";
import { puzzleKeys } from "./puzzle.keys";
import { PuzzleService } from "../../services/puzzle.service";
import { BoardDimensions, Tiles } from "../../models/tiles.model";

export const usePuzzleGenerator = (boardDimensions: BoardDimensions) => {
  return useQuery({
    queryKey: puzzleKeys.generate(boardDimensions),
    queryFn: () => PuzzleService.generatePuzzle(boardDimensions),
    staleTime: Infinity,
    throwOnError: true,
  });
};

export const usePuzzleSolver = (tiles: Tiles) => {
  return useQuery({
    queryKey: puzzleKeys.solve(tiles),
    queryFn: ({ signal }) => PuzzleService.bfsSolve(tiles, signal),
    staleTime: Infinity,
    throwOnError: true,
  });
};
