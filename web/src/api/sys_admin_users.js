import http from "@/utils/request.js";

/**
 * 获取列表
 * @param current_page 当前页
 * @param page_size 每页条数
 * @param filter_map 过滤条件
 * @param order 排序字段
 * @param desc 是否降序
 * @returns {Promise<axios.AxiosResponse<any>>} 返回列表
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
 * @param id 主键
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const getSysAdminUsers = async (id) => {
    return http.get('sys-admin-users/' + id);
}

/**
 * deleteSysAdminUsers 删除
 * @returns {Promise<axios.AxiosResponse<any>>}
 * @param ids 删除的主键
 */
export const deleteSysAdminUsers = async (ids = []) => {
    return http.delete('sys-admin-users', {
        data: {
            ids: ids
        }
    });
}

/**
 * createSysAdminUsers 创建
 * @param data 创建数据
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const createSysAdminUsers = async (data) => {
    return http.post('sys-admin-users', data);
}

/**
 * updateSysAdminUsers 更新
 * @param data 更新数据
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const updateSysAdminUsers = async (data) => {
    return http.put('sys-admin-users/' + data.id, data);
}