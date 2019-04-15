package exe7_8_10

import (
	"sort"
	"testing"
)

func TestMultiTierSorting(t *testing.T) {
	printTracks(guiArray)
	myC := customSort{guiArray, func(c1, c2 *guiTables) bool {
		if c1.firstColumn != c2.firstColumn {
			return c1.firstColumn < c2.firstColumn
		}
		if c1.secondColumn != c2.secondColumn {
			return c1.secondColumn < c2.secondColumn
		}

		if c1.thirdColumn != c2.thirdColumn {
			return c1.thirdColumn < c2.thirdColumn
		}
		return c1.fourthColumn < c2.fourthColumn
	}}
	sort.Sort(&myC)
	printTracks(guiArray)

	sort.Sort(sort.Reverse(&myC))
	printTracks(guiArray)
}
