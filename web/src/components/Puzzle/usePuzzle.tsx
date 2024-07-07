import { App } from "antd";
import { useEffect, useReducer, useState } from "react";
import { MOVE_DIRECTION_ENUM, MoveDirectionNames } from "../../models/moveDirection.model";
import { BoardDimensions, Tiles } from "../../models/tiles.model";
import { move } from "../../utils/board.util";
import { usePuzzleGenerator } from "./queries";

type Move = {
  type: "MOVE";
  payload: { tiles: Tiles; moveDirection: MoveDirectionNames };
};
type Reset = { type: "RESET"; payload: Tiles };

type TilesAction = Move | Reset;

const reducer = (tiles: Tiles, action: TilesAction) => {
  switch (action.type) {
    case "RESET":
      return action.payload;

    case "MOVE":
      return move(tiles, MOVE_DIRECTION_ENUM[action.payload.moveDirection]);

    default:
      return tiles;
  }
};

export const usePuzzle = () => {
  const { message } = App.useApp();
  const [boardDimensions, setBoardDimensions] = useState<BoardDimensions>({
    columnCount: 3,
    rowCount: 3,
  });

  const [tiles, dispatch] = useReducer(reducer, []);

  const { data, isSuccess, isError, error, isLoading, isFetchedAfterMount } =
    usePuzzleGenerator(boardDimensions);

  useEffect(() => {
    if (isSuccess) {
      dispatch({ type: "RESET", payload: data });
    }
  }, [data, isSuccess]);

  useEffect(() => {
    if (isError && error) {
      message.error({
        content: error.message,
        key: "puzzle-generator",
      });
    }
  }, [error, isError, message]);

  return {
    reducer: {
      tiles,
      dispatch,
    },
    query: {
      isLoading,
      isSuccess,
      isFetchedAfterMount,
    },
    boardDimensions: {
      boardDimensions,
      setBoardDimensions,
    },
  };
};
