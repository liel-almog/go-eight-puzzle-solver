import {
  faCancel,
  faPause,
  faPlay,
  faRepeat,
  faStepBackward,
  faStepForward,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { QueryStatus, useQueryClient } from "@tanstack/react-query";
import { Button, Spin } from "antd";
import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { Tiles } from "../../../models/tiles.model";
import { Board } from "../../Board";
import { puzzleKeys } from "../puzzle.keys";
import classes from "./puzzle-solver.module.scss";
import { Algorithms, algorithms } from "../../../services/puzzle.service";

export interface PuzzleSolverProps {
  tiles: Tiles;
  setIsAutoSolving: Dispatch<SetStateAction<boolean>>;
  algorithm: Algorithms
  query: {
    isSuccess: boolean,
    data: Tiles[] | undefined,
    status: QueryStatus
  }
}

export const PuzzleSolver = ({ tiles, setIsAutoSolving, algorithm, query: { data, isSuccess, status } }: PuzzleSolverProps) => {
  const [animationPlaying, setAnimationPlaying] = useState(false);
  const [currentStep, setCurrentStep] = useState(0);
  const queryClient = useQueryClient();


  // We are making a step every 1 second. We clear the interval after the animation is in the last step
  useEffect(() => {
    const interval = setInterval(() => {
      if (animationPlaying && isSuccess && data) {
        if (currentStep === data.length - 1) {
          clearInterval(interval);
          setAnimationPlaying(false);
        } else {
          setCurrentStep((prev) => prev + 1);
        }
      }
    }, 1000);

    return () => clearInterval(interval);
  }, [animationPlaying, currentStep, data?.length, isSuccess]);

  const handleCancelSolving = () => {
    setIsAutoSolving(false);
    queryClient.cancelQueries({
      exact: true,
      queryKey: puzzleKeys.solve(tiles, algorithm),
    });
  };

  if (status === "pending" || !data) {
    return (
      <>
        <Spin className={classes.spin} spinning size="large" delay={300}>
          <Board tiles={tiles} />
        </Spin>
        <Button
          onClick={handleCancelSolving}
          size="large"
          icon={<FontAwesomeIcon size="lg" icon={faCancel} />}
        />
      </>
    );
  }

  if (status === "error") {
    return (
      <>
        <Board tiles={tiles} />
        <p>לצערינו לא ניתן לפתור פאזל זה</p>
      </>
    );
  }

  const isFirstStep = currentStep === 0;
  const isLastStep = currentStep === data.length - 1;

  const handleAnimationClick = () => {
    setAnimationPlaying((prev) => !prev);
  };

  const handleStepBackwardClick = () => {
    if (!isFirstStep) {
      setCurrentStep((prev: number) => prev - 1);
    }
  };

  const handleStepForwardClick = () => {
    if (!isLastStep) {
      setCurrentStep((prev: number) => prev + 1);
    }
  };

  const handleRepeatClick = () => {
    setCurrentStep(0);
    setAnimationPlaying(true);
  };

  const animationIcon = animationPlaying && !isLastStep ? faPause : faPlay;
  const handleStopSolving = () => {
    setIsAutoSolving(false);
  };

  return (
    <>
      <Button onClick={handleStopSolving}>נסה לפתור לבד</Button>
      <Board tiles={data[currentStep]} />
      <Button.Group>
        <Button
          onClick={handleStepBackwardClick}
          size="large"
          disabled={isFirstStep}
          icon={<FontAwesomeIcon size="lg" icon={faStepBackward} />}
        />
        <Button
          onClick={handleAnimationClick}
          size="large"
          icon={<FontAwesomeIcon size="lg" icon={animationIcon} />}
        />
        <Button
          onClick={handleStepForwardClick}
          size="large"
          disabled={isLastStep}
          icon={<FontAwesomeIcon size="lg" icon={faStepForward} />}
        />
        <Button
          size="large"
          onClick={handleRepeatClick}
          icon={<FontAwesomeIcon size="lg" icon={faRepeat} />}
        />
      </Button.Group>
    </>
  );
};
