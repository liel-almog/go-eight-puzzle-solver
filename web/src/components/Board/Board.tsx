import { Tiles } from "../../models/tiles.model";
import classes from "./board.module.scss";
import { Cell } from "./Cell";
import { motion } from "framer-motion";

export interface BoardProps {
  tiles: Tiles;
}

export const Board = ({ tiles }: BoardProps) => {
  const grid = tiles.map((tileRow, rowIndex) => {
    const tile = tileRow.map((value, columnIndex) => (
      <Cell columnIndex={columnIndex} rowIndex={rowIndex} key={value} value={value} />
    ));

    return (
      <article className={classes.row} key={tileRow.toString()}>
        {tile}
      </article>
    );
  });

  return (
    <motion.section layout className={classes.board}>
      {grid}
    </motion.section>
  );
};
