package two_city_scheduling

import "testing"

type Test struct {
	costs [][]int
	want  int
}

var tests = []Test{
	{[][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}, 110},
	{[][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}, 1859},
	{[][]int{{515, 563}, {451, 713}, {537, 709}, {343, 819}, {855, 779}, {457, 60}, {650, 359}, {631, 42}}, 3086},
}

func TestTwoCitySchedCost(t *testing.T) {
	for _, test := range tests {
		if res := twoCitySchedCost(test.costs); res != test.want {
			t.Errorf("for %v wanted %v, got %v", test.costs, test.want, res)
		}
	}
}
