package parser

import (
	"cron/segment"
	"errors"
	"reflect"
	"testing"
)

func TestExactParser_GetPossibilities(t *testing.T) {
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
			fields:  fields{exp: segment.NewDay("1").BaseSegment},
			want:    []int{1},
			wantErr: nil,
		},
		{
			name:    "test2",
			fields:  fields{exp: segment.NewMonth("59").BaseSegment},
			want:    nil,
			wantErr: errors.New("the value for segment is more than the maximum allowed"),
		},
		{
			name:    "test3",
			fields:  fields{exp: segment.NewWeekday("0").BaseSegment},
			want:    nil,
			wantErr: errors.New("the value for segment is less than the minimum allowed"),
		},
		{
			name:    "test4",
			fields:  fields{exp: segment.NewWeekday("mon").BaseSegment},
			want:    []int{1},
			wantErr: nil,
		},
		{
			name:    "test3",
			fields:  fields{exp: segment.NewWeekday("tuds").BaseSegment},
			want:    nil,
			wantErr: errors.New("invalid segment expression"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ExactParser{
				exp: tt.fields.exp,
			}
			got, err := a.GetPossibilities()
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("ExactParser.GetPossibilities() error = %v, wantErr %v, got %v", err, tt.wantErr, got)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExactParser.GetPossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}
