package singleton

import "sync"

type Singleton[T any] interface {
	Get() *T
}

type EagerSingleton[T any] struct {
	instance *T
}

func (s *EagerSingleton[T]) Get() *T {
	return s.instance
}

type Mode int

const (
	Eager Mode = iota
	Lazy
)

func SingletonFactory[T any](constructor func() *T, mode Mode) Singleton[T] {
	switch mode {
	case Eager:
		return NewEagerSingleton(constructor)
	case Lazy:
		return NewLazySingleton(constructor)
	}
	panic("create singleton failed")
}

func NewEagerSingleton[T any](constructor func() *T) *EagerSingleton[T] {
	var singleton EagerSingleton[T]
	singleton.instance = constructor()
	return &singleton
}

type LazySingleton[T any] struct {
	instance    *T
	constructor func() *T
	once        sync.Once
}

func NewLazySingleton[T any](constructor func() *T) *LazySingleton[T] {
	var singleton LazySingleton[T]
	singleton.constructor = constructor
	return &singleton
}

func (s *LazySingleton[T]) Get() *T {
	s.once.Do(
		func() {
			s.instance = s.constructor()
		})
	return s.instance
}
