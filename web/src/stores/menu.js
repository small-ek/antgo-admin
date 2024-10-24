import {defineStore} from 'pinia';


export const useMenu = defineStore('useMenu', {
    state: () => {
        return {
            router: [],
            menu: [],
            menuTree: []
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
                paths: ['router', 'menu','menuTree']
            }
        ]
    }
});