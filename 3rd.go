package kakao2019

import (
	"fmt"
	"sort"
	"strings"
)

// All returns all combinations for a given string array.
// This is essentially a powerset of the given set except that the empty set is disregarded.
func combinations(set []int) (subsets [][]int) {
	length := uint(len(set))

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []int

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets

		// sort. subset
		sort.Slice(subset, func(i, j int) bool {
			return subset[i] < subset[j]
		})

		subsets = append(subsets, subset)
	}
	// sort subsets
	sort.Slice(subsets, func(i, j int) bool {
		if len(subsets[i]) < len(subsets[j]) {
			return true
		} else if len(subsets[i]) > len(subsets[j]) {
			return false
		}

		return subsets[i][0] < subsets[j][0]
	})

	fmt.Println("keys :", subsets)
	return subsets
}

func getkey(row []string, key []int) string {

	var tmp []string
	for _, k := range key {
		tmp = append(tmp, row[k])
	}
	return strings.Join(tmp, ":")
}

func issubset(first, second []int) bool {
	set := make(map[int]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

func k3rd(relation [][]string) int {
	var tempkey []int
	for i := 0; i < len(relation[0]); i++ {
		tempkey = append(tempkey, i)
	}

	allkeys := combinations(tempkey)
	var unikeys [][]int

	result := 0

	for _, ki := range allkeys {
		ignore := false
		for _, uk := range unikeys {
			//fmt.Println("chekc subset", uk, ki, issubset(uk, ki))

			if issubset(uk, ki) {
				fmt.Println("ignore key", uk, ki)
				ignore = true
				break
			}

		}
		if ignore {
			continue
		}

		key := make(map[string]struct{})
		bUnique := true

		for i := 0; i < len(relation); i++ {
			kstr := getkey(relation[i], ki)

			if _, exist := key[kstr]; exist {
				bUnique = false
				break
			} else {
				key[kstr] = struct{}{}
			}
		}
		if bUnique {
			result++
			unikeys = append(unikeys, ki)

			fmt.Println("uni key :", ki)
		}
	}

	return result
}
