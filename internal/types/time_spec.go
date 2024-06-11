package types

import (
	"fmt"
	"time"
	"database/sql/driver"
)

const (
	FORMAT_TIME = "2006-01-02 15:04:05"
)

type WrapTime struct {
	time.Time
}

func (w WrapTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", w.Format(FORMAT_TIME))
	return []byte(formatted), nil
}

func (w *WrapTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+FORMAT_TIME+`"`, string(data), time.Local)
	w.Time = now

	return
}

func (w WrapTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	
	if w.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}

	return w.Time, nil
}

func (w *WrapTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		w.Time = value
		return nil
	}

	return fmt.Errorf("can not convert %v to timestamp", v)
}