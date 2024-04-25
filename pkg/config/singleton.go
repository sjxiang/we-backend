package config

import ( 
	"sync"
)

// 单例

var (
	instance *Config
	once     sync.Once
)

func GetInstance() *Config {
	once.Do(func() {
		var err error
		if instance == nil {
			instance, err = LoadConfig() // not thread safe
			if err != nil {
				panic(err)
			}
		}
	})
	
	return instance
}