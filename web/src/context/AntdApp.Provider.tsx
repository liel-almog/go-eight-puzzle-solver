import { App } from "antd";
import { PropsWithChildren } from "react";

export const AntdAppProvider = ({ children }: PropsWithChildren) => (
  <App
    style={{ width: "100%", height: "100%", fontFamily: "Noto Sans Hebrew" }}
    message={{ duration: 2, rtl: true }}
  >
    {children}
  </App>
);
