import { useQueryClient } from "@tanstack/react-query";
import { Button, Select } from "antd";
import classes from './unsolved-puzzle.module.scss'
import { Dispatch, SetStateAction, useEffect, useMemo, useState } from "react";
import { MoveDirectionNames } from "../../../models/moveDirection.model";
import { BoardDimensions, Tiles } from "../../../models/tiles.model";
import { isSolved } from "../../../utils/board.util";
import { Board } from "../../Board";
import { puzzleKeys } from "../puzzle.keys";
import { usePuzzleSolver } from "../queries";
import { Algorithms, algorithms } from "../../../services/puzzle.service";

type HandleMove = (moveDirection: MoveDirectionNames, tiles: Tiles) => void;
export interface UnsolvedPuzzleProps {
  tiles: Tiles;
  setIsAutoSolving: Dispatch<SetStateAction<boolean>>;
  handleMove: HandleMove;
  boardDimensions: BoardDimensions;
  algorithm: Algorithms
  setAlgorithm: (algorithm: Algorithms) => void
}

export const UnsolvedPuzzle = ({
  handleMove,
  setIsAutoSolving,
  tiles,
  boardDimensions,
  algorithm,
  setAlgorithm
}: UnsolvedPuzzleProps) => {
  const queryClient = useQueryClient();

  useControls(tiles, handleMove);

  const handleClick = () => {
    setIsAutoSolving(true);
    queryClient.fetchQuery({
      queryKey: puzzleKeys.solve(tiles, algorithm),
    })
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

  const algorithmSelection = <Select className={classes.algoSelect} onSelect={(_v, opt) => {
    setAlgorithm(opt.value)
  }} defaultValue='BFS' value={algorithm} options={algorithms.map((a) => {
    return { value: a, label: a }
  })} />


  return (
    <>
      <Board tiles={tiles} />
      {isPuzzleSolved && <h3>הפאזל נפתר בהצלחה</h3>}
      <div>
        <p>{algorithmSelection} פתור באמצעות </p>
      </div>
      <Button.Group size="large">
        <Button disabled={isPuzzleSolved} onClick={handleClick}>
          פתור פאזל
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
