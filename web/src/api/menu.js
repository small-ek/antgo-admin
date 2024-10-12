import http from "@/utils/request.js";

/**
 * getMenu 获取菜单
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const getMenu = async (parent_id=0) => {
    return http.get('sys-menu',{params:{page_size:1000}});
}
