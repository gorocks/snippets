package strings

import (
	"testing"
)

func TestIsDeformation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"'' 和 '' 不互为变形词", args{"", ""}, false},
		{"123 和 231 互为变形词", args{"123", "231"}, true},
		{"123 和 2331 不互为变形词", args{"123", "2331"}, false},
		{"2331 和 123 不互为变形词", args{"2331", "123"}, false},
		{"2331 和 1234 不互为变形词", args{"2331", "1234"}, false},
		{"你好啊 和 好你啊 互为变形词", args{"你好啊", "好你啊"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDeformation(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("IsDeformation() = %v, want %v", got, tt.want)
			}
		})
	}
}
