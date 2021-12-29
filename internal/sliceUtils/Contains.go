package sliceUtils

import "github.com/transip/gotransip/v6/domain"

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsDomain(s []domain.Domain, e string) int {
	for i, a := range s {
		if a.Name == e {
			return i
		}
	}
	return -1
}
