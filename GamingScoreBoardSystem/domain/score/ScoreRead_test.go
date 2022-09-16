package score

import (
	"context"
	"testing"
)

func TestReadScoreFromFile(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{ctx: context.TODO()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadScoreFromFile(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("ReadScoreFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
