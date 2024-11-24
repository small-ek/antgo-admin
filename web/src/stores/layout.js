import {defineStore} from 'pinia';


export const useLayout = defineStore('useLayout', {
    state: () => {
        return {
            //是否显示设置
            setting: true,
            // 布局模式 (垂直：vertical | 经典：classic | 横向：transverse | 分栏：columns)
            layout: "vertical",
            // 侧边栏深色
            isDarkSidebar: true,
            // 菜单手风琴
            isAccordion: false,
            // 侧边栏宽度
            sidebarWidth: 260,
            // 头部深色
            isDarkHeader: false,
            // 面包屑导航
            isBreadcrumb: true,
            // 显示多语言
            isLanguage: true,
            // 显示全屏
            isFullScreen: true,
            // 显示刷新
            isRefresh: true,
            //显示搜索
            isSearch: true,
            // 头部模式 (固定：fixed | 静止：static | 自适应：adaptive)
            header: "adaptive",
            // 显示标签栏
            isTabs: true,
            // 标签栏是否显示图标
            isTabsIcon: true,
            // 标签是否可拖拽
            isTabsDraggable: true,
            // Tabs样式('line' | 'card' | 'card-gutter' | 'text' | 'rounded' | 'capsule')
            tabsType: 'card-gutter',
            // 显示底部栏
            isFooter: true,
            // 底部文案
            footerText: "技术支持 © 2021",
            // 主题模式 (auto | dark | light)
            theme: "auto",
            // 主题颜色
            primary: "rgb(24, 144, 255)",
            // 色弱模式
            isColorBlind: false,
            // 灰色模式
            isGrey: false,
            //显示区宽度
            windowWidth: window.innerWidth,
            //显示区高度
            windowHeight: window.innerHeight,
            //是否显示设置面板
            showSetting: false,
            //菜单收缩
            isCollapsed: false,
            //显示手机菜单
            showMobileMenu: false,
            //显示标签
            showHeader: true,
            //头部是否固定
            isFixedHeader: true,
        };
    },
    getters: {},
    actions: {
        setState(...args) {
            this.$patch({[args[0]]: args[1]});
        },
        setLayout(value) {
            this.setState('sidebarWidth', 260);
            if (value === 'vertical') {
                this.setState('isDarkSidebar', true);
                this.setState('isDarkHeader', false);
            }
            if (value === 'classic') {
                this.setState('isDarkSidebar', false);
                this.setState('isDarkHeader', true);
            }
            if(value==='transverse'){
                this.setState('isDarkSidebar', true)
                this.setState('isDarkHeader', true);
            }
            if(value==='columns'){
                this.setState('isDarkSidebar', false)
                this.setState('isDarkHeader', false);
                this.setState('sidebarWidth', 220);
            }

            this.setState('isFixedHeader', true)
            this.$patch({'layout': value});
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
                paths: [
                    "setting",
                    "layout",
                    "isDarkSidebar",
                    "isAccordion",
                    "sidebarWidth",
                    "isDarkHeader",
                    "isBreadcrumb",
                    "isLanguage",
                    "isFullScreen",
                    "isRefresh",
                    "isSearch",
                    "header",
                    "isTabs",
                    "isTabsIcon",
                    "isTabsDraggable",
                    "isFooter",
                    "footerText",
                    "theme",
                    "primary",
                    "isColorBlind",
                    "isGrey",
                    "windowWidth",
                    "windowHeight",
                    "showSetting",
                    "isFixedHeader"
                ]
            }
        ]
    }
});