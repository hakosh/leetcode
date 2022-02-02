package first_bad_version

var versions []bool

func setVersions(vs []bool) {
	versions = vs
}

func isBadVersion(n int) bool {
	return versions[n-1]
}

func firstBadVersion(n int) int {
	low := 0
	mid := 0
	high := n + 1

	for high-low > 1 {
		mid = (low + high) / 2

		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid
		}
	}

	if isBadVersion(mid) {
		return mid
	} else {
		return mid + 1
	}
}
