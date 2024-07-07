import { BLANK_TILE, CellValue } from "../../../models/cell.model";
import classes from "./cell.module.scss";
import { motion } from "framer-motion";

export interface CellProps {
  value: CellValue;
  rowIndex: number;
  columnIndex: number;
}

export const Cell = ({ value, columnIndex, rowIndex }: CellProps) => {
  return <motion.div className={classes.cell}>{value === BLANK_TILE ? "" : value}</motion.div>;
};
