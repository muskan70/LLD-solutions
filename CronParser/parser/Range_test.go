package parser

import (
	"cron/segment"
	"errors"
	"reflect"
	"testing"
)

func TestRangeParser_GetPossibilities(t *testing.T) {
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
			fields:  fields{exp: segment.NewDay("1-5").BaseSegment},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: nil,
		},
		{
			name:    "test7",
			fields:  fields{exp: segment.NewWeekday("mon-fri").BaseSegment},
			want:    []int{1, 2, 3, 4, 5},
			wantErr: nil,
		},
		{
			name:    "test2",
			fields:  fields{exp: segment.NewMonth("1-").BaseSegment},
			want:    nil,
			wantErr: errors.New("invalid segment expression"),
		},
		{
			name:    "test3",
			fields:  fields{exp: segment.NewWeekday("1-8").BaseSegment},
			want:    nil,
			wantErr: errors.New("range maximum is not valid"),
		},
		{
			name:    "test4",
			fields:  fields{exp: segment.NewMonth("0-6").BaseSegment},
			want:    nil,
			wantErr: errors.New("range minimum is not valid"),
		},
		{
			name:    "test5",
			fields:  fields{exp: segment.NewMonth("13-16").BaseSegment},
			want:    nil,
			wantErr: errors.New("range minimum is not valid"),
		},
		{
			name:    "test6",
			fields:  fields{exp: segment.NewHour("50-20").BaseSegment},
			want:    nil,
			wantErr: errors.New("range minimum/maximum are in wrong order"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &RangeParser{
				exp: tt.fields.exp,
			}
			got, err := a.GetPossibilities()
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("RangeParser.GetPossibilities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RangeParser.GetPossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}
