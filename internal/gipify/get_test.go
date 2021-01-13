package gipify

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {

	type httpResponse struct {
		contentType string
		body        string
		status      int
	}
	type want struct {
		IP   string
		Type int
	}
	tests := []struct {
		name    string
		want    want
		http    httpResponse
		wantErr bool
	}{
		{
			name: "GetIP IPv4",
			http: httpResponse{contentType: `application/json`, body: `{"ip":"98.207.254.136"}`, status: 200},
			want: want{IP: "98.207.254.136", Type: IPV4},
		},
		{
			name: "GetIP IPv6",
			http: httpResponse{contentType: `application/json`, body: `{"ip":"2a00:1450:400f:80d::200e"}`, status: 200},
			want: want{IP: "2a00:1450:400f:80d::200e", Type: IPV6},
		},
		{
			name:    "400 error",
			http:    httpResponse{contentType: `application/json`, body: `{"ip":"2a00:1450:400f:80d::200e"}`, status: 400},
			want:    want{IP: "2a00:1450:400f:80d::200e", Type: IPV6},
			wantErr: true,
		},
		{
			name:    "500 error",
			http:    httpResponse{contentType: `application/json`, body: `{"ip":"2a00:1450:400f:80d::200e"}`, status: 500},
			want:    want{IP: "2a00:1450:400f:80d::200e", Type: IPV6},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.New()
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.http.status)
				w.Header().Set("Content-Type", tt.http.contentType)
				fmt.Fprintln(w, tt.http.body)
			}))
			defer ts.Close()

			ipURL = ts.URL
			ipWant := IP{IP: tt.want.IP, Type: tt.want.Type}
			got, err := GetIP()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, &ipWant) {
				t.Errorf("GetIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		i io.Reader
	}

	type want struct {
		IP   string
		Type int
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name:    "Parse json ip 192.168.0.1",
			args:    args{i: strings.NewReader(`{"ip":"192.168.0.1"}`)},
			want:    want{IP: "192.168.0.1", Type: IPV4},
			wantErr: false,
		},
		{
			name:    "Parse json ip 127.0.0.1",
			args:    args{i: strings.NewReader(`{"ip":"127.0.0.1"}`)},
			want:    want{IP: "127.0.0.1", Type: IPV4},
			wantErr: false,
		},
		{
			name:    "Parse json ip ::1",
			args:    args{i: strings.NewReader(`{"ip":"::1"}`)},
			want:    want{IP: "::1", Type: IPV6},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			ipWant := IP{IP: tt.want.IP, Type: tt.want.Type}

			if !reflect.DeepEqual(got, &ipWant) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ipType(t *testing.T) {
	type args struct {
		i string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "IPv4: 192.168.0.1",
			args: args{"192.168.0.1"},
			want: IPV4,
		},
		{
			name: "IPv4: 127.0.0.1",
			args: args{"192.168.0.1"},
			want: IPV4,
		},
		{
			name: "IPv6: ::1",
			args: args{"::1"},
			want: IPV6,
		},
		{
			name: "IPv6: 2001:4860:4860::8888",
			args: args{"2001:4860:4860::8888"},
			want: IPV6,
		},
		{
			name: "Error: 127.0.0.",
			args: args{"127.0.0."},
			want: UNKNOWN,
		},
		{
			name: "Error: Empty",
			args: args{""},
			want: UNKNOWN,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ipType(tt.args.i); got != tt.want {
				t.Errorf("ipType() = %v, want %v", got, tt.want)
			}
		})
	}
}
