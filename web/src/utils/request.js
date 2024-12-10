import axios from 'axios';
import config from "./config.js"
import {useUserLoginStore} from "@/stores/userLogin.js";
import {Message} from "@arco-design/web-vue";
import i18n from "@/utils/i18n.js";
import EventBus from "@/utils/eventBus.js";

const http = axios.create({
    baseURL: config.url,
    timeout: config.timeout,
});
let loadingTimer;

// 显示loading
const showLoading = () => {
    if (loadingTimer) {
        clearTimeout(loadingTimer);
    }
    loadingTimer = setTimeout(() => {
        EventBus.emit('setLoading', true);
    }, 200);
};

// 关闭loading
const closeLoading = () => {
    if (loadingTimer) {
        clearTimeout(loadingTimer);
        loadingTimer = null;
    }
    EventBus.emit('setLoading', false);
};

// 请求拦截器
http.interceptors.request.use(
    config => {
        showLoading()
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
        closeLoading()
        // 对响应数据做处理，例如只返回data部分
        const data = response.data


        const msg = i18n.global.t('tip.' + data.message) !== 'tip.' + data.message ? i18n.global.t('tip.' + data.message) : data.message;

        if (data.code !== 0 && msg) {
            Message.error(msg)
        }


        return data === undefined ? {} : data
    },
    res => {
        closeLoading()
        const data = res.response.data
        const msg = i18n.global.t('tip.' + data.message) !== 'tip.' + data.message ? i18n.global.t('tip.' + data.message) : data.message;
        if (res.status === 500) {
            Message.error("请求失败,请重新尝试一下")
        } else if (res.status === 401 && msg) {
            Message.error(msg)
            useUserLoginStore().logout()
            setTimeout(() => {
                window.location.reload()
            }, 500)
        } else if (res.status !== 200 && msg) {
            Message.error(msg)
        }
        // 响应错误处理
        return Promise.reject(error);
    }
);

export default http;