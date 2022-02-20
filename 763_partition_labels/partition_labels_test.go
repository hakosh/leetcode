package partition_labels

import (
	"reflect"
	"testing"
)

type Test struct {
	s   string
	out []int
}

var tests = []Test{
	{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
}

func TestPartitionLabels(t *testing.T) {
	for _, test := range tests {
		if res := partitionLabels(test.s); !reflect.DeepEqual(res, test.out) {
			t.Errorf("For %v expected %v, got %v", test.s, test.out, res)
		}
	}
}
