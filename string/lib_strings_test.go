package string

import (
	"strings"
	"testing"
)

// Write a test code for all functions in lib\lib_strings.go
func TestGzipBase64String(t *testing.T) {
	input := "Hello World"
	expected := "H4sIAAAAAAAA//JIzcnJVwjPL8pJAQQAAP//VrEXSgsAAAA="
	actual, _ := GzipBase64String(input)
	if actual != expected {
		t.Errorf("GzipBase64String(%s) = %s; want %s", input, actual, expected)
	}
}

func TestRemoveSpacesForHyphens(t *testing.T) {
	// Write test code here
	t.Run("RemoveSpacesForHyphens", func(t *testing.T) {
		input := "Hello World"
		expected := "hello-world"
		actual, _ := RemoveSpacesForHyphens(input)
		if actual != expected {
			t.Errorf("RemoveSpacesForHyphens(%s) = %s; want %s", input, actual, expected)
		}
	})
}

func TestRemoveSpacesForUnderscores(t *testing.T) {
	input := "Hello World"
	expected := "hello_world"
	actual := RemoveSpacesForUnderscores(input)
	if actual != expected {
		t.Errorf("RemoveSpacesForUnderscores(%s) = %s; want %s", input, actual, expected)
	}
}

func TestConcat(t *testing.T) {
	input := []string{"Hello", "World"}
	expected := "Hello-World"
	actual := Concat("-", "Hello", "World")
	if actual != expected {
		t.Errorf("Concat(%s) = %s; want %s", input, actual, expected)
	}
}

func TestIsNumber(t *testing.T) {
	//Should be true
	input := "123"
	expected := true
	actual := IsNumber(input)
	if actual != expected {
		t.Errorf("IsNumber(%s) = %t; want %t", input, actual, expected)
	}
	//Should be false
	input = "fish"
	expected = false
	actual = IsNumber(input)
	if actual != expected {
		t.Errorf("IsNumber(%s) = %t; want %t", input, actual, expected)
	}
}

func TestGetDomainOrIP(t *testing.T) {
	input := "https://www.google.com"
	expected := "www.google.com"
	actual := GetDomainOrIP(input)
	if actual != expected {
		t.Errorf("GetDomainOrIP(%s) = %s; want %s", input, actual, expected)
	}
}

func TestGetPortString(t *testing.T) {
	input := "https://www.google.com:443"
	expected := "443"
	actual := GetPortString(input)
	if actual != expected {
		t.Errorf("GetPortString(%s) = %s; want %s", input, actual, expected)
	}
}

func TestGetPortInt(t *testing.T) {
	input := "https://www.google.com:443"
	expected := 443
	actual := GetPortInt(input)
	if actual != expected {
		t.Errorf("GetPortInt(%s) = %d; want %d", input, actual, expected)
	}
}

func TestClean(t *testing.T) {
	input := "https://www.google.com:443"
	expected := "httpswwwgooglecom443"
	actual := Clean(input, "")
	if actual != expected {
		t.Errorf("Clean(%s) = %s; want %s", input, actual, expected)
	}
}

func TestInsertString(t *testing.T) {
	//	Write test code here
	input := []string{"Hello", "Fish"}
	expected := []string{"Hello", "World", "Fish"}
	actual := InsertString(input, 1, "World")
	if strings.Join(actual, "") != strings.Join(expected, "") {
		t.Errorf("InsertString(%s) = %s; want %s", input, actual, expected)
	}

}

func TestCastStringToType(t *testing.T) {
	input := "123"
	var expected int64
	expected = 123

	actual_int := CastStringToType(INPUT_TYPE_INT, input).(int64)
	if actual_int != expected {
		t.Errorf("CastStringToType(%s) = %d; want %d", input, actual_int, expected)
	}

	input = "true"
	expected_bool := true
	actual_bool := CastStringToType(INPUT_TYPE_BOOL, input).(bool)
	if actual_bool != expected_bool {
		t.Errorf("CastStringToType(%s) = %v; want %d", input, actual_bool, expected)
	}

	input = "1.56"
	expected_float := 1.56
	actual_float := CastStringToType(INPUT_TYPE_FLOAT, input).(float64)
	if actual_float != expected_float {
		t.Errorf("CastStringToType(%s) = %v; want %d", input, actual_bool, expected)
	}

}

func TestCommaListContainsString(t *testing.T) {
	input := "123,456,789"
	expected := true
	actual := CommaListContainsString(input, "123")
	if actual != expected {
		t.Errorf("CommaListContainsString(%s) = %t; want %t", input, actual, expected)
	}
}

func TestCheckNotBlank(t *testing.T) {
	type args struct {
		value []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "CheckNotBlank",
			args: args{
				value: []string{"Hello", "World"},
			},
			wantErr: nil,
		},
		{
			name: "CheckNotBlank",
			args: args{
				value: []string{"Hello", "", "World"},
			},
			wantErr: Err_ValueIsBlank,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckNotBlank(tt.args.value...); err != tt.wantErr {
				t.Errorf("CheckNotBlank() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGzipBase64(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "GzipBase64",
			args: args{
				data: "Hello World",
			},
			want:    "H4sIAAAAAAAA//JIzcnJVwjPL8pJAQQAAP//VrEXSgsAAAA=",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GzipBase64(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GzipBase64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GzipBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}
