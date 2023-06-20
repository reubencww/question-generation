import { defineConfig } from "vite";
import { resolve } from "path";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    server: {
        open: false,
        origin: "http://localhost:5173",
    },
    resolve: {
        alias: {
            "@": resolve(__dirname, "./src"),
        },
    },
    build: {
        manifest: "manifest.json",
        rollupOptions: {
            input: {
                main: resolve("./src/main.ts"),
            },
        },
    },
});
