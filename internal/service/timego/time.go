package timego

import (
	"fmt"
	"time"
)

func NewTimeEntity(date string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"

	dateEntity, err := time.Parse(layout, fmt.Sprintf("%s 00:00:00", date))
	if err != nil {
		return time.Time{}, err
	}

	return dateEntity, nil
}

func IsActive(date string) (bool, error) {
	t, err := NewTimeEntity(date)
	if err != nil {
		return false, err
	}

	todayEntity, err := NewTimeEntity(time.Now().Format("2006-01-02"))
	if err != nil {
		return false, err
	}

	if t.Before(todayEntity) || t.Equal(todayEntity) {
		return true, nil
	}

	return false, nil
}

func IsWeekend(date string) (bool, error) {
	t, err := NewTimeEntity(date)
	if err != nil {
		return false, err
	}

	weekday := int(t.Weekday())

	if weekday == 6 || weekday == 7 {
		return true, nil
	}

	return false, nil
}
