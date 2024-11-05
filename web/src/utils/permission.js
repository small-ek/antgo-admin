import {useMenu} from "@/stores/menu.js";
import router from "@/routers/index.js";

export const initRouter = async () => {
    const modules = import.meta.glob(['../views/*/*.vue'])
    const menu = await useMenu().menu
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

}