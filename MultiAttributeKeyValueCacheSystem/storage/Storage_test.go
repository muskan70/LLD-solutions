package storage

import (
	"reflect"
	"testing"
)

func TestStorage_Add(t *testing.T) {
	type fields struct {
		KeyStore       map[string][]AttributeValue
		AttributeType  map[string]reflect.Type
		AttributeStore map[string]map[string]interface{}
	}
	stg := &Storage{KeyStore: make(map[string][]AttributeValue),
		AttributeType:  make(map[string]reflect.Type),
		AttributeStore: make(map[string]map[string]interface{}),
	}
	type args struct {
		key   string
		value map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{name: "test1",
			args: args{key: "delhi", value: map[string]interface{}{
				"pollution-level": "moderate",
				"latitude":        12.94,
				"longitude":       77.64,
				"free_food":       "true",
			},
			},
			want: nil,
		},
		{
			name: "test2",
			args: args{key: "kolkata",
				value: map[string]interface{}{
					"pollution-level": "moderate",
					"latitude":        "vip",
					"longitude":       77.64,
					"free_food":       "true",
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stg.Add(tt.args.key, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
