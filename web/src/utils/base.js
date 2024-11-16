// 用于处理公共逻辑，如跳转、返回等
import {useMenu} from "@/stores/menu.js";
import EventBus from "@/utils/eventBus.js";

export function useNavigation(router) {
    /**
     * 跳转,跳转到应用内的某个页面
     * @param url 跳转地址，前面要有斜杠/page/*
     * @param queryParams 跳转参数
     */
    const jump = (url, queryParams = {}) => {
        router.push({
            path: url,
            query: queryParams
        })
    }
    /**
     * 跳转,跳转到应用内的某个页面
     * @param url 跳转地址，前面要有斜杠/page/*
     * @param queryParams 跳转参数
     */
    const jumpTab = (url, queryParams = {}) => {
        const getTabs = useMenu().tabs.find(item => "/" + item.path === url)
        if (!getTabs) {
            const getMenu = useMenu().menu.find(item => "/" + item.path === url);
            const index = useMenu().tabs.length
            useMenu().tabs.push({
                key: useMenu().tabs.length + "",
                title: getMenu.title,
                content: getMenu.title,
                path: getMenu.path,
                id: getMenu.id,
                parent_id: getMenu.parent_id,
            })
            useMenu().tabsActiveKey = useMenu().tabs.length + ""
            EventBus.emit('setActiveKey', index + "");
            console.log(getMenu)
            EventBus.emit('setMenuCheck', getMenu);
        } else {
            useMenu().setState("tabsActiveKey", getTabs.key)
            EventBus.emit('setActiveKey', getTabs.key);
            console.log(getTabs)
            EventBus.emit('setMenuCheck', getTabs);
        }
        router.push({
            path: url,
            query: queryParams
        })
    }
    /**
     * 打开url
     * @param url
     */
    const openUrl = (url) => {
        window.open(url)
    }
    /**
     * 获取资源
     * @param url
     * @returns {*}
     */
    const getAssets = (url) => {
        const modules = import.meta.glob('../static/*', {eager: true});
        return modules["../assets/" + url].default
    };
    // 通过返回值暴露所管理的状态
    return {
        jump, openUrl, getAssets, jumpTab
    }
}