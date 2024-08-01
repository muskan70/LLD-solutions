package main

import (
	"fmt"
	"testing"
)

func Test_getKeywordOccurences(t *testing.T) {
	type args struct {
		urls        []string
		keywords    []string
		noofWorkers int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				urls:        []string{"https://www.paytm.com", "https://www.google.com"},
				keywords:    []string{"paytm", "google"},
				noofWorkers: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response := getKeywordOccurences(tt.args.urls, tt.args.keywords, tt.args.noofWorkers)
			fmt.Println(response)
		})
	}
}
