package pagination

import (
	"fmt"
)

// Footer representes the configuration of the footer
// pages presentation for the webapp
type Footer struct {
	CurrentPage, TotalPages, Boundaries, Around int
}

// ComputePagination calculates the footer pagination offsets
func (f *Footer) ComputePagination() []string {
	pagesWithSeparators := []string{}
	separator := "..."
	pageSet := map[int]int{}

	if f.CurrentPage == 0 || f.TotalPages == 0 {
		return pagesWithSeparators
	}

	pageSet[f.CurrentPage] = 1
	for i := 0; i < f.Boundaries; i++ {
		pageSet[1+i] = 1
		pageSet[f.TotalPages-i] = 1
	}

	for i := 1; i <= f.Around; i++ {
		pageSet[f.CurrentPage+i] = 1
		pageSet[f.CurrentPage-i] = 1
	}

	// At the beggining, no separator is needed because paginations should start with 1
	needsSeparator := false
	for i := 1; i <= f.TotalPages; i++ {
		if pageSet[i] == 1 {
			pagesWithSeparators = append(pagesWithSeparators, fmt.Sprint(i))
			needsSeparator = true
		}
		if pageSet[i] == 0 && needsSeparator {
			pagesWithSeparators = append(pagesWithSeparators, separator)
			needsSeparator = false
		}
	}

	// Remove trailing separators
	if pagesWithSeparators[len(pagesWithSeparators)-1] == separator {
		pagesWithSeparators = pagesWithSeparators[:len(pagesWithSeparators)-1]
	}

	return pagesWithSeparators
}
