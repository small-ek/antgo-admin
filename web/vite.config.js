import {defineConfig} from 'vite'
import {resolve} from 'path'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite';
import viteCompression from 'vite-plugin-compression'
import {ArcoResolver} from 'unplugin-vue-components/resolvers';
import {vitePluginForArco} from '@arco-plugins/vite-vue'
import seoPrerender from 'vite-plugin-seo-prerender'

// https://vitejs.dev/config/
export default defineConfig({
    resolve: {
        alias: {
            '@': resolve(__dirname, 'src'),
        }
    },
    plugins: [
        vue(),
        //预加载
        seoPrerender({
            hashHistory: false,　// 使用hash路由，需设置为true
            routes: ['/test', "/test2"], // 需要生成的路由,无法使用二级目录
            callback: (html, route) => { //渲染和处理public下.html文件处理回调事件

            },
            removeStyle: false,// 移除预览服务生成多余样式，默认true。如样式丢失，可设置为false
            network: true // 构建时是否获取异步数据并注入到预渲染的 HTML 文件中
        }),
        AutoImport({
            resolvers: [ArcoResolver()],

        }),
        Components({
            resolvers: [
                ArcoResolver({
                    sideEffect: true
                })
            ]
        }),
        vitePluginForArco({
            style: 'css'
        }),
        // 压缩插件
        viteCompression({
            filter: /\.(js|css|json|txt|ico|svg)(\?.*)?$/i, // 需要压缩的文件
            verbose: true, // 默认即可
            disable: false, // 开启压缩(不禁用)，默认即可
            deleteOriginFile: false, // 删除源文件
            threshold: 10240, // 设置只有大于 10kb 的文件才会被压缩, 单位b
            algorithm: 'gzip', // 压缩算法
            ext: '.gz' // 文件类型
        }),

    ],
    //base: "web",
    server: {
        host: "0.0.0.0",
        open: true
    },
    build: {
        target: 'es2015',
        cssTarget: 'chrome61', // 此选项允许用户为 CSS 的压缩设置一个不同的浏览器 target，此处的 target 并非是用于 JavaScript 转写目标。
        cssCodeSplit: true, // 启用/禁用 CSS 代码拆分
        rollupOptions: {
            output: {// 分包
                chunkFileNames: 'static/js/[name]-[hash].js',
                entryFileNames: 'static/js/[name]-[hash].js',
                assetFileNames: 'static/[ext]/[name]-[hash].[ext]',
                manualChunks(id) {
                    if (id.includes("node_modules")) {
                        return id.toString().split("node_modules/")[1].split("/")[0].toString()
                    }
                },
            }
        },
        terserOptions: {
            compress: {
                keep_infinity: true,  // 防止 Infinity 被压缩成 1/0，这可能会导致 Chrome 上的性能问题
                drop_console: true,   // 生产环境去除 console
                drop_debugger: true   // 生产环境去除 debugger
            },
        },
        reportCompressedSize: false, // 启用/禁用 gzip 压缩大小报告
        chunkSizeWarningLimit: 1500, // 规定触发警告的 chunk 大小
    }
});