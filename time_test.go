package gouse

import (
	"testing"
	"time"
)

func TestSecond(t *testing.T) {
	if Second() != time.Now().Second() {
		t.Error("Second() != time.Now().Second()")
	}
}

func TestMinute(t *testing.T) {
	if Minute() != time.Now().Minute() {
		t.Error("Minute() != time.Now().Minute()")
	}
}

func TestHour(t *testing.T) {
	if Hour() != time.Now().Hour() {
		t.Error("Hour() != time.Now().Hour()")
	}
}

func TestDay(t *testing.T) {
	if Day() != time.Now().Day() {
		t.Error("Day() != time.Now().Day()")
	}
}

func TestMonth(t *testing.T) {
	if Month() != int(time.Now().Month()) {
		t.Error("Month() != int(time.Now().Month())")
	}
}

func TestYear(t *testing.T) {
	if Year() != time.Now().Year() {
		t.Error("Year() != time.Now().Year()")
	}
}

func TestWeekday(t *testing.T) {
	if Weekday() != int(time.Now().Weekday()) {
		t.Error("Weekday() != int(time.Now().Weekday())")
	}
}

func TestUnix(t *testing.T) {
	if Unix() != time.Now().Unix() {
		t.Error("Unix() != time.Now().Unix()")
	}
}

func TestUnixMilli(t *testing.T) {
	if UnixMilli() != time.Now().UnixNano()/int64(time.Millisecond) {
		t.Error("UnixMilli() != time.Now().UnixNano()/int64(time.Millisecond)")
	}
}

func TestUnixMicro(t *testing.T) {
	result := UnixMicro()
	currentTime := time.Now().UnixNano() / int64(time.Microsecond)

	acceptableRange := int64(1000) // 1 millisecond in microseconds

	if result < currentTime-acceptableRange || result > currentTime+acceptableRange {
		t.Errorf("UnixMicro() result %d is not within an acceptable range of current time %d", result, currentTime)
	}
}

func TestUnixNano(t *testing.T) {
	result := UnixNano()
	currentTime := time.Now().UnixNano()

	acceptableRange := int64(1000000000) // 1 second in nanoseconds

	if result < currentTime-acceptableRange || result > currentTime+acceptableRange {
		t.Errorf("UnixNano() result %d is not within an acceptable range of current time %d", result, currentTime)
	}
}

func TestUnixMilliToTime(t *testing.T) {
	if UnixMilliToTime(0) != time.Unix(0, 0) {
		t.Error("UnixMilliToTime(0) != time.Unix(0, 0)")
	}
}

func TestUnixMicroToTime(t *testing.T) {
	if UnixMicroToTime(0) != time.Unix(0, 0) {
		t.Error("UnixMicroToTime(0) != time.Unix(0, 0)")
	}
}

func TestUnixNanoToTime(t *testing.T) {
	if UnixNanoToTime(0) != time.Unix(0, 0) {
		t.Error("UnixNanoToTime(0) != time.Unix(0, 0)")
	}
}

func TestFormatTime(t *testing.T) {
	tests := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "Single digit hour, minute, and second",
			input:    time.Date(2024, 2, 4, 5, 6, 7, 0, time.UTC),
			expected: "05:06:07",
		},
		{
			name:     "Double digit hour, minute, and second",
			input:    time.Date(2024, 2, 4, 15, 16, 17, 0, time.UTC),
			expected: "15:16:17",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := _time2String(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %s, but got %s", tc.expected, result)
			}
		})
	}
}

/* Testing utilities */

func TestToSecond(t *testing.T) {
	second := 1
	expected := time.Duration(second) * time.Second
	actual := ToSecond(second)
	if expected != actual {
		t.Errorf(DESC_TEST, expected, actual)
	}
}

func TestToMinute(t *testing.T) {
	minute := 1
	expected := time.Duration(minute) * time.Minute
	actual := ToMinute(minute)
	if expected != actual {
		t.Errorf(DESC_TEST, expected, actual)
	}
}

func TestToHour(t *testing.T) {
	hour := 1
	expected := time.Duration(hour) * time.Hour
	actual := ToHour(hour)
	if expected != actual {
		t.Errorf(DESC_TEST, expected, actual)
	}
}

func TestSleepSecond(t *testing.T) {
	nowSecond := time.Now().Second()
	second := 1
	SleepSecond(second)
	if nowSecond == time.Now().Second() {
		t.Errorf("Expected %v but it got %v", nowSecond+second, time.Now().Second())
	}
}

// skip this test because it will take a long time
// func TestSleepMinute(t *testing.T) {
// 	nowMinute := time.Now().Minute()
// 	minute := 1

// 	SleepMinute(minute)
// 	if nowMinute == time.Now().Minute() {
// 		t.Errorf("Expected %v but it got %v", nowMinute+minute, time.Now().Minute())
// 	}
// }

// skip this test because it will take a long time
// func TestSleepHour(t *testing.T) {
// 	nowHour := time.Now().Hour()
// 	hour := 1
// 	SleepHour(hour)
// 	if nowHour == time.Now().Hour() {
// 		t.Errorf("Expected %v but it got %v", nowHour+hour, time.Now().Hour())
// 	}
// }
