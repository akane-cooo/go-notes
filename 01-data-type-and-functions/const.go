package ch01

/*
Key takeaways:

1. const value should be initialized
2. normally we use uppercase for const variable names
3. for iota, the first value is 0, the second is one, so on so forth
4. if not specified, the const's will inherit expression from the previous const
*/

// const and iota

const (
	A = 20 // 20
	B      // 20
	C      // 20
	D      // 20
)

const (
	NOT_USE = 1 << (10 * iota) // iota=0
	KB                         // iota=1
	MB                         // iota=2
	GB                         // iota=3
)

const (
	S, M = iota + 1, iota + 2 // 1, 2, iota = 0
	K, L                      // 2, 3 iota = 1
	E, G                      // 3, 4 iota = 2
)
