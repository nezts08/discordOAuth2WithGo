package models

import (
	"encoding/json"
	"strconv"
)

type CustomInt64 int64

func (ci *CustomInt64) UnmarshalJSON(data []byte) error {
	if len(data) >= 2 && data[0] == '"' && data[len(data)-1] == '"' {
		s := string(data[1 : len(data)-1])
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return err
		}
		*ci = CustomInt64(val)
		return nil
	}
	var val int64
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	*ci = CustomInt64(val)
	return nil
}

type GuildDC struct {
	ID          string      `json:"id" bson:"id"`
	Name        string      `json:"name" bson:"name"`
	Icon        string      `json:"icon" bson:"icon"`
	Owner       bool        `json:"owner" bson:"owner"`
	Permissions CustomInt64 `json:"permissions" bson:"permissions"` // Usamos nuestro tipo
	InServer    bool        `json:"in_server" bson:"in_server"`
}
