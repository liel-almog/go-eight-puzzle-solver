import { PropsWithChildren } from "react";
import { TanStackQueryClientProvider } from "./TanStackQueryProvider";
import { AntdConfigProvider } from "./AntdConfig.Provider";
import { AntdAppProvider } from "./AntdApp.Provider";

export const GlobalProvider = ({ children }: PropsWithChildren) => {
  return (
    <TanStackQueryClientProvider>
      <AntdConfigProvider>
        <AntdAppProvider>{children}</AntdAppProvider>
      </AntdConfigProvider>
    </TanStackQueryClientProvider>
  );
};
