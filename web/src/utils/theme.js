// 用于界面设置修改
import {useLayout} from "@/stores/layout.js";
import {getLightColor} from "@/utils/color.js";

export function useTheme() {
    /**
     * 切换暗黑模式
     */
    const updateDark = () => {
        console.log(111)
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
        // document.documentElement.style.setProperty(
        //     "--el-color-primary-dark-2",
        //     isDark.value ? `${getLightColor(val, 0.2)}` : `${getDarkColor(val, 0.3)}`
        // );
        for (let i = 5; i >= 1; i--) {
            document.body.style.setProperty(`--primary-${i}`, getLightColor(colorRgb[0].split(","), i / 10));
        }
        // globalStore.setGlobalState("primary", val);
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
        // const propName = type === "grey" ? "isWeak" : "isGrey";

    };
    return {
        changePrimary, updateDark, changeGreyOrWeak
    }
}