import http from "@/utils/request.js";

/**
 * getSysAdminUsersList 获取列表
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const getSysAdminUsersList = async (current_page, page_size, filter_map, order = [], desc = []) => {
    return http.get('sys-admin-users', {
        params: {
            current_page: current_page,
            page_size: page_size,
            filter_map: JSON.stringify(filter_map),
            order: order.length > 0 ? order : ['id'],
            desc: desc.length > 0 ? desc : [true],
        }
    });
}

/**
 * getSysAdminUsers 获取详情
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const getSysAdminUsers = async (id) => {
    return http.get('sys-admin-users/' + id);
}

/**
 * deleteSysAdminUsers 删除
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const deleteSysAdminUsers = async (id) => {
    return http.delete('sys-admin-users/' + id);
}

/**
 * createSysAdminUsers 创建
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const createSysAdminUsers = async (data) => {
    return http.post('sys-admin-users', data);
}

/**
 * updateSysAdminUsers 修改
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const updateSysAdminUsers = async (data) => {
    return http.put('sys-admin-users/' + data.id, data);
}