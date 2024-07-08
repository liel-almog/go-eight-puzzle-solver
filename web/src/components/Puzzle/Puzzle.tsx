import { InputNumber, InputNumberProps, Space } from "antd";
import { Dispatch, SetStateAction, useState } from "react";
import { MoveDirectionNames } from "../../models/moveDirection.model";
import { BoardDimensions, Tiles } from "../../models/tiles.model";
import { PuzzleSolver } from "./PuzzleSolver";
import { UnsolvedPuzzle } from "./UnsolvedPuzzle";
import classes from "./puzzle.module.scss";
import { usePuzzle } from "./usePuzzle";

export interface PuzzleProps {}

export interface SolveButtonProps {
  tiles: Tiles;
  setSolving: Dispatch<SetStateAction<boolean>>;
}

export const Puzzle = () => {
  const [isAutoSolving, setIsAutoSolving] = useState(false);
  const {
    reducer: { dispatch, tiles },
    boardDimensions: { setBoardDimensions, boardDimensions },
    query: { isSuccess },
  } = usePuzzle();

  const handleMove = (moveDirection: MoveDirectionNames, tiles: Tiles) => {
    dispatch({
      type: "MOVE",
      payload: {
        moveDirection,
        tiles,
      },
    });
  };

  if (isSuccess) {
    return (
      <main className={classes.board}>
        <BoardDimensionsInput
          boardDimensions={boardDimensions}
          setBoardDimensions={setBoardDimensions}
          isAutoSolving={isAutoSolving}
        />
        {isAutoSolving ? (
          <PuzzleSolver tiles={tiles} setIsAutoSolving={setIsAutoSolving} />
        ) : (
          <UnsolvedPuzzle
            boardDimensions={boardDimensions}
            handleMove={handleMove}
            setIsAutoSolving={setIsAutoSolving}
            tiles={tiles}
          />
        )}
      </main>
    );
  }

  return <h1>אנחנו מכינים את הפאזל</h1>;
};

export interface BoardDimensionsInputProps {
  boardDimensions: BoardDimensions;
  isAutoSolving: boolean;
  setBoardDimensions: Dispatch<SetStateAction<BoardDimensions>>;
}

export const BoardDimensionsInput = ({
  boardDimensions: { columnCount, rowCount },
  setBoardDimensions,
  isAutoSolving: solving,
}: BoardDimensionsInputProps) => {
  const handleColumnCountChange: InputNumberProps["onChange"] = (columnCount) => {
    if (typeof columnCount === "number") {
      setBoardDimensions({
        columnCount,
        rowCount,
      });
    }
  };

  const handleRowCountChange: InputNumberProps["onChange"] = (rowCount) => {
    if (typeof rowCount === "number") {
      setBoardDimensions({
        columnCount,
        rowCount,
      });
    }
  };

  return (
    <Space.Compact size="large">
      <InputNumber
        value={columnCount}
        min={2}
        max={8}
        disabled={solving}
        onChange={handleColumnCountChange}
      />
      <InputNumber
        value={rowCount}
        min={2}
        max={8}
        disabled={solving}
        onChange={handleRowCountChange}
      />
    </Space.Compact>
  );
};
