import { Outlet } from "react-router-dom";
import classes from "./layout.module.scss";
import { Header } from "../../components/Header";
import { Footer } from "../../components/Footer";

export const Layout: React.FC = () => {
  return (
    <div className={classes.container}>
      <Header />
      <main>
        <Outlet />
      </main>
      <Footer />
    </div>
  );
};
