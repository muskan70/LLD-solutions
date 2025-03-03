package parser

import (
	"cron/segment"
	"errors"
	"reflect"
	"testing"
)

func TestStepParser_GetPossibilities(t *testing.T) {
	type fields struct {
		exp segment.BaseSegment
	}
	tests := []struct {
		name    string
		fields  fields
		want    []int
		wantErr error
	}{
		{
			name:    "test1",
			fields:  fields{exp: segment.NewMinute("*/15").BaseSegment},
			want:    []int{0, 15, 30, 45},
			wantErr: nil,
		},
		{
			name:    "test2",
			fields:  fields{exp: segment.NewMonth("*/").BaseSegment},
			want:    nil,
			wantErr: errors.New("step does not have valid expression"),
		},
		{
			name:    "test3",
			fields:  fields{exp: segment.NewDay("a/15").BaseSegment},
			want:    nil,
			wantErr: errors.New("step does not have valid expression"),
		},
		{
			name:    "test4",
			fields:  fields{exp: segment.NewHour("*/60").BaseSegment},
			want:    nil,
			wantErr: errors.New("step size is more than maximum value"),
		},
		{
			name:    "test5",
			fields:  fields{exp: segment.NewHour("60/15").BaseSegment},
			want:    nil,
			wantErr: errors.New("step start is more than maximum value"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &StepParser{
				exp: tt.fields.exp,
			}
			got, err := a.GetPossibilities()
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("StepParser.GetPossibilities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StepParser.GetPossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}
