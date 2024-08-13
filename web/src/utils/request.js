import axios from 'axios';
import config from "./config.js"
import {useUserLoginStore} from "@/stores/userLogin.js";

const http = axios.create({
    baseURL: config.url,
    timeout: config.timeout,
});

// 请求拦截器
http.interceptors.request.use(
    config => {
        if (useUserLoginStore().authorization !== "") {
            config.headers['x-token'] = useUserLoginStore().authorization
        }
        config.headers['device-id'] = useUserLoginStore().getDeviceId
        return config;
    },
    error => {
        // 请求错误处理
        return Promise.reject(error);
    }
);

// 响应拦截器
http.interceptors.response.use(
    response => {
        // 对响应数据做处理，例如只返回data部分
        const data = response.data
        if (response.status !== 200) {
            console.error("请求失败，请稍后再试")
        }
        return data === undefined ? {} : data
    },
    error => {
        // 响应错误处理
        return Promise.reject(error);
    }
);

export default http;