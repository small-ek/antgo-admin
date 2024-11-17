import {useLayout} from "@/stores/layout.js";

//生产唯一key
export function generateUniqueKey() {
    const randomPart = Math.floor(Math.random() * 1000); // 3位随机数
    const timePart = Date.now().toString().slice(-5);   // 时间戳后5位
    return Number(timePart + randomPart);
}

//弹窗宽度
export function modalWidth(width) {
    return useLayout().windowWidth > 768 ? width : '95%'
}