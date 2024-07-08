import { useQueryClient } from "@tanstack/react-query";
import { Button } from "antd";
import { Dispatch, SetStateAction, useEffect, useMemo } from "react";
import { MoveDirectionNames } from "../../../models/moveDirection.model";
import { BoardDimensions, Tiles } from "../../../models/tiles.model";
import { isSolved } from "../../../utils/board.util";
import { Board } from "../../Board";
import { puzzleKeys } from "../puzzle.keys";

type HandleMove = (moveDirection: MoveDirectionNames, tiles: Tiles) => void;
export interface UnsolvedPuzzleProps {
  tiles: Tiles;
  setIsAutoSolving: Dispatch<SetStateAction<boolean>>;
  handleMove: HandleMove;
  boardDimensions: BoardDimensions;
}

export const UnsolvedPuzzle = ({
  handleMove,
  setIsAutoSolving,
  tiles,
  boardDimensions,
}: UnsolvedPuzzleProps) => {
  const queryClient = useQueryClient();

  useControls(tiles, handleMove);

  const handleClick = () => {
    setIsAutoSolving(true);
  };

  const handleGenerateNewPuzzle = () => {
    queryClient.refetchQueries({
      exact: true,
      queryKey: puzzleKeys.generate(boardDimensions),
    });
  };

  const isPuzzleSolved = useMemo(() => {
    return isSolved(tiles);
  }, [tiles]);

  return (
    <>
      <Board tiles={tiles} />
      {isPuzzleSolved && <h3>הפאזל נפתר בהצלחה</h3>}
      <Button.Group size="large">
        <Button disabled={isPuzzleSolved} onClick={handleClick}>
          הצגת שלבי פתרון
        </Button>
        <Button onClick={handleGenerateNewPuzzle}>יצירת לוח חדש</Button>
      </Button.Group>
    </>
  );
};

const useControls = (tiles: Tiles, handleMove: HandleMove) => {
  useEffect(() => {
    const keyListener = (event: KeyboardEvent) => {
      switch (event.code) {
        case "Down":
        case "KeyS":
        case "ArrowDown":
          handleMove("DOWN", tiles);
          break;
        case "Up":
        case "KeyW":
        case "ArrowUp":
          handleMove("UP", tiles);
          break;
        case "Left":
        case "KeyA":
        case "ArrowLeft":
          handleMove("LEFT", tiles);
          break;
        case "Right":
        case "KeyD":
        case "ArrowRight":
          handleMove("RIGHT", tiles);
          break;
        default:
          break;
      }
    };

    document.addEventListener("keydown", keyListener);

    return () => {
      document.removeEventListener("keydown", keyListener);
    };
  }, [tiles, handleMove]);
};
