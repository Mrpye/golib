package path

import "testing"

func TestStripTrailingSlashes(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty slice",
			args: args{
				path: "c:\\test\\",
			},
			want: "c:\\test",
		},
		{
			name: "empty slice",
			args: args{
				path: "./test/",
			},
			want: "./test",
		},
		{
			name: "empty slice",
			args: args{
				path: "./test",
			},
			want: "./test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StripTrailingSlashes(tt.args.path); got != tt.want {
				t.Errorf("StripTrailingSlashes() = %v, want %v", got, tt.want)
			}
		})
	}
}
