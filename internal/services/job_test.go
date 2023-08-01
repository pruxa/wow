package services

import (
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestJobService_AcceptJob(t *testing.T) {
	type fields struct {
		currentJob job
	}
	type args struct {
		number int64
		hash   string
	}
	created := time.Now().Unix()
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			"Wrong hash",
			fields{
				job{
					last:    "",
					number:  0,
					created: time.Now().Unix(),
					hash:    "0x000000001",
				},
			},
			args{
				number: 0,
				hash:   "0x000000000",
			},
			false,
			true,
		},
		{
			"Wrong number",
			fields{
				job{
					last:    "0x68c54b74b7b74b69821f6755236402b32b7d184eb928cd5e694e5d934c44fb52",
					number:  10,
					created: time.Now().Unix(),
					hash:    "0x68c54b74b7b74b69821f6755236402b32b7d184eb928cd5e694e5d934c44fb34",
				},
			},
			args{
				number: 0,
				hash:   "0x68c54b74b7b74b69821f6755236402b32b7d184eb928cd5e694e5d934c44fb34",
			},
			false,
			false,
		},
		{
			"Good job",
			fields{
				job{
					last:    "last_hash",
					number:  100,
					created: created,
					hash:    getHexedHash("last_hash" + strconv.FormatInt(created, 10) + strconv.FormatInt(100, 10)),
				},
			},
			args{
				number: 100,
				hash:   getHexedHash("last_hash" + strconv.FormatInt(created, 10) + strconv.FormatInt(100, 10)),
			},
			true,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js := &JobService{
				currentJob: tt.fields.currentJob,
			}
			got, err := js.AcceptJob(tt.args.number, tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("AcceptJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AcceptJob() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJobService_GetCurrentJob(t *testing.T) {
	type fields struct {
		currentJob job
	}
	tests := []struct {
		name   string
		fields fields
		want   Job
	}{
		{
			"Get current job",
			fields{
				job{
					last:       "0x68c54b74b7b74b69821f6755236402b32b7d184eb928cd5e694e5d934c44fb34",
					number:     100,
					created:    1690884797,
					hash:       "0x68c54b74b7b74b69821f6755236402b32b7d184eb928cd5e694e5d934c44fb54",
					difficulty: 1000000,
				},
			},
			Job{
				Created:    1690884797,
				Last:       "0x68c54b74b7b74b69821f6755236402b32b7d184eb928cd5e694e5d934c44fb34",
				Difficulty: 1000000,
				Hash:       "0x68c54b74b7b74b69821f6755236402b32b7d184eb928cd5e694e5d934c44fb54",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			js := &JobService{
				currentJob: tt.fields.currentJob,
			}
			if got := js.GetCurrentJob(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrentJob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_h(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Hash",
			args{
				message: "0x6714aeeff2c6229c734677e0e41968989d9cc9a3d93016d1450dddef2dbd5c05123131212110",
			},
			"0x68c54b74b7b74b69821f6755236402b32b7d184eb928cd5e694e5d934c44fb34",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHexedHash(tt.args.message); got != tt.want {
				t.Errorf("getHexedHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
