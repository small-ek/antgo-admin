import http from "@/utils/request.js";

/**
 * getCaptcha 获取验证码
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const getCaptcha = async () => {
    return http.get('captcha');
}

/**
 * login 登录
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const login = async (data) => {
    return http.post('login', data);
}