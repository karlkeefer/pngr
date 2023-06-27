/// <reference types="vitest" />

import path from "path";

import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";
import svgrPlugin from "vite-plugin-svgr";
import viteTsconfigPaths from "vite-tsconfig-paths";

// https://vitejs.dev/config/
export default defineConfig({
  root: "./src",
  plugins: [react(), viteTsconfigPaths(), svgrPlugin()],
  test: {
    globals: true,
    environment: 'jsdom',
    coverage: {
      reporter: ['text', 'html'],
      exclude: [
        'node_modules/',
      ],
    },
    cache: false
  },
  build: {
    outDir: "../build",
  },
  server: {
    host: true,
    port: 3000,
    hmr: {
      host: 'localhost',
      protocol: 'wss',
      path: '/vite-development-wss'
    }
  },
  css: {
    preprocessorOptions: {
      less: {
        math: "always",
      }
    },
  },
  resolve: {
    // semantic-ui theming requires the import path of theme.config to be rewritten to our local theme.config file
    alias: {
      "../../theme.config": path.resolve(
        __dirname,
        "./src/semantic-ui/theme.config"
      ),
    },
  },
});
