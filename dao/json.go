package dao

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type JsonEntity[T any] struct {
	entity T
}

func (j *JsonEntity[T]) Scan(value any) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var result T
	err := json.Unmarshal(bytes, &result)
	*j = JsonEntity[T]{entity: result}
	return err
}

func (j JsonEntity[T]) Value() (driver.Value, error) {
	return json.Marshal(j.entity)
}

func (j JsonEntity[T]) Entity() T {
	return j.entity
}

type StringArray []string

func (StringArray) GormDataType() string {
	return "JSON"
}

func (s *StringArray) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s StringArray) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}

type UnitArray []uint

func (UnitArray) GormDataType() string {
	return "JSON"
}

func (s *UnitArray) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s UnitArray) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}
