import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";
import svgrPlugin from "vite-plugin-svgr";
import viteTsconfigPaths from "vite-tsconfig-paths";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), viteTsconfigPaths(), svgrPlugin()],
  server: {
    host: true,
    port: 3000,
    hmr: {
      host: 'localhost',
      protocol: 'wss',
      path: '/vite-development-wss'
    }
  },
});
