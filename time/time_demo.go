package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func timeRange() {
	start := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)

	/*
		On systems where Unix time is stored as a signed 32-bit integer, the largest value that can be
		recorded is 2147483647 (2^31 − 1), which is 03:14:07 Tuesday, 19 January 2038 UTC.
	*/
	//	Duration receives nanoseconds. 10^-9 seconds
	dur := time.Duration(int64(math.Pow(2, 31)-1) * 1e9)
	fmt.Println(dur)
	end := start.Add(dur) // 2038-01-19 03:14:07 +0000 UTC
	fmt.Println(end)

	dur2 := time.Duration(int64(math.Pow(2, 63) - 1))
	fmt.Println(dur2)
	end2 := start.Add(dur2) // 2038-01-19 03:14:07 +0000 UTC
	fmt.Println(end2)
}

func unixTime() {
	// max1 gets time.Time struct: {-9223371974719179009 999999999}
	max1 := time.Unix(1<<63-1, 999999999)
	fmt.Println(max1)
	// number of seconds between Year 1 and 1970 (62135596800 seconds)
	unixToInternal := int64((1969*365 + 1969/4 - 1969/100 + 1969/400) * 24 * 60 * 60)

	// max2 gets time.Time struct: {9223372036854775807 999999999}
	max2 := time.Unix(1<<63-1-unixToInternal, 999999999)
	fmt.Println(max2)
}

func seedRandom() {
	/*
		UnixNano returns t as a Unix time, the number of nanoseconds elapsed since January 1, 1970 UTC.
		The result is undefined if the Unix time in nanoseconds cannot be represented by an int64
		(a date before the year 1678 or after 2262).
	*/
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 //
	fmt.Println(freq)
}
func main() {
	timeRange()
	unixTime()
}
