import axios from 'axios';
import config from "./config.js"
import {useUserLoginStore} from "@/stores/userLogin.js";

const http = axios.create({
    baseURL: config.url, // 替换为你的API基地址
    timeout: config.timeout, // 请求超时时间
});

// 请求拦截器
http.interceptors.request.use(
    config => {
        // 可以在这里添加例如token等请求头
        if (useUserLoginStore().authorization !== "") {
            // 可以在此通过vm引用vuex中的变量，具体值在vm.$store.state中
            config.header['x-token'] = useUserLoginStore().authorization
        }
        config.header['device-id'] = useUserLoginStore().getDeviceId
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
        return response.data;
    },
    error => {
        // 响应错误处理
        return Promise.reject(error);
    }
);

export default http;