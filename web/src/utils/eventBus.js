// 用于子组件和子组件之间的通信
import mitt from 'mitt'

const EventBus = mitt()
export default EventBus

//TODO: 不建议大量使用，会导致代码难以维护事件总线

// 发布一个事件
//bus.emit('foo', { a: 'b' })

// 订阅一个具体的事件
//bus.on('foo', (e) => console.log('foo', e))

// 订阅所有事件
//bus.on('*', (type, e) => console.log(type, e))

// 取消订阅同名事件
//bus.off('foo')

// 取消所有事件
//bus.all.clear()