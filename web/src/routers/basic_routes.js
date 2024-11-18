// 基础路由
const routes = [
    {
        name: "login",
        path: '/login',
        component: () => import('@/views/login/index.vue'),
        meta: {
            title: '后台登录',
            keywords: '后台登录',
            description: '后台登录'
        },
    },
    {
        path: '/',
        name: 'admin',
        component: () => import('@/layout/index.vue'),
        children: [
            {
                name: "settings",
                path: '/settings',
                component: () => import('@/views/settings/index.vue'),
                meta: {
                    title: '个人设置',
                    keywords: '个人设置',
                    description: '个人设置'
                },
            },
        ],
    },

]


export default routes