import axios from 'axios';
import config from "./config.js"
import {useUserLoginStore} from "@/stores/userLogin.js";
import {Message} from "@arco-design/web-vue";
import i18n from "@/utils/i18n.js";

const http = axios.create({
    baseURL: config.url,
    timeout: config.timeout,
});

// 请求拦截器
http.interceptors.request.use(
    config => {
        if (useUserLoginStore().authorization !== "") {
            config.headers['Authorization'] = "Bearer " + useUserLoginStore().authorization
        }
        config.headers['Device-ID'] = useUserLoginStore().getDeviceId
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


        const msg = i18n.global.t('tip.' + data.message) !== 'tip.' + data.message ? i18n.global.t('tip.' + data.message) : data.message;

        if (data.code !== 0 && data.error) {
            const errorMsg = i18n.global.t('tip.' + data.error) !== 'tip.' + data.error ? i18n.global.t('tip.' + data.error) : data.error;
            if (errorMsg) {
                Message.error(errorMsg)
                return
            }
        }
        if (data.code !== 0 && msg) {
            Message.error(msg)
        }


        return data === undefined ? {} : data
    },
    response => {
        const data = response.data
        const errorMsg = i18n.global.t('tip.' + data.error)
        const msg = i18n.global.t('tip.' + data.message)
        if (response.status === 500) {
            Message.error("接口请求失败，请重新尝试一下")
        } else if (response.status === 401 && errorMsg) {
            Message.error(msg)
        } else if (response.status !== 200 && msg) {
            Message.error(msg)
        }
        // 响应错误处理
        return Promise.reject(error);
    }
);

export default http;