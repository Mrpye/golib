package encrypt

import (
	"testing"
)

func TestBase64EncString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: args{
				value: "",
			},
			want: "",
		},
		{
			name: "name:password",
			args: args{
				value: "name:password",
			},
			want: "bmFtZTpwYXNzd29yZA==",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Base64EncString(tt.args.value); got != tt.want {
				t.Errorf("Base64EncString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBase64DecString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "name:password",
			args: args{
				value: "bmFtZTpwYXNzd29yZA==",
			},
			want: "name:password",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Base64DecString(tt.args.value); got != tt.want {
				t.Errorf("Base64DecString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncryptDecryptString(t *testing.T) {

	value := "name:password"
	pass_word := "password"

	got, err := EncryptString(value, pass_word)
	if err != nil {
		t.Errorf("EncryptString() error = %v", err)
		return
	}
	got, err = DecryptString(got, pass_word)
	if err != nil {
		t.Errorf("EncryptString() error = %v", err)
		return
	}
	if got != value {
		t.Errorf("DecryptString() = %v, want %v", got, value)
	}
}
