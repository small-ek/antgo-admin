import {defineStore} from 'pinia';


export const useLayout = defineStore('useLayout', {
    state: () => {
        return {
            // 布局模式 (纵向：vertical | 经典：classic | 横向：transverse | 分栏：columns)
            layout: "vertical",
            // element 组件大小
            assemblySize: "default",
            // 当前系统语言
            language: null,
            // 当前页面是否全屏
            maximize: false,
            // 主题颜色
            primary: "#009688",
            // 深色模式
            isDark: false,
            // 灰色模式
            isGrey: false,
            // 色弱模式
            isWeak: false,
            // 侧边栏反转
            asideInverted: false,
            // 头部反转
            headerInverted: false,
            // 折叠菜单
            isCollapse: false,
            // 菜单手风琴
            accordion: true,
            // 面包屑导航
            breadcrumb: true,
            // 面包屑导航图标
            breadcrumbIcon: true,
            // 标签页
            tabs: true,
            // 标签页图标
            tabsIcon: true,
            // 页脚
            footer: true,
            //手机端是否可见
            mobileVisible: false,
            //显示区宽度
            windowWidth: window.innerWidth,
            //显示区高度
            windowHeight: window.innerHeight,

            //是否显示设置
            isSetting: false
        };
    },
    getters: {},
    actions: {
        setState(...args) {
            this.$patch({[args[0]]: args[1]});
        },
        resetData() {
            this.$reset();
        }
    },
    persist: {
        // 开启持久化
        enabled: true,
        // 选择存储方式和内容
        strategies: [
            {
                storage: localStorage,
                paths: ["windowWidth", "windowHeight"]
            }
        ]
    }
});