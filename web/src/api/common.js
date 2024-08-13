import http from "@/utils/request.js";

/**
 * getCaptcha 获取验证码
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const getCaptcha = async () => {
    return http.get('captcha');
}