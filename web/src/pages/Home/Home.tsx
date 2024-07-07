import { useNavigate } from "react-router-dom";
import classes from "./home.module.scss";
import { Controls } from "../../components/Controls";

export interface HomeProps {}

export const Home = (props: HomeProps) => {
  const navigate = useNavigate();
  return (
    <main className={classes.container}>
      <h1>ברוכים הבאים</h1>
      <section className={classes.getStarted}>
        <section className={classes.instructions}>
          <p>
            זהו משחק פאזל ההזנה המטרה שלכם היא להזיז את הפאזל כך שהמספרים יהיו מסודרים בסדר עולה
            משמאל לימין ומלמעלה למטה.
          </p>
          <p>אם אתם מסתבכים תמיד יש את האופציה לבקש מהאלגוריתם לנסות לפתור את הפאזל בשבילכם.</p>
        </section>
        <section className={classes.instructions}>
          <h3>אמצעי הזזה</h3>
          <p>ניתן להזיז את הפאזל בעזרת המקשים הללו</p>
          <div className={classes.controls}>
            <Controls up="W" left="A" right="D" down="S" />
            <Controls up="▲" left="◄" right="►" down="▼" />
          </div>
        </section>
        <button onClick={() => navigate("puzzle")} className={classes.btn}>
          בואו נתחיל
        </button>
      </section>
    </main>
  );
};
