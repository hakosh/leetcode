package valid_mountain_array

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	i := 0

	for {
		if i == len(arr)-1 {
			return false
		}

		if arr[i] > arr[i+1] {
			break
		}

		if arr[i] == arr[i+1] {
			return false
		}

		i++
	}

	if i == 0 {
		return false
	}

	for {
		if i == len(arr)-1 {
			break
		}

		if arr[i] <= arr[i+1] {
			return false
		}

		i++
	}

	return true
}

//func validMountainArray(arr []int) bool {
//	if len(arr) < 3 {
//		return false
//	}
//
//	li := 0
//	ri := len(arr) - 1
//
//	for {
//		l := true
//
//		if li < len(arr)-2 && arr[li] < arr[li+1] {
//			li++
//		} else if li == 0 {
//			return false
//		} else {
//			l = false
//		}
//
//		if ri > 0 && arr[ri] < arr[ri-1] {
//			ri--
//		} else if ri == len(arr)-1 {
//			return false
//		} else if !l {
//			break
//		}
//	}
//
//	return ri-li <= 1
//}
