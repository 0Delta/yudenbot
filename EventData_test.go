package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestGetEventData(t *testing.T) {
	type args struct {
		jsonBytes []byte
	}
	tests := []struct {
		name     string
		args     args
		wantData EventData
		wantErr  bool
		loadErr  bool
	}{
		{
			name: "test1",
			args: args{
				jsonBytes: []byte{},
			},
			wantData: EventData{
				ID:          5212,
				URL:         "wp.infra-workshop.tech?id=5212",
				Title:       "\u30a4\u30f3\u30d5\u30e9\u30a8\u30f3\u30b8\u30cb\u30a2\u306e\u305f\u3081\u306e\u30c0\u30a4\u30a8\u30c3\u30c8",
				Description: "",
			},
			wantErr: true,
		},
		{
			name: "test2",
		},
		// TODO: Add test cases.
	}
	// loading json-data
	for n := range tests {
		raw, err := ioutil.ReadFile("EventData_testcase/" + tests[n].name + ".json")
		if err != nil {
			t.Errorf("Can't load json : %v", tests[n].name)
			tests[n].loadErr = true
		} else {
			tests[n].loadErr = false
		}
		tests[n].args.jsonBytes = raw
	}
	// test execute
	for _, tt := range tests {
		if tt.loadErr {
			t.Logf("%v is Skipping...", tt.name)
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			gotData, err := GetEventData(tt.args.jsonBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("GetEventData() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestGetEventsFromWordpress(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name       string
		args       args
		wantEvents []EventData
		wantErr    bool
	}{
		{
			name: "test1",
			args: args{
				"wp.infra-workshop.tech",
			},
			wantEvents: []EventData{{
				ID:          5212,
				URL:         "wp.infra-workshop.tech?id=5212",
				Title:       "\u30a4\u30f3\u30d5\u30e9\u30a8\u30f3\u30b8\u30cb\u30a2\u306e\u305f\u3081\u306e\u30c0\u30a4\u30a8\u30c3\u30c8",
				Description: "",
			}},
			wantErr: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEvents, err := GetEventsFromWordpress(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEventsFromWordpress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotEvents, tt.wantEvents) {
				t.Errorf("GetEventsFromWordpress() = %v, want %v", gotEvents, tt.wantEvents)
			}
		})
	}
}
