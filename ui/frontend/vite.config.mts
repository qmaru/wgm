import tailwindcss from "@tailwindcss/vite"
import react from "@vitejs/plugin-react"
import { defineConfig } from "vite"

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    sourcemap: false,
  },
  resolve: {
    tsconfigPaths: true,
  },
  plugins: [react(), tailwindcss()],
})
