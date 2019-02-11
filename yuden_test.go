// yuden

package yudenbot

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestYudenbot(t *testing.T) {
	type args struct {
		execList []Executor
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "bot test",
			args: args{
				execList: []Executor{
					Executor{
						name: "fizz",
						fnc:  func(ctx context.Context) (err error) { log.Println("fizz"); return nil },
						tick: 3 * time.Second,
					},
					Executor{
						name: "buzz",
						fnc:  func(ctx context.Context) (err error) { log.Println("buzz"); return nil },
						tick: 5 * time.Second,
					},
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Yudenbot(tt.args.execList)
		})
	}
}
