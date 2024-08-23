//用于编码国际化的文件，这里使用了vue-i18n，首先引入vue-i18n，然后引入中英文的json文件，然后创建i18n实例，设置locale和fallbackLocale，最后导出i18n实例。
import {createI18n} from 'vue-i18n';
import zh from '@/lang/zh-cn.json';
import en from '@/lang/en.json';

const messages = {
    en,
    zh,
};
const language = (navigator.language || 'zh').toLocaleLowerCase(); // 这是获取浏览器的语言

const i18n = createI18n({
    legacy: false,
    locale: localStorage.getItem('lang') || language.split('-')[0] || 'zh', // 首先从缓存里拿，没有的话就用浏览器语言，
    fallbackLocale: 'en', // 设置备用语言
    messages,
});

export default i18n;