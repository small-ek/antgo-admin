// 用于处理CSV文件的工具函数
import screenfull from 'screenfull';
import {ref} from 'vue';
// 全屏状态
const screenFullStatus = ref(false);
// 执行全屏
const onScreenFull = (idName = '') => {
    if (screenfull.isEnabled) {
        console.log(idName)
        if (idName) {
            const element = document.getElementById(idName);
            screenfull.request(element);
        } else {
            screenfull.toggle();
        }


        // 在全屏状态切换时监听
        screenfull.on("change", () => {
            screenFullStatus.value = screenfull.isFullscreen;
        });
    } else {
        console.log('当前浏览器不支持全屏');
    }
};

export {onScreenFull, screenFullStatus};
