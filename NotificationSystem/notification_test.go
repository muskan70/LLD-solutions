package main

import (
	"testing"
	"time"
)

func TestNotification_UpdateScheduledTime(t *testing.T) {
	type fields struct {
		Id            uint64
		UserId        int
		Channel       int
		Message       string
		ScheduledTime time.Time
		Priority      int
		Status        int
	}
	type args struct {
		newTime time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			fields:  fields{Id: 1, UserId: 1, Channel: 1, Message: "hello I am muskan", ScheduledTime: time.Now().Add(time.Second * 10), Priority: 3, Status: 2},
			args:    args{newTime: time.Now().Add(time.Second * 20)},
			wantErr: false,
		},
		{
			name:    "test2",
			fields:  fields{Id: 1, UserId: 1, Channel: 1, Message: "hello I am muskan", ScheduledTime: time.Now().Add(time.Second * 10), Priority: 3, Status: 2},
			args:    args{newTime: time.Now()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Notification{
				Id:            tt.fields.Id,
				UserId:        tt.fields.UserId,
				Channel:       tt.fields.Channel,
				Message:       tt.fields.Message,
				ScheduledTime: tt.fields.ScheduledTime,
				Priority:      tt.fields.Priority,
				Status:        tt.fields.Status,
			}
			if err := n.UpdateScheduledTime(tt.args.newTime); (err != nil) != tt.wantErr {
				t.Errorf("Notification.UpdateScheduledTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotification_CancelNotification(t *testing.T) {
	type fields struct {
		Id            uint64
		UserId        int
		Channel       int
		Message       string
		ScheduledTime time.Time
		Priority      int
		Status        int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "test1",
			fields:  fields{Id: 1, UserId: 1, Channel: 1, Message: "hello I am muskan", ScheduledTime: time.Now().Add(time.Second * 10), Priority: 3, Status: 2},
			wantErr: false,
		},
		{
			name:    "test2",
			fields:  fields{Id: 1, UserId: 1, Channel: 1, Message: "hello I am muskan", ScheduledTime: time.Now().Add(-time.Second * 10), Priority: 3, Status: 1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Notification{
				Id:            tt.fields.Id,
				UserId:        tt.fields.UserId,
				Channel:       tt.fields.Channel,
				Message:       tt.fields.Message,
				ScheduledTime: tt.fields.ScheduledTime,
				Priority:      tt.fields.Priority,
				Status:        tt.fields.Status,
			}
			if err := n.CancelNotification(); (err != nil) != tt.wantErr {
				t.Errorf("Notification.CancelNotification() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotification_UpdateMessage(t *testing.T) {
	type fields struct {
		Id            uint64
		UserId        int
		Channel       int
		Message       string
		ScheduledTime time.Time
		Priority      int
		Status        int
	}
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "test1",
			fields:  fields{Id: 1, UserId: 1, Channel: 1, Message: "hello I am muskan", ScheduledTime: time.Now().Add(time.Second * 10), Priority: 3, Status: 2},
			args:    args{message: "its okay"},
			wantErr: false,
		},
		{
			name:    "test2",
			fields:  fields{Id: 1, UserId: 1, Channel: 1, Message: "hello I am muskan", ScheduledTime: time.Now().Add(time.Second * 10), Priority: 3, Status: 2},
			args:    args{message: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Notification{
				Id:            tt.fields.Id,
				UserId:        tt.fields.UserId,
				Channel:       tt.fields.Channel,
				Message:       tt.fields.Message,
				ScheduledTime: tt.fields.ScheduledTime,
				Priority:      tt.fields.Priority,
				Status:        tt.fields.Status,
			}
			if err := n.UpdateMessage(tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Notification.UpdateMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
