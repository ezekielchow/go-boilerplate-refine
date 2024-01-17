package utils

import "net/url"

func ExtractFilters(queries url.Values, searchables []string) map[string]string {
	filters := map[string]string{}

	for i := 0; i < len(searchables); i++ {
		field := searchables[i]

		if queries.Has(field + "_like") {
			filters[field] = queries.Get(field + "_like")
		}
	}

	return filters
}
