import {createRouter, createWebHistory} from "vue-router";
import routes from "./basic_routes"
import NProgress from '../utils/nprogress';

//创建路由
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: routes
});


//全局前置守卫
router.beforeEach(async (to, from) => {
    NProgress.start();
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