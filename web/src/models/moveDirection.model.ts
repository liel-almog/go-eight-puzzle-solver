import { z } from "zod";

export const MOVE_DIRECTION_NAMES_ENUM = {
  LEFT: "LEFT",
  RIGHT: "RIGHT",
  UP: "UP",
  DOWN: "DOWN",
} as const;

export const MOVE_DIRECTION_NAMES = [
  MOVE_DIRECTION_NAMES_ENUM.LEFT,
  MOVE_DIRECTION_NAMES_ENUM.RIGHT,
  MOVE_DIRECTION_NAMES_ENUM.UP,
  MOVE_DIRECTION_NAMES_ENUM.DOWN,
] as const;

export const moveDirectionNamesSchema = z.enum(MOVE_DIRECTION_NAMES);
export type MoveDirectionNames = z.infer<typeof moveDirectionNamesSchema>;

type MoveValue = -1 | 0 | 1;
export type MoveDirection = {
  row: MoveValue;
  column: MoveValue;
};

export const MOVE_DIRECTION_ENUM = {
  LEFT: {
    column: -1,
    row: 0,
  },
  RIGHT: {
    column: 1,
    row: 0,
  },
  UP: {
    column: 0,
    row: -1,
  },
  DOWN: {
    column: 0,
    row: 1,
  },
} satisfies Record<MoveDirectionNames, MoveDirection>;
