package game

import (
	"testing"
)

func TestIncludes(t *testing.T) {
	type args struct {
		str string
		target []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"string a contains", args{"a", []string{"a", "b", "c"}}, true},
		{"string z is not contains", args{"z", []string{"a", "b", "c"}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := contains(tt.args.str, tt.args.target)
			if tt.want {
				if !got {
					t.Errorf("contains() = %v, want %v", got, tt.want)
				}
			} else {
				if got {
					t.Errorf("contains() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestFuga(t *testing.T) {
	type args struct {
		inputs   []string
		corrects []string
	}
	type want struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"One match, one included",
			args{[]string{"1", "3", "0"}, []string{"1", "2", "3"}},
			want{1, 1},
		},
		{
			"Zero match, three included",
			args{[]string{"2", "3", "1"}, []string{"1", "2", "3"}},
			want{0, 3},
		},
		{
			"Zero match, zero included",
			args{[]string{"9", "8", "0"}, []string{"1", "2", "3"}},
			want{0, 0},
		},
		{
			"Three match, zero included",
			args{[]string{"1", "2", "3"}, []string{"1", "2", "3"}},
			want{3, 0},
		},
	}

	g := New()
	g.Digit = 3
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g.PlayerNums = tt.args.inputs
			g.CorrectNums = tt.args.corrects
			gotA, gotB := g.hoge()
			if !(gotA == tt.want.a && gotB == tt.want.b) {
				t.Errorf("g.hoge() = %v, %v, want %v, %v", gotA, gotB, tt.want.a, tt.want.b)
			}
		})
	}
}
