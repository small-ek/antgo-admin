import {defineStore} from 'pinia';
import {useDevice} from "@/utils/device.js";


export const useUserLoginStore = defineStore('userLogin', {
    state: () => {
        return {
            authorization: "",
            userInfo: {},
            expiresAt: "",
            deviceId: ""
        };
    },
    getters: {
        getAuthorization(state) {
            return state.authorization;
        },
        getUserInfo(state) {
            return state.userInfo;
        },
        getDeviceId: (state) => {
            if (state.deviceId === "") {
                state.deviceId = useDevice().getUniqueNumber()
            }
            return state.deviceId;
        },

        getExpiresAt(state) {
            return state.expiresAt;
        }
    },
    actions: {
        setAuthorization(authorization) {
            this.authorization = authorization;
        },
        setUserInfo(userInfo) {
            this.userInfo = userInfo;
        },
        setExpiresAt(expiresAt) {
            this.expiresAt = expiresAt;
        },
        // 登录
        login(userInfo) {
            this.setUserInfo(userInfo.user);
            this.setAuthorization(userInfo.token);
            this.setExpiresAt(userInfo.expiresAt);
        },
        // 退出
        logout() {
            // 退出登录
            this.setAuthorization("");
            this.setUserInfo({});
            this.setExpiresAt("");
        }
    },
    persist: {
        // 开启持久化
        enabled: true,
        // 选择存储方式和内容
        strategies: [
            {
                storage: localStorage,
                paths: ['authorization', 'userInfo', 'expiresAt', 'deviceId']
            }
        ]
    }
});