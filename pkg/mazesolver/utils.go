package mazesolver

import "log"

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
