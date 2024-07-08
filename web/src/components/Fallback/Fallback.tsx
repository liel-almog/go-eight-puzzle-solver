import { QueryStatus } from "@tanstack/react-query";
import { Spin } from "antd";
import clsx from "clsx";
import { PropsWithChildren, Suspense } from "react";
import classes from "./fallback.module.scss";

export interface StatusIndicatorProps {
  status: QueryStatus;
}

export const StatusIndicator = ({ children, status }: PropsWithChildren<StatusIndicatorProps>) => {
  if (status === "pending") {
    return (
      <div className={clsx(classes.spinner)}>
        <Spin spinning size="large" delay={300} />
      </div>
    );
  }

  if (!children) return <></>;

  if (status === "success") {
    return <>{children}</>;
  }

  return <></>;
};
export interface FallbackProps extends StatusIndicatorProps {}

export const Fallback = ({ children, status }: PropsWithChildren<FallbackProps>) => (
  <Suspense fallback={<StatusIndicator status={status} />}>{children}</Suspense>
);
