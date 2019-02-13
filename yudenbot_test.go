// yuden

package main

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestYudenbot(t *testing.T) {
	type args struct {
		ctx      context.Context
		execList []Executor
	}
	ctx, cancel := context.WithCancel(context.TODO())
	tests := []struct {
		name string
		args args
	}{
		{
			name: "bot test",
			args: args{
				ctx: ctx,
				execList: []Executor{
					Executor{
						Name: "fizz",
						Fnc:  func(ctx context.Context) (err error) { log.Println("fizz"); return nil },
						Tick: 3 * time.Second,
					},
					Executor{
						Name: "buzz",
						Fnc:  func(ctx context.Context) (err error) { log.Println("buzz"); return nil },
						Tick: 5 * time.Second,
					},
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go Yudenbot(tt.args.ctx, tt.args.execList)
			time.Sleep(25 * time.Second)
			cancel()
		})
	}
}
