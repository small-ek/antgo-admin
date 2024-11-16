import {defineStore} from 'pinia';


export const useMenu = defineStore('useMenu', {
    state: () => {
        return {
            menu: [],
            menuTree: [],
            subMenu: [],
            tabs: [{
                key: '0',
                title: '首页',
                content: '首页',
                path: "index",
                id: 1,
                parent_id: 1,
            }],
            tabsActiveKey: '0',
        };
    },
    getters: {},
    actions: {
        setState(...args) {
            this.$patch({[args[0]]: args[1]});
        },
    },
    persist: {
        // 开启持久化
        enabled: true,
        // 选择存储方式和内容
        strategies: [
            {
                storage: localStorage,
                paths: ['menu', 'menuTree', 'tabs', 'tabsActiveKey', 'subMenu'],
            }
        ]
    }
});