package singleton

import "sync"

func LazySingleton[T any](constructor func() *T) func() *T {
	var once sync.Once
	var instance *T
	return func() *T {
		once.Do(
			func() {
				instance = constructor()
			})
		return instance
	}
}

func EagerSingleton[T any](constructor func() *T) func() *T {
	instance := constructor()
	return func() *T {
		return instance
	}
}
