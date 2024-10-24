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
                path: 'test',
                component: () => import('@/views/home/index.vue'),
                meta: {
                    title: '首页',
                    keywords: '首页',
                    description: '首页'
                },
            }
        ],
    },

]


export default routes