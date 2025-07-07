package mapext

// Map allows mapping of a map[K]V -> U.
func Map[K comparable, V any, U any](m map[K]V, init U, fn func(accum U, key K, value V) U) U {
	accum := init
	for k, v := range m {
		accum = fn(accum, k, v)
	}

	return accum
}
