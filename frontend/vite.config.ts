import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import tsconfigPaths from "vite-tsconfig-paths";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), tsconfigPaths()],
  server: {
    port: 3000,
    proxy: {
      "/query": {
        target: "http://localhost:8081/",
      },
      "/oauth": {
        target: "http://localhost:8081/",
      },
    },
  },
});
