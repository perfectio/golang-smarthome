package xiaomi_cube

import "encoding/json"

type Event struct {
	Action         string  `json:"action"`
	ActionFromSide int     `json:"action_from_side"`
	ActionSide     int     `json:"action_side"`
	ActionToSide   int     `json:"action_to_side"`
	ActionAngle    float32 `json:"action_angle"`
	Battery        int     `json:"battery"`
	Linkquality    int     `json:"linkquality"`
	Side           int     `json:"side"`
	Voltage        int     `json:"voltage"`
}

func (c *Event) LoadFromMap(m map[string]interface{}) error {
	data, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(data, c)
	}
	return err
}
