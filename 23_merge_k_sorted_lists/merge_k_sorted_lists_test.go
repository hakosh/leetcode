package merge_k_sorted_lists

import (
	"reflect"
	"testing"
)

type Test struct {
	lists []*ListNode
	out   []int
}

var list1 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 4,
		Next: &ListNode{
			Val:  5,
			Next: &ListNode{Val: 6},
		},
	},
}

var list2 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val:  3,
		Next: &ListNode{Val: 4},
	},
}

var list3 = &ListNode{
	Val: 2,
	Next: &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  10,
				Next: &ListNode{Val: 12},
			},
		},
	},
}

var list4 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
		},
	},
}

var tests = []Test{
	{[]*ListNode{}, []int{}},
	{[]*ListNode{list4}, []int{1, 2, 3}},
	{[]*ListNode{list1, list2, list3, list4}, []int{1, 1, 1, 2, 2, 3, 3, 4, 4, 5, 6, 6, 9, 10, 12}},
}

func TestMergeKLists(t *testing.T) {
	for _, test := range tests {
		res := mergeKLists(test.lists)

		if !reflect.DeepEqual(traverse(res), test.out) {
			t.Errorf("For %v expected %v, got %v", test.lists, test.out, traverse(res))
		}
	}
}

func BenchmarkMergeKLists(b *testing.B) {
	test := tests[len(tests)-1]
	for i := 0; i < b.N; i++ {
		mergeKLists(test.lists)
	}
}

func traverse(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	return append([]int{head.Val}, traverse(head.Next)...)
}
