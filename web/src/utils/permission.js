import {useMenu} from "@/stores/menu.js";
import router from "@/routers/index.js";
import {getMenu} from "@/api/menu.js";
import {useTree} from "@/utils/tree.js";

/**
 * initRouter 刷新初始化路由
 * @returns {Promise<void>}
 */
export const initRouter =  async () => {
    const modules = import.meta.glob(['../views/*/*.vue'])
    const menu =  await useMenu().menu
    if (menu.length > 0) {
        for (let i = 0; i < menu.length; i++) {
            const row = menu[i]
            if (row.path !== "/" && row.path !== "" && row.component) {
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
}

/**
 * setMenu 设置路由
 * @returns {Promise<void>}
 */
export const setMenu =  () => {
    return getMenu().then(res => {
        useMenu().setState("menu", res.data.items)
        useMenu().setState("menuTree", useTree().buildTree(res.data.items))
        initRouter()
    })
}