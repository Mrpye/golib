package convert

import (
	"reflect"
	"testing"
)

func TestToPassFail(t *testing.T) {
	type args struct {
		value bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test ToPassFail() with true",
			args: args{
				value: true,
			},
			want: "Pass",
		},
		{
			name: "Test ToPassFail() with false",
			args: args{
				value: false,
			},
			want: "Fail",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPassFail(tt.args.value); got != tt.want {
				t.Errorf("ToPassFail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToBool(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test ToBool() with true",
			args: args{
				value: true,
			},
			want: true,
		},
		{
			name: "Test ToBool() with false",
			args: args{
				value: false,
			},
			want: false,
		},
		{
			name: "Test ToBool() with 1",
			args: args{
				value: 1,
			},
			want: true,
		},
		{
			name: "Test ToBool() with 0",
			args: args{
				value: 0,
			},
			want: false,
		},
		{
			name: "Test ToBool() with \"yes\"",
			args: args{
				value: "yes",
			},
			want: true,
		},
		{
			name: "Test ToBool() with \"no\"",
			args: args{
				value: "no",
			},
			want: false,
		},
		{
			name: "Test ToBool() with \"y\"",
			args: args{
				value: "y",
			},
			want: true,
		},
		{
			name: "Test ToBool() with \"n\"",
			args: args{
				value: "n",
			},
			want: false,
		},
		{
			name: "Test ToBool() with \"pass\"",
			args: args{
				value: "pass",
			},
			want: true,
		},
		{
			name: "Test ToBool() with \"fail\"",
			args: args{
				value: "fail",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBool(tt.args.value); got != tt.want {
				t.Errorf("ToBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test ToString() with true",
			args: args{
				value: true,
			},
			want: "true",
		},
		{
			name: "Test ToString() with false",
			args: args{
				value: false,
			},
			want: "false",
		},
		{
			name: "Test ToString() with 1",
			args: args{
				value: 1,
			},
			want: "1",
		},
		{
			name: "Test ToString() with 1.25",
			args: args{
				value: 1.25,
			},
			want: "1.25",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.value); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test ToInt() with true",
			args: args{
				value: true,
			},
			want: 1,
		},
		{
			name: "Test ToInt() with false",
			args: args{
				value: false,
			},
			want: 0,
		},
		{
			name: "Test ToInt() with 1",
			args: args{
				value: 1,
			},
			want: 1,
		},
		{
			name: "Test ToInt() with 1.25",
			args: args{
				value: 1.25,
			},
			want: 1,
		},
		{
			name: "Test ToInt() with \"1\"",
			args: args{
				value: "1",
			},
			want: 1,
		},
		{
			name: "Test ToInt() with \"1.25\"",
			args: args{
				value: "1.25",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.value); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt8(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "Test ToInt8() with true",
			args: args{
				value: true,
			},
			want: 1,
		},
		{
			name: "Test ToInt8() with false",
			args: args{
				value: false,
			},
			want: 0,
		},
		{
			name: "Test ToInt8() with 1",
			args: args{
				value: 1,
			},
			want: 1,
		},
		{
			name: "Test ToInt8() with 1.25",
			args: args{
				value: 1.25,
			},
			want: 1,
		},
		{
			name: "Test ToInt8() with \"1\"",
			args: args{
				value: "1",
			},
			want: 1,
		},
		{
			name: "Test ToInt8() with \"1.25\"",
			args: args{
				value: "1.25",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt8(tt.args.value); got != tt.want {
				t.Errorf("ToInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt32(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "Test ToInt32() with true",
			args: args{
				value: true,
			},
			want: 1,
		},
		{
			name: "Test ToInt32() with false",
			args: args{
				value: false,
			},
			want: 0,
		},
		{
			name: "Test ToInt32() with 1",
			args: args{
				value: 1,
			},
			want: 1,
		},
		{
			name: "Test ToInt32() with 1.25",
			args: args{
				value: 1.25,
			},
			want: 1,
		},
		{
			name: "Test ToInt32() with \"1\"",
			args: args{
				value: "1",
			},
			want: 1,
		},
		{
			name: "Test ToInt32() with \"1.25\"",
			args: args{
				value: "1.25",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt32(tt.args.value); got != tt.want {
				t.Errorf("ToInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Test ToInt64() with true",
			args: args{
				value: true,
			},
			want: 1,
		},
		{
			name: "Test ToInt64() with false",
			args: args{
				value: false,
			},
			want: 0,
		},
		{
			name: "Test ToInt64() with 1",
			args: args{
				value: 1,
			},
			want: 1,
		},
		{
			name: "Test ToInt64() with 1.25",
			args: args{
				value: 1.25,
			},
			want: 1,
		},
		{
			name: "Test ToInt64() with \"1\"",
			args: args{
				value: "1",
			},
			want: 1,
		},
		{
			name: "Test ToInt64() with \"1.25\"",
			args: args{
				value: "1.25",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt64(tt.args.value); got != tt.want {
				t.Errorf("ToInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat32(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "Test ToFloat32() with true",
			args: args{
				value: true,
			},
			want: 1,
		},
		{

			name: "Test ToFloat32() with false",
			args: args{
				value: false,
			},
			want: 0,
		},
		{
			name: "Test ToFloat32() with 1",
			args: args{
				value: 1,
			},
			want: 1,
		},
		{
			name: "Test ToFloat32() with 1.25",
			args: args{
				value: 1.25,
			},
			want: 1.25,
		},
		{
			name: "Test ToFloat32() with \"1\"",
			args: args{
				value: "1",
			},
			want: 1,
		},
		{
			name: "Test ToFloat32() with \"1.25\"",
			args: args{
				value: "1.25",
			},
			want: 1.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat32(tt.args.value); got != tt.want {
				t.Errorf("ToFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test ToFloat64() with true",
			args: args{
				value: true,
			},
			want: 1,
		},
		{

			name: "Test ToFloat64() with false",
			args: args{
				value: false,
			},
			want: 0,
		},
		{
			name: "Test ToFloat64() with 1",
			args: args{
				value: 1,
			},
			want: 1,
		},
		{
			name: "Test ToFloat64() with 1.25",
			args: args{
				value: 1.25,
			},
			want: 1.25,
		},
		{
			name: "Test ToFloat64() with \"1\"",
			args: args{
				value: "1",
			},
			want: 1,
		},
		{
			name: "Test ToFloat64() with \"1.25\"",
			args: args{
				value: "1.25",
			},
			want: 1.25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat64(tt.args.value); got != tt.want {
				t.Errorf("ToFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapInterface(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "Test ToMapInterface() with map[string]interface{}",
			args: args{
				value: map[string]interface{}{
					"key":  "value",
					"Test": "test22",
				},
			},
			want: map[string]interface{}{
				"key":  "value",
				"Test": "test22",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMapInterface(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapString(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "Test ToMapString() with map[string]string",
			args: args{
				value: map[string]interface{}{
					"key":  "value",
					"Test": "test22",
				},
			},
			want: map[string]string{
				"key":  "value",
				"Test": "test22",
			},
		},
		{
			name: "Test ToMapString() with map[string]string",
			args: args{
				value: map[string]string{
					"key":  "value",
					"Test": "test22",
				},
			},
			want: map[string]string{
				"key":  "value",
				"Test": "test22",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMapString(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapInt(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Test ToMapInt() with map[string]int",
			args: args{
				value: map[string]interface{}{
					"key":  1,
					"Test": 2,
				},
			},
			want: map[string]int{
				"key":  1,
				"Test": 2,
			},
		},
		{
			name: "Test ToMapInt() with map[string]int",
			args: args{
				value: map[string]int{
					"key":  1,
					"Test": 2,
				},
			},
			want: map[string]int{
				"key":  1,
				"Test": 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMapInt(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapFloat32(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]float32
	}{
		{
			name: "Test ToMapFloat32() with map[string]float32",
			args: args{
				value: map[string]float32{
					"key":  1.25,
					"Test": 2.25,
				},
			},
			want: map[string]float32{
				"key":  1.25,
				"Test": 2.25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMapFloat32(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToMapFloat64(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]float64
	}{
		{
			name: "Test ToMapFloat64() with map[string]float64",
			args: args{
				value: map[string]float64{
					"key":  1.25,
					"Test": 2.25,
				},
			},
			want: map[string]float64{
				"key":  1.25,
				"Test": 2.25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToMapFloat64(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMapFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToArrayInterface(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "Test ToArrayInterface() with []interface{}",
			args: args{
				value: []interface{}{
					"key",
					"value",
					"Test",
					"test22",
				},
			},
			want: []interface{}{

				"key",
				"value",
				"Test",
				"test22",
			},
		},
		{
			name: "Test ToArrayInterface() with []string",
			args: args{
				value: []string{
					"key",
					"value",
					"Test",
					"test22",
				},
			},
			want: []interface{}{
				"key",
				"value",
				"Test",
				"test22",
			},
		},
		{
			name: "Test ToArrayInterface() with []int",
			args: args{
				value: []int{
					1,
					2,
					3,
					4,
				},
			},
			want: []interface{}{
				1,
				2,
				3,
				4,
			},
		},
		{
			name: "Test ToArrayInterface() with []float32",
			args: args{
				value: []float64{
					1.25,
					2.25,
					3.25,
					4.25,
				},
			},
			want: []interface{}{
				1.25,
				2.25,
				3.25,
				4.25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToArrayInterface(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToArrayInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToArrayString(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test ToArrayString() with []string",
			args: args{
				value: []string{
					"key",
					"value",
					"Test",
					"test22",
				},
			},
			want: []string{
				"key",
				"value",
				"Test",
				"test22",
			},
		},
		{
			name: "Test ToArrayString() with []interface{}",
			args: args{
				value: []interface{}{
					"key",
					"value",
					"Test",
					"test22",
				},
			},
			want: []string{
				"key",
				"value",
				"Test",
				"test22",
			},
		},
		{
			name: "Test ToArrayString() with []int",
			args: args{
				value: []int{
					1,
					2,
					3,
					4,
				},
			},
			want: []string{
				"1",
				"2",
				"3",
				"4",
			},
		},
		{
			name: "Test ToArrayString() with []float64",
			args: args{
				value: []float64{
					1.25,
					2.25,
					3.25,
					4.25,
				},
			},
			want: []string{
				"1.25",
				"2.25",
				"3.25",
				"4.25",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToArrayString(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToArrayString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToArrayInt(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{
			name: "Test ToArrayInt() with []int",
			args: args{
				value: []int{
					1,
					2,
					3,
					4,
				},
			},
			want: []int{
				1,
				2,
				3,
				4,
			},
		},
		{
			name: "Test ToArrayInt() with []interface{}",
			args: args{
				value: []interface{}{
					1,
					2,
					3,
					4,
				},
			},
			want: []int{
				1,
				2,
				3,
				4,
			},
		},
		{
			name: "Test ToArrayInt() with []string",
			args: args{
				value: []string{
					"1",
					"2",
					"3",
					"4",
				},
			},
			want: []int{
				1,
				2,
				3,
				4,
			},
		},
		{
			name: "Test ToArrayInt() with []float32",
			args: args{
				value: []float32{
					1.25,
					2.25,
					3.25,
					4.25,
				},
			},
			want: []int{
				1,
				2,
				3,
				4,
			},
		},
		{
			name: "Test ToArrayInt() with []float64",
			args: args{
				value: []float64{
					1.25,
					2.25,
					3.25,
					4.25,
				},
			},
			want: []int{
				1,
				2,
				3,
				4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToArrayInt(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToArrayInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToArrayFloat32(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "Test ToArrayFloat32() with []float32",
			args: args{
				value: []float32{
					1.25,
					2.25,
					3.25,
					4.25,
				},
			},
			want: []float32{
				1.25,
				2.25,
				3.25,
				4.25,
			},
		},
		{
			name: "Test ToArrayFloat32() with []interface{}",
			args: args{
				value: []interface{}{
					1.25,
					2.25,
					3.25,
					4.25,
				},
			},
			want: []float32{
				1.25,
				2.25,
				3.25,
				4.25,
			},
		},
		{
			name: "Test ToArrayFloat32() with []string",
			args: args{
				value: []string{
					"1.25",
					"2.25",
					"3.25",
					"4.25",
				},
			},
			want: []float32{
				1.25,
				2.25,
				3.25,
				4.25,
			},
		},
		{
			name: "Test ToArrayFloat32() with []int",
			args: args{
				value: []int{
					1,
					2,
					3,
					4,
				},
			},
			want: []float32{
				1,
				2,
				3,
				4,
			},
		},
		{
			name: "Test ToArrayFloat32() with []float64",
			args: args{
				value: []float64{
					1.25,
					2.25,
					3.25,
					4.25,
				},
			},
			want: []float32{
				1.25,
				2.25,
				3.25,
				4.25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToArrayFloat32(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToArrayFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToArrayFloat64(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "Test ToArrayFloat32() with []float32",
			args: args{
				value: []float64{
					1.25,
					2.25,
					3.25,
					4.25,
				},
			},
			want: []float64{
				1.25,
				2.25,
				3.25,
				4.25,
			},
		},
		{
			name: "Test ToArrayFloat32() with []interface{}",
			args: args{
				value: []interface{}{
					1.25,
					2.25,
					3.25,
					4.25,
				},
			},
			want: []float64{
				1.25,
				2.25,
				3.25,
				4.25,
			},
		},
		{
			name: "Test ToArrayFloat32() with []string",
			args: args{
				value: []string{
					"1.25",
					"2.25",
					"3.25",
					"4.25",
				},
			},
			want: []float64{
				1.25,
				2.25,
				3.25,
				4.25,
			},
		},
		{
			name: "Test ToArrayFloat32() with []int",
			args: args{
				value: []int{
					1,
					2,
					3,
					4,
				},
			},
			want: []float64{
				1,
				2,
				3,
				4,
			},
		},
		{
			name: "Test ToArrayFloat32() with []float64",
			args: args{
				value: []float64{
					1.25,
					2.25,
					3.25,
					4.25,
				},
			},
			want: []float64{
				1.25,
				2.25,
				3.25,
				4.25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToArrayFloat64(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToArrayFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToArrayBool(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "Test ToArrayBool() with []bool",
			args: args{
				value: []bool{
					true,
					false,
					true,
					false,
				},
			},
			want: []bool{
				true,
				false,
				true,
				false,
			},
		},
		{
			name: "Test ToArrayBool() with []interface{}",
			args: args{
				value: []interface{}{
					true,
					false,
					true,
					false,
				},
			},
			want: []bool{
				true,
				false,
				true,
				false,
			},
		},
		{
			name: "Test ToArrayBool() with []string",
			args: args{
				value: []string{
					"true",
					"false",
					"true",
					"false",
				},
			},
			want: []bool{
				true,
				false,
				true,
				false,
			},
		},
		{
			name: "Test ToArrayBool() with []int",
			args: args{
				value: []int{
					1,
					0,
					1,
					0,
				},
			},
			want: []bool{
				true,
				false,
				true,
				false,
			},
		},
		{
			name: "Test ToArrayBool() with []float64",
			args: args{
				value: []float64{
					1.25,
					2.25,
					3.25,
					4.25,
				},
			},
			want: []bool{
				true,
				true,
				true,
				true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToArrayBool(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToArrayBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetType(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test GetType() with []bool",
			args: args{
				value: []bool{
					true,
					false,
					true,
					false,
				},
			},
			want: "slice",
		},
		{
			name: "Test GetType() with []interface{}",
			args: args{
				value: []interface{}{
					true,
					false,
					true,
					false,
				},
			},
			want: "slice",
		},
		{
			name: "Test GetType() with []string",
			args: args{
				value: []string{
					"true",
					"false",
					"true",
					"false",
				},
			},
			want: "slice",
		},
		{
			name: "Test GetType() with []int",
			args: args{
				value: []int{
					1,
					0,
					1,
					0,
				},
			},
			want: "slice",
		},
		{
			name: "Test GetType() with []float64",
			args: args{
				value: []float64{
					1.25,
					2.25,
					3.25,
					4.25,
				},
			},
			want: "slice",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetType(tt.args.value); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}
