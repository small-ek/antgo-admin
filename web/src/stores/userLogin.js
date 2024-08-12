import {defineStore} from 'pinia';

export const useUserLoginStore = defineStore('userLogin', {
    state: () => {
        return {
            authorization: uni.getStorageSync("authorization") || "",
            userInfo: uni.getStorageSync("userInfo") || {},
            expiresAt: uni.getStorageSync("expiresAt") || "",
            deviceId: uni.getStorageSync("deviceId") || "",
            isLogin: uni.getStorageSync("authorization") !== "",
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
                state.deviceId = uni.getSystemInfoSync()?.deviceId;
                uni.setStorageSync('deviceId', state.deviceId)
            }
            return state.deviceId;
        },
        getIsLogin(state) {
            return state.authorization !== "";
        },
        getExpiresAt(state) {
            return state.expiresAt;
        }
    },
    actions: {
        setAuthorization(authorization) {
            this.authorization = authorization;
            uni.setStorageSync("authorization", authorization);
        },
        setUserInfo(userInfo) {
            this.userInfo = userInfo;
            uni.setStorageSync("userInfo", userInfo);
        },
        setExpiresAt(expiresAt) {
            this.expiresAt = expiresAt;
            uni.setStorageSync("expiresAt", expiresAt);
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
});