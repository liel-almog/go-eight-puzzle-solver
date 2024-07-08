import { UseQueryOptions, useQuery } from "@tanstack/react-query";
import { puzzleKeys } from "./puzzle.keys";
import { Algorithms, PuzzleService } from "../../services/puzzle.service";
import { BoardDimensions, Tiles } from "../../models/tiles.model";

export const usePuzzleGenerator = (boardDimensions: BoardDimensions) => {
  return useQuery({
    queryKey: puzzleKeys.generate(boardDimensions),
    queryFn: () => PuzzleService.generatePuzzle(boardDimensions),
    staleTime: Infinity,
    throwOnError: true,
  });
};

export const usePuzzleSolver = (tiles: Tiles, algorithm: Algorithms, options?: Omit<UseQueryOptions<Tiles[]>, 'queryKey'>) => {
  return useQuery<Tiles[]>({
    queryKey: puzzleKeys.solve(tiles, algorithm),
    queryFn: ({ signal }) => PuzzleService.solve(tiles, algorithm, signal),
    staleTime: Infinity,
    throwOnError: true,
    ...options
  });
};
