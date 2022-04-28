package day4

import (
	"reflect"
	"testing"
)

/*
get  [[{22 false} {13 false} {17 false} {11 false} {0 false} {0 false}] [{0 false} {8 false} {0 false} {2 false} {23 false} {0 false} {4 false} {24 false}] [{21 false} {0 false} {9 false} {14 false} {16 false} {0 false} {7 false}] [{0 false} {6 false} {10 false} {0 false} {3 false} {18 false} {0 false} {5 false}] [{0 false} {1 false} {12 false} {20 false} {15 false} {19 false}]],
want [[{22 false} {13 false} {17 false} {11 false}] [{8 false} {2 false} {23 false} {4 false} {24 false}] [{21 false} {9 false} {14 false} {16 false} {7 false}] [{6 false} {10 false} {3 false} {18 false} {5 false}] [{1 false} {12 false} {20 false} {15 false} {19 false}]]

*/

func TestTextToBoard(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    Board
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Simple board",
			args: args{
				s: "22 13 17 11  0\n 8  2 23  4 24\n21  9 14 16  7\n 6 10  3 18  5\n 1 12 20 15 19",
			},
			want: Board{
				{Point{i: 22}, Point{i: 13}, Point{i: 17}, Point{i: 11}, Point{i: 0}},
				{Point{i: 8}, Point{i: 2}, Point{i: 23}, Point{i: 4}, Point{i: 24}},
				{Point{i: 21}, Point{i: 9}, Point{i: 14}, Point{i: 16}, Point{i: 7}},
				{Point{i: 6}, Point{i: 10}, Point{i: 3}, Point{i: 18}, Point{i: 5}},
				{Point{i: 1}, Point{i: 12}, Point{i: 20}, Point{i: 15}, Point{i: 19}},
			},
			wantErr: false,
		},
		{
			name: "Simple board, with error",
			args: args{
				s: "a 13 17 11  0\n 8  2 23  4 24\n21  9 14 16  7\n 6 10  3 18  5\n 1 12 20 15 19",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TextToBoard(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("TextToBoard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TextToBoard() got = %v, want %v", got, tt.want)
			}
		})
	}
}
