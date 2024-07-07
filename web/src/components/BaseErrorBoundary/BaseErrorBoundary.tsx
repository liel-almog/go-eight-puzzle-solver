import { faTriangleExclamation } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Button, Collapse } from "antd";
import { useNavigate, useRouteError } from "react-router-dom";
import classes from "./base-error-boundary.module.scss";
export const BaseErrorBoundary = () => {
  const error = useRouteError() as Error;
  const navigate = useNavigate();

  return (
    <div className={classes.container}>
      <FontAwesomeIcon icon={faTriangleExclamation} />
      <h1>אירעה שגיאה</h1>

      <Button className={classes.btn} onClick={() => navigate("/")}>
        למעבר לעמוד הראשי
      </Button>
      <Collapse
        style={{
          position: "fixed",
          bottom: "1rem",
        }}
        size="small"
        items={[
          {
            key: "1",
            label: "לצפייה בפרטים נוספים",
            children: <p>{error.message}</p>,
          },
        ]}
      />
    </div>
  );
};
