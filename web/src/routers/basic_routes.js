// 基础路由
const routes = [
    {
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
        name: 'index',
        component: () => import('@/layout/index.vue'),
        children: [
            {
                path: '/',
                component: () => import('@/views/user/index.vue'),
                meta: {
                    title: '关于我们33',
                    keywords: '关键词44, 关键词55',
                    description: '关于我们描述4666'
                },
            },
        ],
    },
    {
        path: '/user',
        component: () => import('@/views/user/index.vue'),
        meta: {
            title: '关于我们33',
            keywords: '关键词44, 关键词55',
            description: '关于我们描述4666'
        },
    },
]


export default routes