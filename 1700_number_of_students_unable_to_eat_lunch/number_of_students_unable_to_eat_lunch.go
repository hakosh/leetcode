package number_of_students_unable_to_eat_lunch

func countStudents(students []int, sandwiches []int) int {
	skip := 0

	for i := 0; i < len(sandwiches)/2; i++ {
		sandwiches[i], sandwiches[len(sandwiches)-1-i] = sandwiches[len(sandwiches)-1-i], sandwiches[i]
	}

	for {
		if len(students) == 0 {
			return 0
		}

		if students[0] == sandwiches[len(sandwiches)-1] {
			students = students[1:]
			sandwiches = sandwiches[:len(sandwiches)-1]
			skip = 0
		} else {
			if skip == len(students) {
				break
			}

			students = append(students[1:], students[0])
			skip++
		}
	}

	return skip
}
