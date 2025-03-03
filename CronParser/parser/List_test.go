package parser

import (
	"cron/segment"
	"errors"
	"reflect"
	"testing"
)

func TestListParser_GetPossibilities(t *testing.T) {
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
			fields:  fields{exp: segment.NewDay("1,5").BaseSegment},
			want:    []int{1, 5},
			wantErr: nil,
		},
		{
			name:    "test2",
			fields:  fields{exp: segment.NewMonth("a,15").BaseSegment},
			want:    nil,
			wantErr: errors.New("invalid segment expression"),
		},
		{
			name:    "test3",
			fields:  fields{exp: segment.NewWeekday("1,8").BaseSegment},
			want:    nil,
			wantErr: errors.New("the value for segment is more than the maximum allowed"),
		},
		{
			name:    "test4",
			fields:  fields{exp: segment.NewMonth("1,").BaseSegment},
			want:    nil,
			wantErr: errors.New("invalid segment expression"),
		},
		{
			name:    "test5",
			fields:  fields{exp: segment.NewMonth("1-3,5-6").BaseSegment},
			want:    []int{1, 2, 3, 5, 6},
			wantErr: nil,
		},
		{
			name:    "test6",
			fields:  fields{exp: segment.NewHour("1-4,3-6").BaseSegment},
			want:    []int{1, 2, 3, 4, 5, 6},
			wantErr: nil,
		},
		{
			name:    "test7",
			fields:  fields{exp: segment.NewWeekday("wed,thr-sun").BaseSegment},
			want:    []int{3, 4, 5, 6, 7},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ListParser{
				exp: tt.fields.exp,
			}
			got, err := a.GetPossibilities()
			if !reflect.DeepEqual(err, tt.wantErr) {
				t.Errorf("ListParser.GetPossibilities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListParser.GetPossibilities() = %v, want %v", got, tt.want)
			}
		})
	}
}
