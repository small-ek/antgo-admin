import http from "@/utils/request.js";

/**
 * getSysAdminUsersList 获取列表
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const getSysAdminUsersList = async () => {
    return http.get('sys-admin-users');
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
export const createSysAdminUsers = async (id) => {
    return http.post('sys-admin-users/' + id);
}

/**
 * updateSysAdminUsers 修改
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const updateSysAdminUsers = async (data) => {
    return http.put('sys-admin-users/' + data.id, data);
}