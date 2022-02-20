package minimum_number_of_steps_to_make_two_strings_anagram

func minSteps(s string, t string) int {
	freqS := [27]int{}
	freqT := [27]int{}

	for i := range s {
		freqS[s[i]-'a']++
		freqT[t[i]-'a']++
	}

	steps := 0

	for i := range freqS {
		if freqS[i] > freqT[i] {
			steps += freqS[i] - freqT[i]
		}
	}

	return steps
}
