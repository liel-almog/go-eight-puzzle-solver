import { useNavigate } from "react-router-dom";
import classes from "./header.module.scss";

export interface HeaderProps {}

export const Header = () => {
  const navigate = useNavigate();
  return (
    <header className={classes.container}>
      <div onClick={() => navigate("/")} className={classes.home}>
        <h1>Eight Puzzle Solver</h1>
      </div>
    </header>
  );
};
