package common

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDateValid(t *testing.T) {
	d := NewDate(2024, 11, 22)
	assert.True(t, d.Valid)
	assert.Equal(t, 2024, d.Year)
	assert.Equal(t, time.November, d.Month)
	assert.Equal(t, 22, d.Day)
}

func TestDateInvalid(t *testing.T) {
	d := NewDate(0, 0, 0)
	assert.False(t, d.Valid)
	assert.Equal(t, 0, d.Year)
	assert.Equal(t, time.Month(0), d.Month)
	assert.Equal(t, 0, d.Day)
}

func TestDateString(t *testing.T) {
	d := NewDate(2024, 11, 22)
	assert.Equal(t, "2024-11-22", d.String())
}

func TestDateMarshalJSON(t *testing.T) {
	d := NewDate(2024, 11, 22)
	b, err := json.Marshal(d)
	assert.NoError(t, err)
	assert.Equal(t, `"2024-11-22"`, string(b))
}

func TestDateUnmarshalJSON(t *testing.T) {
	d := NewDate(0, 0, 0)
	err := json.Unmarshal([]byte(`"2024-11-22"`), &d)
	assert.NoError(t, err)
	assert.Equal(t, 2024, d.Year)
	assert.Equal(t, time.November, d.Month)
	assert.Equal(t, 22, d.Day)
}

func TestDateUnmarshalJSONWhenNotString(t *testing.T) {
	d := NewDate(0, 0, 0)
	err := json.Unmarshal([]byte(`12345`), &d)
	assert.ErrorIs(t, err, ErrInvalidDate)
}

func TestDateUnmarshalJSONWhenInvalidFormat(t *testing.T) {
	d := NewDate(0, 0, 0)
	err := json.Unmarshal([]byte(`"2024-11"`), &d)
	assert.ErrorIs(t, err, ErrInvalidDate)
}

func TestDateUnmarshalJSONWhenYearNotNumber(t *testing.T) {
	d := NewDate(0, 0, 0)
	err := json.Unmarshal([]byte(`"invalid-13-22"`), &d)
	assert.ErrorIs(t, err, ErrInvalidDate)
}

func TestDateUnmarshalJSONWhenMonthNotNumber(t *testing.T) {
	d := NewDate(0, 0, 0)
	err := json.Unmarshal([]byte(`"2024-invalid-22"`), &d)
	assert.ErrorIs(t, err, ErrInvalidDate)
}

func TestDateUnmarshalJSONWhenDayNotNumber(t *testing.T) {
	d := NewDate(0, 0, 0)
	err := json.Unmarshal([]byte(`"2024-11-invalid"`), &d)
	assert.ErrorIs(t, err, ErrInvalidDate)
}
