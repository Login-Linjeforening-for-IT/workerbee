package db

import "encoding/json"

func (ns NullTimeTypeEnum) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.TimeTypeEnum)
	} else {
		return json.Marshal(nil)
	}
}

func (ns *NullTimeTypeEnum) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "null", `""`:
		ns.Valid = false
		return nil
	}

	var v TimeTypeEnum
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	ns.TimeTypeEnum = v
	ns.Valid = true

	return nil
}
