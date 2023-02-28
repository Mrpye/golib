package math

import (
	"testing"
)

func TestAND(t *testing.T) {
	type args struct {
		a bool
		b bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty slice",
			args: args{
				a: true,
				b: true,
			},
			want: true,
		},
		{
			name: "empty slice",
			args: args{
				a: false,
				b: true,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: false,
				b: false,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: false,
				b: true,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AND(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("AND() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOR(t *testing.T) {
	type args struct {
		a bool
		b bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty slice",
			args: args{
				a: false,
				b: true,
			},
			want: true,
		},
		{
			name: "empty slice",
			args: args{
				a: true,
				b: false,
			},
			want: true,
		},
		{
			name: "empty slice",
			args: args{
				a: false,
				b: false,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: true,
				b: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OR(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("OR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNOT(t *testing.T) {
	type args struct {
		a bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty slice",
			args: args{
				a: true,
			},
			want: false,
		},
		{
			name: "empty slice",
			args: args{
				a: false,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NOT(tt.args.a); got != tt.want {
				t.Errorf("NOT() = %v, want %v", got, tt.want)
			}
		})
	}
}
