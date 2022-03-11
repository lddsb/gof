package singleton

// Singleton 是单例模式接口，导出的
// 通过该接口可以避免GetInstance返回一个包私有类型的指针
type Singleton interface {
	foo()
}

// singleton 是单例模式类，包私有的
type singleton struct{}

func (s singleton) foo() {}

var instance *singleton

func init() {
	instance = &singleton{}
}

// GetInstance 获取实例
func GetInstance() Singleton {
	return instance
}
