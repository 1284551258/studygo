package split_string

import (
	"reflect"
	"testing"
)

// func TestSplit(t *testing.T) {

// 	ret := Split("abcd", "b")
// 	want := []string{"a", "cde"}
// 	if !reflect.DeepEqual(ret, want) {
// 		t.Errorf("fail! want:%v,but got:%v", want, ret)
// 	}
// }

// func TestSplit(t *testing.T) {

// 	type testCase struct {
// 		str string
// 		req string
// 		ret []string
// 	}

// 	testGroup := []testCase{
// 		testCase{"abcd", "b", []string{"a", "cd"}},
// 		testCase{"abcd", "bc", []string{"a", "d"}},
// 		testCase{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
// 	}

// 	for _, tc := range testGroup {
// 		ret := Split(tc.str, tc.req)
// 		if !reflect.DeepEqual(ret, tc.ret) {
// 			t.Fatalf("fail! want:%v,but got:%v", tc.ret, ret)
// 		}
// 	}
// }

func TestSplit(t *testing.T) {

	type testCase struct {
		str string
		req string
		ret []string
	}

	testGroup := map[string]testCase{
		"case_1": testCase{"abcd", "b", []string{"a", "cd"}},
		"case_2": testCase{"abcd", "bc", []string{"a", "d"}},
		"case_3": testCase{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}

	for name, tc := range testGroup {
		t.Run(name, func(t *testing.T) {
			ret := Split(tc.str, tc.req)
			if !reflect.DeepEqual(ret, tc.ret) {
				t.Fatalf("fail! want:%v,but got:%v", tc.ret, ret)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d", ":")
	}
}
