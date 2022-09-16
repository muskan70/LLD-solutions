package score

import (
	"context"
	"intuitMc/requests"
	"testing"
)

func TestPushScoreToFile(t *testing.T) {
	type args struct {
		ctx context.Context
		req *requests.PushScoreRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{ctx: context.TODO(), req: &requests.PushScoreRequest{UserId: 4, Score: 10}},
			wantErr: false,
		},
		{
			name:    "test1",
			args:    args{ctx: context.TODO(), req: &requests.PushScoreRequest{UserId: 12, Score: 20}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PushScoreToFile(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("PushScoreToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
