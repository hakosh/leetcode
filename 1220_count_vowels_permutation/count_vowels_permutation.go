package count_vowels_permutation

import (
	"math"
	"math/big"
)

var vowels = []rune{'a', 'e', 'i', 'o', 'u'}

var follow = make([][]rune, 127)
var preceed = make([][]rune, 127)

func init() {
	follow[' '] = []rune{'a', 'e', 'i', 'o', 'u'}

	follow['a'] = []rune{'e'}
	follow['e'] = []rune{'a', 'i'}
	follow['i'] = []rune{'a', 'e', 'o', 'u'}
	follow['o'] = []rune{'i', 'u'}
	follow['u'] = []rune{'a'}

	preceed[0] = []rune{1, 2, 4}
	preceed[1] = []rune{0, 2}
	preceed[2] = []rune{1, 3}
	preceed[3] = []rune{2}
	preceed[4] = []rune{2, 3}
}

// BOTTOM UP

func countVowelPermutation(n int) int {
	if n == 0 {
		return 0
	}

	m := int64(1000000007)

	var last [5]int64
	var this [5]int64

	last[0] = 1
	last[1] = 1
	last[2] = 1
	last[3] = 1
	last[4] = 1

	for k := 1; k < n; k++ {
		for i := range vowels {
			var p int64

			for _, j := range preceed[i] {
				p += last[j]
			}

			this[i] = p % m
		}

		last, this = this, [5]int64{}
	}

	var sum int64
	for _, v := range last {
		sum += v
	}

	return int(sum % m)
}

// TOP DOWN

var cache [][]*big.Int

func countVowelPermutationR(n int) int {
	cache = make([][]*big.Int, n+1)
	for i := range cache {
		cache[i] = make([]*big.Int, 127)
	}

	m := big.NewInt(int64(math.Pow10(9)) + 7)
	r := dp(n, ' ')
	r.Mod(r, m)

	return int(r.Int64())
}

func dp(n int, last rune) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}

	if cache[n][last] == nil {
		sum := big.NewInt(0)

		for _, vowel := range follow[last] {
			sum.Add(sum, dp(n-1, vowel))
		}

		cache[n][last] = sum
	}

	return cache[n][last]
}
