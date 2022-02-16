package sort_characters_by_frequency

import "testing"

type Test struct {
	s   string
	out string
}

var tests = []Test{
	{"tree", "eetr"},
}

func TestFrequencySort(t *testing.T) {
	for _, test := range tests {
		if res := frequencySortHeap(test.s); res != test.out {
			t.Errorf("For %v expected %v, got %v", test.s, test.out, res)
		}
	}
}

func BenchmarkFrequencySort(b *testing.B) {
	s := "C3UJ6yMnKbAiErai4hcLHNoc2GSCEed2MPcCJnAMx0ZSwCN9ANTZoyp9rka8A0gFxFH26YMbvLB1ZjVAV5W3MtrpImPLIJzcJWAUHDBR1mN5G1hkfSkaC7M1sDhqE5EcZcVjIo8fcstFkTwCGD663pIv9fZFdeILnZkx8fA8y44DI06H7cCakCfnoNGaXGLiswycxqhrzKPi6zusFOPOyvcoM4dZT2lVkyCnX6R47FcB54QJMuOs8FxrjkcLswhVM3sHb7pCGuGkSS2IO4t6qvSZKTlplS6s0T0plJ4bp2ieX9ceSOheSJGNAy3gryHZ3L6mMXg5x7EmL7OeByWTlH3UKaRWBgImxzzvOCb0Fb23aDWHxqaHcmcknwkelddrJjg6VF3JQfQUDcoDjNPOrNRqLhH4zjGtfMCybcDPbakHEz4MbnnERIghda0Up07Tkb7kUddQxtK43CotAXLYPOsY0emaJIpB8oPHkPRrOIZz9VCKWnYJzHsRILLm1bOmZbZ3UQ1LnfTHE1NBIUidTDcUWlJdF9ETSwBv7AHjfoJm1gmJqoSH8QIQ7Pm0axgRWIWMH5NZFA2XAOmIi6qE9J54AykdG238BhVjSDAGWOTgBvqVxLAtP34nKOiUDLQhqRuMfJubG2S1DIiL6PdGY30F010nE0KZLHP4Rt"
	for i := 0; i < b.N; i++ {
		frequencySortHeap(s)
	}
}
