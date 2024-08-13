// 用于生成浏览器唯一设备指纹

import {ClientJS} from "clientjs";

export function useDevice() {
    /**
     * 获取唯一设备号
     * @returns {*}
     */
    const getUniqueNumber = () => {
        var client = new ClientJS(); // Create A New Client Object
        return client.getFingerprint().toString();
    };


// 通过返回值暴露所管理的状态
    return {
        getUniqueNumber
    }
}