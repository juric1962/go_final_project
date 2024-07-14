package nextdate

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/juric1962/go_final_project/tasks"
	_ "github.com/mattn/go-sqlite3"
)

func NextDate(now time.Time, date string, repeat string) (string, error) {
	dayInMonth := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	err := errors.New("неправильный формат")
	if len(date) != 8 {
		return "", err
	}
	if len(repeat) == 0 {
		return "", err
	}
	start, err1 := time.Parse(tasks.TimeFormat, date)
	if err1 != nil {
		return "", err
	}
	s := strings.SplitAfterN(date, "", 8)
	year, err0 := strconv.Atoi(s[0] + s[1] + s[2] + s[3])
	month, err1 := strconv.Atoi(s[4] + s[5])
	day, err2 := strconv.Atoi(s[6] + s[7])
	if err1 != nil || err2 != nil || err0 != nil {
		return "", err
	}
	if year%4 == 0 {
		dayInMonth[2] = 29
	}
	if month > 12 || month <= 0 {
		return "", err
	}
	if dayInMonth[month] < day {
		return "", err
	}
	if repeat[0] == 'y' {
		for {
			next := start.AddDate(1, 0, 0)
			if next.After(now) {
				date := next.Format(tasks.TimeFormat)
				return date, nil
			}
			start = next
		}
	}
	if repeat[0] == 'd' {
		va := strings.Split(repeat, " ")
		if len(va) != 2 {
			return "", err
		}
		dayS, err1 := strconv.Atoi(va[1])
		if dayS <= 0 || dayS > 400 || err1 != nil {
			return "", err
		}
		start, _ := time.Parse(tasks.TimeFormat, date)
		for {
			next := start.AddDate(0, 0, dayS)
			if next.After(now) {
				date := next.Format(tasks.TimeFormat)
				return date, nil
			}
			start = next
		}
	}
	return "", err
}
