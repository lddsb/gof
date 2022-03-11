package singleton

import "sync"

var (
	lazySingleton *singleton
	once          = &sync.Once{}
)

// GetLazyInstance 懒汉式
func GetLazyInstance() Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &singleton{}
		})
	}

	return lazySingleton
}
