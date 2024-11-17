//用于树形结构的操作
import {pinyin} from "pinyin-pro";

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
    /**
     * 获取所有子节点
     * @param items 数据
     * @returns {*}
     */
    const subTree = (items) => {
        return items
            .filter(item => item.path !== "" &&  item.component !== "").map(item => ({
                ...item,
                pinyin: pinyin(item.title, {toneType: "none"}).replace(/\s+/g, ''),
            }));
    };
    /**
     * findKeys 查找菜单树中的默认展开项和默认选中项
     * @param tree
     * @param path
     * @returns {{defaultOpenKeys: *[], defaultSelectedKey: null}}
     */
    const findKeys = (tree, path) => {
        let openKeys = [];
        let selectedKeys = [];

        const traverse = (nodes, parentKey = null) => {
            for (const node of nodes) {
                if ("/" + node.path === path) {
                    selectedKeys.push(node.id)
                    if (parentKey) {
                        openKeys.push(parentKey);
                    }
                    return true;
                }
                if (node.children && node.children.length > 0) {
                    if (traverse(node.children, node.id)) {
                        if (!openKeys.includes(node.id)) {
                            openKeys.push(node.id);
                        }
                        return true;
                    }
                }
            }
            return false;
        };

        traverse(tree);
        return {openKeys, selectedKeys};
    };

    // 通过返回值暴露所管理的状态
    return {
        buildTree, findKeys, subTree
    }
}