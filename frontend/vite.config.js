import { defineConfig } from "vite";
import reactRefresh from "@vitejs/plugin-react-refresh";

export default defineConfig({
  plugins: [reactRefresh()],
  server: {
    proxy: {
      "/query": {
        target: "http://localhost:8081/",
      },
      "/oauth": {
        target: "http://localhost:8081/"
      }
    },
  },
});
