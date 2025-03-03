package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestCronStringParser(t *testing.T) {
	type args struct {
		cronStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name:    "test1",
			args:    args{cronStr: "*/15 0 1,15 * 1-5 /usr/bin/find"},
			wantErr: nil,
		},
		{
			name:    "test2",
			args:    args{cronStr: "*/105 0 1,5 * 1-5 /usr/bin/find"},
			wantErr: errors.New("step size is more than maximum value"),
		},
		{
			name:    "test3",
			args:    args{cronStr: "*/105 0 1,5 /usr/bin/find"},
			wantErr: errors.New("invalid cron string : unable to break into following format (minute, hour, day of month, month, and day of week) plus a command"),
		},
		{
			name:    "test4",
			args:    args{cronStr: "*/15 0 1,15 * 1-5 /usr/bin/find 1 2 3"},
			wantErr: nil,
		},
		{
			name:    "test5",
			args:    args{cronStr: "*/15 0 1,15 * mon-fri /usr/bin/find 1 2 3"},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CronStringParser(tt.args.cronStr); !reflect.DeepEqual(got, tt.wantErr) {
				t.Errorf("CronStringParser() error = %v, wantErr %v", got, tt.wantErr)
			}
		})
	}
}
