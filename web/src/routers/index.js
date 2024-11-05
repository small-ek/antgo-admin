import {createRouter, createWebHistory} from "vue-router";
import routes from "./basic_routes.js"
import NProgress from '@/utils/nprogress';
import {useUserLoginStore} from "@/stores/userLogin.js";
import config from "@/utils/config.js";
import {useMenu} from "@/stores/menu.js";

//创建路由
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: routes
});


//setDynamicRouter 设置动态路由
const setDynamicRouter = async () => {
    const modules = import.meta.glob(['../views/*/*.vue'])
    const menu = useMenu().menu
    if (menu.length > 0) {
        for (let i = 0; i < menu.length; i++) {
            const row = menu[i]
            if (row.path !== "index" && row.path !== "/" && row.path !== "" && row.component) {
                if (!router.hasRoute(row.path)) {
                    router.addRoute("admin", {
                        name: row.path,
                        path: "/" + row.path,
                        component: () => modules[`../${row.component}`](),
                        meta: {
                            title: row.title,
                            keywords: row.title,
                            description: row.title
                        },
                    })
                }


            }
        }
    }
    // getMenu().then(res=>{
    //     useMenu().setState("menu",res.data.items)
    //     // const children = []
    //     for (let i = 0; i < res.data.items.length; i++) {
    //         const row = res.data.items[i]
    //         if (row.path !== "index" && row.path !== "/" && row.path !== "") {
    //             router.addRoute("admin", {
    //                 path: "/" + row.path,
    //                 component: () => modules[`../${row.component}`](),
    //                 meta: {
    //                     title: row.title,
    //                     keywords: row.title,
    //                     description: row.title
    //                 },
    //             })
    //
    //         }
    //     }
    // })
}

// 全局前置守卫
router.beforeEach(async (to, from, next) => {
    NProgress.start();
    const isLogin = useUserLoginStore().getAuthorization === "";
    const isLoginUrl = config.noLoginUrls.includes(to.path);

    if (isLogin && !isLoginUrl) {
        next({path: '/login', query: {redirect: to.fullPath}});
    } else {
        next();
    }

    NProgress.done();
    return {...to, replace: true}
});

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