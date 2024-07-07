import { BLANK_TILE } from "../models/cell.model";
import { MoveDirection } from "../models/moveDirection.model";
import { Position } from "../models/position.model";
import { Tiles } from "../models/tiles.model";

const getBlankTilePosition = (tiles: Tiles) => {
  for (let i = 0; i < tiles.length; i++) {
    for (let j = 0; j < tiles[0].length; j++) {
      if (tiles[i][j] === BLANK_TILE) {
        return {
          row: i,
          column: j,
        } satisfies Position;
      }
    }
  }

  throw new Error("Could not find blank tile");
};

export const removeBlankTile = (tiles: Tiles) => {
  return tiles.flat().filter((tile) => tile !== BLANK_TILE);
};

export const moveBlankTile = (tiles: Tiles, moveDirection: MoveDirection) => {
  const blankTilePosition = getBlankTilePosition(tiles);
  const newBlankTileRow = blankTilePosition.row + moveDirection.row;
  const newBlankTileColumn = blankTilePosition.column + moveDirection.column;

  return {
    column: newBlankTileColumn,
    row: newBlankTileRow,
  } satisfies Position;
};

export const canMove = (tiles: Tiles, moveDirection: MoveDirection) => {
  try {
    const { column, row } = moveBlankTile(tiles, moveDirection);
    return column >= 0 && column <= tiles[0].length - 1 && row >= 0 && row <= tiles.length - 1;
  } catch (error) {
    return false;
  }
};

export const move = (tiles: Tiles, moveDirection: MoveDirection) => {
  if (!canMove(tiles, moveDirection)) {
    return structuredClone(tiles);
  }

  // After the canMove we can be sure that the blank tile is inside the tiles
  const { column: blankTileColumn, row: blankTileRow } = getBlankTilePosition(tiles);
  const { column: newBlankTileColumn, row: newBlankTileRow } = moveBlankTile(tiles, moveDirection);
  const newTiles = structuredClone(tiles);
  newTiles[newBlankTileRow][newBlankTileColumn] = BLANK_TILE;
  newTiles[blankTileRow][blankTileColumn] = tiles[newBlankTileRow][newBlankTileColumn];

  return newTiles;
};

export const isSolved = (tiles: Tiles) => {
  if (tiles.length === 0) {
    return false;
  }

  if (tiles[tiles.length - 1][tiles[0].length - 1] !== BLANK_TILE) {
    return false;
  }

  return removeBlankTile(tiles)
    .flat()
    .every((value, index, array) => index === 0 || value >= array[index - 1]);
};
