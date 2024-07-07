import react from "@vitejs/plugin-react-swc";
import { setDefaultResultOrder } from "dns";
import * as path from "path";
import { defineConfig } from "vite";
import Inspect from "vite-plugin-inspect";
import tsonfigpathes from "vite-tsconfig-paths";

setDefaultResultOrder("verbatim");
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), Inspect(), tsonfigpathes()],
  css: {
    modules: {
      generateScopedName: (name, filePath) => {
        const matches = path.basename(filePath).match(/^([a-z-]+)(.module)?.s?css/);
        if (!matches) {
          throw new Error("Could not match filename");
        }

        const baseFilename = matches[1];
        return `${baseFilename}-${name}`;
      },
      localsConvention: "camelCaseOnly",
    },
  },
  server: {
    port: 3000,
    open: false,
    strictPort: true,
    host: true,
  },
});
