import { ReactNode } from "react";
import classes from "./controls.module.scss";

export interface ControlsProps {
  up: ReactNode;
  down: ReactNode;
  left: ReactNode;
  right: ReactNode;
}

export const Controls = ({ down, left, right, up }: ControlsProps) => {
  return (
    <article className={classes.container}>
      <span className={classes.key}>{up}</span>
      <footer className={classes.wrapper}>
        <span className={classes.key}>{right}</span>
        <span className={classes.key}>{down}</span>
        <span className={classes.key}>{left}</span>
      </footer>
    </article>
  );
};
