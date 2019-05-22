package stringstudy

import "strconv"

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		// int to string
		result = strconv.Itoa(lsb) + result
	}
	return result
}
