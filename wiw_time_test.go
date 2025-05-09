package wheniwork_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/quiqupltd/wheniwork"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type ExampleData struct {
	Time wheniwork.WIWTime `json:"time"`
}

func Test_WIWTime(t *testing.T) {
	exampleTimeString := "Fri, 09 May 2025 01:20:22 -0500"
	exampleTime, err := time.Parse(wheniwork.WIWTimeFormat, exampleTimeString)
	// convert to WIWTime
	exampleWIWTime := wheniwork.WIWTime(exampleTime)
	require.NoError(t, err)

	t.Run("should unmarshal WIWTime", func(t *testing.T) {
		var data ExampleData
		err := json.Unmarshal([]byte(`{"time": "`+exampleTimeString+`"}`), &data)

		require.NoError(t, err)
		assert.Equal(t, exampleWIWTime, data.Time)
	})

	t.Run("should marshal WIWTime", func(t *testing.T) {
		data := ExampleData{Time: wheniwork.WIWTime(exampleTime)}
		jsonData, err := json.Marshal(data)

		require.NoError(t, err)
		assert.Equal(t, `{"time":"Fri, 09 May 2025 01:20:22 -0500"}`, string(jsonData))
	})

	t.Run("should return blank if the time cannot Unmarshal", func(t *testing.T) {
		var data ExampleData
		err := json.Unmarshal([]byte(`{"time": "Fri, 09 May -001 01:20:22 -0341"}`), &data)

		require.NoError(t, err)
		assert.Equal(t, time.Time{}, time.Time(data.Time))
	})
}
