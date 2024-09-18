// 用于界面设置修改
import {useLayout} from "@/stores/layout.js";
import {getLightColor} from "@/utils/color.js";
import throttle from "@/utils/throttle.js";

let lastScrollTop = 0;

export function useTheme() {
    // 默认主题设置
    const defaultTheme = () => {
        updateDark();
        if (useLayout().isColorBlind) {
            changeGreyOrWeak('weak', useLayout().isColorBlind)
        }
        if (useLayout().isGrey) {
            changeGreyOrWeak('grey', useLayout().isGrey)
        }
        changePrimary()
    }
    /**
     * 切换暗黑模式
     */
    const updateDark = () => {
        if (useLayout().theme === 'dark') {
            document.body.setAttribute('arco-theme', 'dark')
        }
        if (useLayout().theme === 'light') {
            document.body.setAttribute('arco-theme', 'arco-theme')
        }
        if (useLayout().theme === 'auto') {
            document.body.setAttribute('arco-theme', 'arco-theme')
        }
    };
    /**
     * changePrimary 修改主题颜色
     * @param val
     */
    const changePrimary = (val) => {
        if (!val) {
            val = useLayout().primary;
        }
        // 计算主题颜色变化
        const colorRgb = val.match(/\d{1,3}, \d{1,3}, \d{1,3}(?:, \d?\.?\d+)?/);
        document.body.style.setProperty("--primary-6", colorRgb[0]);
        for (let i = 5; i >= 1; i--) {
            document.body.style.setProperty(`--primary-${i}`, getLightColor(colorRgb[0].split(","), i / 10));
        }
    }

    /**
     * 灰色或者弱色切换
     * @param type
     * @param value
     */
    const changeGreyOrWeak = (type, value) => {
        const body = document.body
        if (!value) return body.removeAttribute("style");
        const styles = {
            grey: "filter: grayscale(1)",
            weak: "filter: invert(80%)"
        };
        body.setAttribute("style", styles[type]);
    };

    const handleScroll = (event) => {
        throttle(() => {
            if (useLayout().header !== 'adaptive') return;
            const container = event.target;
            const scrollTop = container.scrollTop;
            if (scrollTop > lastScrollTop) {
                useLayout().setState("isFixedHeader", false)
            } else {
                useLayout().setState("isFixedHeader", true)
            }
            lastScrollTop = scrollTop <= 0 ? 0 : scrollTop; // For Mobile or negative scrolling
        }, 100);
    }
    const handleResize = () => {
        useLayout().setState("windowWidth", window.innerWidth)
        useLayout().setState("windowHeight", window.innerHeight)
        if (useLayout().windowWidth < 768) {
            useLayout().setState("isCollapsed", false)
        } else {
            useLayout().setState("mobileVisible", false)
        }
    };
    return {
        defaultTheme, changePrimary, updateDark, changeGreyOrWeak, handleScroll, handleResize
    }
}