package collections

// Index returns the first index of the target interface{} t, or -1 if no match is found.
func Index(vs []interface{}, t interface{}) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include returns true if the target interface{} t is in the slice.
func Include(vs []interface{}, t interface{}) bool {
	return Index(vs, t) >= 0
}

// Any returns true if one of the interface{}s in the slice satisfies the predicate f.
func Any(vs []interface{}, f func(interface{}) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if all of the interface{}s in the slice satisfy the predicate f.
func All(vs []interface{}, f func(interface{}) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing all interface{}s in the slice that satisfy the predicate f.
func Filter(vs []interface{}, f func(interface{}) bool) []interface{} {
	vsf := make([]interface{}, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map returns a new slice containing the results of applying the function f to each interface{} in the original slice.
func Map(vs []interface{}, f func(interface{}) interface{}) []interface{} {
	vsm := make([]interface{}, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
