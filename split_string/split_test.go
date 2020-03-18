package splitstring

import (
	"reflect"
	"testing"
)

func Test1Split(t *testing.T) {
	ret := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(ret, want) {
		//测试用例失败
		t.Errorf("want:%v but got:%v\n", want, ret)
	}
}

func Test2Split(t *testing.T) {
	ret := Split("abcef", "bc")
	want := []string{"a", "ef"}
	if !reflect.DeepEqual(ret, want) {
		//测试用例失败
		t.Errorf("want:%v but got:%v\n", want, ret)
	}
}

func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	// testGroup := []testCase{
	// 	testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
	// 	testCase{"a:b:c", ":", []string{"a", "b", "c"}},
	// 	testCase{"abcef", "bc", []string{"a", "ef"}},
	// }

	// for _, tc := range testGroup {
	// 	got := Split(tc.str, tc.sep)
	// 	want := []string{"a", "ef"}
	// 	if !reflect.DeepEqual(got, tc.want) {
	// 		//测试用例失败
	// 		t.Errorf("want:%v but got:%v\n", want, got)
	// 	}
	// }

	testGroup := map[string]testCase{
		"case_1": testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case_2": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"case_3": testCase{"abcef", "bc", []string{"a", "ef"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			want := []string{"a", "ef"}
			if !reflect.DeepEqual(got, tc.want) {
				//测试用例失败
				t.Errorf("want:%v but got:%v\n", want, got)
			}
		})
	}
}

// 性能基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
