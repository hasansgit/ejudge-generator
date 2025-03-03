package test_gen

import "math/rand"

func genInt(a, b int64) int64 {
	return rand.Int63n(b-a+1) + a
}

func genFloat(a, b float64) (x float64) {
	return rand.Float64()*(b-a) + a
}

func genString(r string, sz int) string {
	ps := make([]rune, sz)
	for i := range ps {
		ps[i] = rune(r[rand.Intn(len(r))])
	}
	return string(ps)
}
