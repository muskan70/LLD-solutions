package parser

import (
	"cron/segment"
	"reflect"
	"testing"
)

func TestAllParser_GetPossibilities(t *testing.T) {
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
			name:    "test2",
			fields:  fields{exp: segment.NewDay("*").BaseSegment},
			want:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
			wantErr: nil,
		},
		{
			name:    "test1",
			fields:  fields{exp: segment.NewWeekday("*").BaseSegment},
			want:    []int{1, 2, 3, 4, 5, 6, 7},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AllParser{
				exp: tt.fields.exp,
			}
			got, err := a.GetPossibilities()
			if err != tt.wantErr {
				t.Errorf("AllParser.GetPossibilities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllParser.GetPossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}
