package split_string

import (
	"reflect"
	"testing"
)

// 测试组
//func TestSplit(t *testing.T) {
//	type testCase struct {
//		str	string
//		sep string
//		want []string
//	}
//	testGroup := []testCase{
//		testCase{"bbdbfds", "b", []string{"", "", "d", "fds"}},
//		testCase{"a:b:c", ":", []string{"a", "b", "c"}},
//		testCase{"abcef", "bce", []string{"a", "f"}},
//		testCase{"ab中文cef", "b", []string{"a", "中文cef"}},
//	}
//	for _, tc := range testGroup{
//		got := Split(tc.str, tc.sep)
//		if !reflect.DeepEqual(got, tc.want){
//			t.Fatalf("want:%#v got:%#v\n", tc.want, got)
//		}
//	}
//}

// 子测试
func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"case_1": testCase{"bbdbfds", "b", []string{"", "", "d", "fds"}},
		"case_2": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"case_3": testCase{"abcef", "bce", []string{"a", "f"}},
		"case_4": testCase{"ab中文cef", "b", []string{"a", "中文cef"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.str, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("want:%#v got:%#v\n", tc.want, got)
			}
		})
	}
}

// BenchamsrtSplit 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}
