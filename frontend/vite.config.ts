import {defineConfig} from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [react()],
    resolve: {
        alias: {
            '@': path.resolve(__dirname, './src'),
        },
    },
    build: {
        sourcemap: false,
        minify: 'esbuild',
        cssMinify: 'esbuild',
        // gzip size for every asset is slow on low-memory servers with many static files
        reportCompressedSize: false,
    },
    server: {
        proxy: {
            '/api': {
                target: 'https://176.123.167.9:443/api',
                //target: 'https://osudashboard.ru/api',
                //target: 'http://localhost:8080/api',
                secure: false,
                changeOrigin: true,
                rewrite: (path) => path.replace(/^\/api/, ''),
            },
        },
    },
})
