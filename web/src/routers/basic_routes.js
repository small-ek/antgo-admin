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

        ],
    },

]


export default routes