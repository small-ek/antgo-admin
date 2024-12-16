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
export const getSysMenuList = async (current_page, page_size = 1000, filter_map, order = [], desc = []) => {
    return http.get('sys-menu', {
        params: {
            current_page: current_page,
            page_size: page_size,
            filter_map: JSON.stringify(filter_map),
            order: order.length > 0 ? order : ['sort', 'id'],
            desc: desc.length > 0 ? desc : [true, false],
        }
    });
}

/**
 * getSysMenu 获取详情
 * @param id 主键
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const getSysMenu = async (id) => {
    return http.get('sys-menu/' + id);
}

/**
 * deleteSysAdminUsers 删除
 * @returns {Promise<axios.AxiosResponse<any>>}
 * @param ids 删除的主键
 */
export const deleteSysMenu = async (ids = []) => {
    return http.delete('sys-menu', {
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
export const createSysMenu = async (data) => {
    return http.post('sys-menu', data);
}

/**
 * updateSysAdminUsers 更新
 * @param data 更新数据
 * @returns {Promise<axios.AxiosResponse<any>>}
 */
export const updateSysMenu = async (data) => {
    return http.put('sys-menu/' + data.id, data);
}
