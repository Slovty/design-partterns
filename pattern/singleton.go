package pattern

import "sync"

// 单例模式 全局只有一个实例  redis
// 适用于实例可复用，节约对象频繁创建销毁的开销
// 2种模式
// 饿汉模式，在包加载时就进行初始化
// 懒汉模式， 在使用时才进行初始化

type Singleton struct {
}

var singleton *Singleton
var once sync.Once

func GetSingleton() *Singleton {
	once.Do(func() {
		singleton = new(Singleton)
	})
	return singleton
}
