// 用于界面设置修改
import {useLayout} from "@/stores/layout.js";
import {getLightColor, rgbToHex} from "@/utils/color.js";

export function useTheme() {
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
            document.body.style.setProperty(`--primary-${i}`, getLightColor(colorRgb[0].split(",")  , i / 10));
        }
        // globalStore.setGlobalState("primary", val);
    }
    return {
        changePrimary
    }
}