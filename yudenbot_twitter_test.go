package yudenbot

import (
	"io/ioutil"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func getTestToken(t *testing.T) *TwitterAuth {
	buf, err := ioutil.ReadFile("./.test.token.yml")
	if err != nil {
		t.Fatal("Error while load token : ", err)
	}
	var auth TwitterAuth
	err = yaml.Unmarshal(buf, &auth)
	if err != nil {
		t.Fatal("Error while unmarshal token: ", err)
	}
	return &auth
}
func Test_tweet(t *testing.T) {
	type args struct {
		message string
		auth    *TwitterAuth
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "tweet test",
			args: args{
				message: "Test from golang code.",
				auth:    getTestToken(t),
			},
			wantErr: false,
		},
		{
			name: "tweet fail test",
			args: args{
				message: "Test from golang code. But It's ERROR !",
				auth: &TwitterAuth{
					AccessSecret:   "",
					AccessToken:    "",
					ConsumerKey:    "",
					ConsumerSecret: "",
				},
			},
			wantErr: true,
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tweet(tt.args.message, tt.args.auth); (err != nil) != tt.wantErr {
				t.Errorf("tweet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	t.Log("PLEASE. Check tweet posted")
}

func Test_getTwitterAPI(t *testing.T) {
	type args struct {
		auth *TwitterAuth
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
	}{
		{
			name:    "test1",
			args:    args{getTestToken(t)},
			wantNil: false,
		},
		{
			name:    "test2 cache test",
			args:    args{getTestToken(t)},
			wantNil: false,
		},
		{
			name: "test3 reAuth test",
			args: args{
				&TwitterAuth{
					AccessSecret:   "",
					AccessToken:    "",
					ConsumerKey:    "",
					ConsumerSecret: "",
				},
			},
			wantNil: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantNil {
				if got := getTwitterAPI(tt.args.auth); got != nil {
					t.Errorf("getTwitterAPI() is return, want nil. : %v", got)
				}
			} else {
				if got := getTwitterAPI(tt.args.auth); got == nil {
					t.Errorf("getTwitterAPI() is return nil, want something")
				}
			}
		})
	}
}
