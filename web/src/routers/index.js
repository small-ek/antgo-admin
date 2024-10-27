import {createRouter, createWebHistory} from "vue-router";
import routes from "./basic_routes.js"
import NProgress from '@/utils/nprogress';
import {useUserLoginStore} from "@/stores/userLogin.js";
import config from "@/utils/config.js";
import {getMenu} from "@/api/menu.js";

//创建路由
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: routes
});


//setDynamicRouter 设置动态路由
const setDynamicRouter = () => {
    const modules = import.meta.glob(['../views/*/*.vue', '../views/*/*/*.vue'])
    getMenu().then((res) => {
        console.log(router)
        const children = []
        for (let i = 0; i < res.data.items.length; i++) {
            const row = res.data.items[i]
            if (row.path !== "index" && row.path !== "/" && row.path !== "") {
                router.addRoute("admin", {
                    path: "/" + row.path,
                    component: () => import('@/views/home/index.vue'),
                    meta: {
                        title: row.title,
                        keywords: row.title,
                        description: row.title
                    },
                })

            }
        }
        router.addRoute("admin", {
            path: "/tt",
            component: () => import('@/views/home/index.vue'),
            meta: {},
        })

        console.log(router.getRoutes())
        console.log(routes);
    });
}
setDynamicRouter()
//全局前置守卫
router.beforeEach(async (to, from, next) => {
    NProgress.start();
    // setDynamicRouter()
    const isLogin = useUserLoginStore().getAuthorization === ""
    const isLoginUrl = config.noLoginUrls.includes(to.path)
    if (isLogin && !isLoginUrl) {
        next({path: '/login', query: {redirect: to.fullPath}});
    } else {
        next();
    }
    return true
})

//全局后置钩子
router.afterEach((to) => {
    const {title, keywords, description} = to.meta

    if (title) {
        document.title = title
    }

    if (keywords) {
        const metaKeywords = document.querySelector('meta[name="keywords"]')
        if (metaKeywords) {
            metaKeywords.content = keywords
        }
    }

    if (description) {
        const metaDescription = document.querySelector('meta[name="description"]')
        if (metaDescription) {
            metaDescription.content = description
        }
    }

    NProgress.done();

});

router.onError(error => {
    NProgress.done();
    console.warn("路由错误", error.message);
});

export default router;