package design_pattern

import (

    "sync"

)

// Singleton 饿汉式单例
type Singleton struct{}

var singleton *Singleton

func init() {
    singleton = &Singleton{}
}

// GetInstance 获取实例
func GetInstance() *Singleton {
    return singleton
}


var (
    lazySingleton *Singleton
    once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() *Singleton {
    if lazySingleton == nil {
        once.Do(func() {
            lazySingleton = &Singleton{}
        })
    }

    return lazySingleton
}
