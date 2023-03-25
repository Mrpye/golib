// This package contains helper functions for that are related to converting interfaces to other types.

package convert

import (
	"testing"
)

func TestIsString(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "TestIsString",
			args: args{
				value: "Hello",
			},
			want: true,
		},
		{
			name: "TestIsString",
			args: args{
				value: 123,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsString(tt.args.value); got != tt.want {
				t.Errorf("IsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBool(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "TestIsBool",
			args: args{
				value: true,
			},
			want: true,
		},
		{
			name: "TestIsBool",
			args: args{
				value: 123,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBool(tt.args.value); got != tt.want {
				t.Errorf("IsBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInt(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "TestIsInt",
			args: args{
				value: 123,
			},
			want: true,
		},
		{
			name: "TestIsInt",
			args: args{
				value: "Hello",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInt(tt.args.value); got != tt.want {
				t.Errorf("IsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFloat(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "TestIsFloat",
			args: args{
				value: 123.45,
			},
			want: true,
		},
		{
			name: "TestIsFloat",
			args: args{
				value: "Hello",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFloat(tt.args.value); got != tt.want {
				t.Errorf("IsFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMapInterface(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "TestIsMapInterface",
			args: args{
				value: map[string]interface{}{
					"Hello": "World",
				},
			},
			want: true,
		},
		{
			name: "TestIsMapInterface",
			args: args{
				value: "Hello",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMapInterface(tt.args.value); got != tt.want {
				t.Errorf("IsMapInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}
