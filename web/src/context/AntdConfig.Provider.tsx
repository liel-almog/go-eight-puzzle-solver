import { ConfigProvider } from "antd";
import { PropsWithChildren } from "react";
import locale from "antd/locale/he_IL";

export const AntdConfigProvider = ({ children }: PropsWithChildren) => (
  <ConfigProvider
    theme={{
      token: {
        fontFamily: "Noto Sans Hebrew",
      },
    }}
    direction="rtl"
    locale={locale}
  >
    {children}
  </ConfigProvider>
);
