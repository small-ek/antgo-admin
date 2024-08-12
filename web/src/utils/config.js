let config = {
    url: "",//请求地址
    timeout: 10000,
    token_expiration_time: 60 * 60 * 24 * 7,//token过期时间
}

if (process.env.NODE_ENV === 'development') {
    // 开发环境
    config.url = 'http://127.0.0.1:9001'
} else {
    // 生产环境
    config.url = 'http://127.0.0.1:9001'
}

export default config