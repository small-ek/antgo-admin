//用于树形结构的操作

export function useTree() {
    /**
     * 一维转树形结构
     * @param items 数据
     * @param parentId 父级id
     * @returns {*}
     */
    const buildTree = (items, parentId = 0) => {
        return items
            .filter(item => item.parent_id === parentId)
            .map(item => ({
                ...item,
                children: buildTree(items, item.id)
            }));
    };


// 通过返回值暴露所管理的状态
    return {
        buildTree
    }
}