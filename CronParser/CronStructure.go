package main

import (
	"cron/parser"
	"cron/segment"
	"errors"
	"fmt"
	"strings"
)

type CronString struct {
	Minute           segment.Minute
	Hour             segment.Hour
	DayOfMonth       segment.Day
	Month            segment.Month
	WeekDay          segment.Weekday
	Command          string
	CommandArguments []string
}

func NewCronString(cronStr string) (*CronString, error) {
	cronParts := strings.Split(cronStr, " ")
	if err := ValidateCronString(cronParts); err != nil {
		return nil, err
	}
	return &CronString{
		Minute:           segment.NewMinute(cronParts[0]),
		Hour:             segment.NewHour(cronParts[1]),
		DayOfMonth:       segment.NewDay(cronParts[2]),
		Month:            segment.NewMonth(cronParts[3]),
		WeekDay:          segment.NewWeekday(cronParts[4]),
		Command:          cronParts[5],
		CommandArguments: cronParts[6:],
	}, nil
}

func ValidateCronString(cronStr []string) error {
	if len(cronStr) < 6 {
		return errors.New("invalid cron string : unable to break into following format (minute, hour, day of month, month, and day of week) plus a command")
	}
	return nil
}

func CronStringParser(cronStr string) error {
	c, err := NewCronString(cronStr)
	if err != nil {
		return err
	}

	if minuteParser, err := parser.GetParser(c.Minute.BaseSegment); err != nil {
		return err
	} else if c.Minute.PossibleValues, err = minuteParser.GetPossibilities(); err != nil {
		return err
	}
	if hourParser, err := parser.GetParser(c.Hour.BaseSegment); err != nil {
		return err
	} else if c.Hour.PossibleValues, err = hourParser.GetPossibilities(); err != nil {
		return err
	}
	if dayParser, err := parser.GetParser(c.DayOfMonth.BaseSegment); err != nil {
		return err
	} else if c.DayOfMonth.PossibleValues, err = dayParser.GetPossibilities(); err != nil {
		return err

	}
	if monthParser, err := parser.GetParser(c.Month.BaseSegment); err != nil {
		return err
	} else if c.Month.PossibleValues, err = monthParser.GetPossibilities(); err != nil {
		return err
	}
	if weekdayParser, err := parser.GetParser(c.WeekDay.BaseSegment); err != nil {
		return err
	} else if c.WeekDay.PossibleValues, err = weekdayParser.GetPossibilities(); err != nil {
		return err
	}

	fmt.Println("Parsed cron string Output:")
	fmt.Println("minute        ", c.Minute.GetPossibleValuesString())
	fmt.Println("hour          ", c.Hour.GetPossibleValuesString())
	fmt.Println("day of month  ", c.DayOfMonth.GetPossibleValuesString())
	fmt.Println("month         ", c.Month.GetPossibleValuesString())
	fmt.Println("day of week   ", c.WeekDay.GetPossibleValuesString())
	fmt.Println("command       ", c.Command, strings.Join(c.CommandArguments, " "))
	return nil
}
