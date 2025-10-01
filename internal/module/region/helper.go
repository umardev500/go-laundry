package region

func MapEntity[E any, D any](ents []E, fn func(E) D) []D {
	results := make([]D, len(ents))
	for i, ent := range ents {
		results[i] = fn(ent)
	}
	return results
}
