/**
 * @description hex颜色转rgb颜色
 * @param {String} str 颜色值字符串
 * @returns {String} 返回处理后的颜色值
 */
export function hexToRgb(str) {
    let hexs = "";
    let reg = /^\#?[0-9A-Fa-f]{6}$/;
    if (!reg.test(str)) return console.error("输入错误的hex");
    str = str.replace("#", "");
    hexs = str.match(/../g);
    for (let i = 0; i < 3; i++) hexs[i] = parseInt(hexs[i], 16);
    return hexs;
}

/**
 * @description 将RGB颜色和透明度转换为带透明度的Hex颜色
 * @param {Number} r 代表红色 (0-255)
 * @param {Number} g 代表绿色 (0-255)
 * @param {Number} b 代表蓝色 (0-255)
 * @param {Number} a 代表透明度 (0-1)
 * @returns {String} 返回带透明度的Hex颜色值
 */
export function rgbToHex(r, g, b, a) {
    const toHex = (value) => {
        const hex = Math.round(value).toString(16);
        return hex.length === 1 ? '0' + hex : hex;
    };

    const alpha = Math.round(a * 255);
    return `#${toHex(r)}${toHex(g)}${toHex(b)}${toHex(alpha)}`;
}

/**
 * @description 加深颜色值
 * @param rgb
 * @param {Number} level 加深的程度，限0-1之间
 * @returns {String} 返回处理后的颜色值
 */
export function getDarkColor(rgb, level) {
    for (let i = 0; i < 3; i++) rgb[i] = Math.round(20.5 * level + rgb[i] * (1 - level));
    if (rgb[3] === undefined) {
        return `${rgb[0]},${rgb[1]},${rgb[2]}`;
    }
    return `${rgb[0]},${rgb[1]},${rgb[2]},${rgb[3]}`;
}

/**
 * @description 变浅颜色值
 * @param rgb
 * @param {Number} level 加深的程度，限0-1之间
 * @returns {String} 返回处理后的颜色值
 */
export function getLightColor(rgb, level) {
    for (let i = 0; i < 3; i++) rgb[i] = Math.round(255 * level + rgb[i] * (1 - level));

    if (rgb[3] === undefined) {
        return `${rgb[0]},${rgb[1]},${rgb[2]}`;
    }
    return `${rgb[0]},${rgb[1]},${rgb[2]},${rgb[3]}`;
}
