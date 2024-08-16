// 用于处理公共逻辑，如跳转、返回等
export function useNavigation(router) {
    /**
     * 跳转,跳转到应用内的某个页面
     * @param url 跳转地址，前面要有斜杠/page/*
     * @param queryParams 跳转参数
     */
    const jump = (url, queryParams = {}) => {
        router.push({
            path: url,
            query: queryParams
        })
    }

    /**
     * 打开url
     * @param url
     */
    const openUrl = (url) => {
        window.open(url)
    }
    // 通过返回值暴露所管理的状态
    return {
        jump, openUrl
    }
}