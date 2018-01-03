package defers

import "testing"

func Test_doubleInt(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{0}, 0},
		{"t2", args{40}, 80},
		{"t3", args{50}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doubleInt(tt.args.i); got != tt.want {
				t.Errorf("doubleInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doubleInt0(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name  string
		args  args
		wantR int
	}{
		{"t1", args{0}, 0},
		{"t2", args{40}, 80},
		{"t3", args{50}, 50},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := doubleInt0(tt.args.i); gotR != tt.wantR {
				t.Errorf("doubleInt0() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
