package convert

import "encoding/json"

func ConvertByJSON[F, T any](from F, to *T) error {
	j, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(j, to)
}
