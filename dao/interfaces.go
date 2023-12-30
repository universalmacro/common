package dao

type Dao[T any] interface {
	Create(dest T)
}

type Condition interface {
	Mixin()
}
