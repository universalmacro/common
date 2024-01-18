package utils

import (
	"encoding/json"
	"errors"
	"fmt"
)

func ScanJson[T any](value any, j *T) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var result T
	err := json.Unmarshal(bytes, &result)
	*j = T(result)
	return err
}
