package map_utils

import (
	"reflect"
	"testing"
)

func TestCopyMap(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}

	//create a map with a nested map to test copy
	test_map := map[string]interface{}{
		"key1": "value1",
		"key2": map[string]interface{}{
			"key3": "value3",
			"key4": map[string]interface{}{
				"key5": "value5",
			},
		},
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "empty slice",
			args: args{
				m: test_map,
			},
			want: test_map,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CopyMap(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopyMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
