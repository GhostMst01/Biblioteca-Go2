package fp

// Filter evalúa un slice usando una función criterio y retorna una colección filtrada nueva.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce procesa un slice acumulando secuencialmente sus valores en un único resultado
func Reduce[T any, U any](slice []T, initial U, accumulator func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = accumulator(result, v)
	}
	return result
}