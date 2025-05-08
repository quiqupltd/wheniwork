package wheniwork

import (
	"encoding/json"
	"time"
)

const WIWTimeFormat = "Mon, 02 Jan 2006 15:04:05 -0700"

type WIWTime time.Time

func (t WIWTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(WIWTimeFormat))
}

func (t *WIWTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsedTime, err := time.Parse(WIWTimeFormat, s)
	if err != nil {
		return err
	}

	*t = WIWTime(parsedTime)
	return nil
}
