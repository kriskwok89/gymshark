package services

import (
	"sort"
)

type Pack struct {
	Size int `json:"size"`
}

type PackCalculator struct {
	packSizes []int
}

func NewPackCalculator(packSizes []int) *PackCalculator {
	return &PackCalculator{packSizes: packSizes}
}

func (p *PackCalculator) CalculatePacks(order int) []Pack {
	packs := []Pack{}
	sort.Sort(sort.Reverse(sort.IntSlice(p.packSizes)))

	for _, size := range p.packSizes {
		if order <= 0 {
			break
		}
		count := order / size
		if count > 0 {
			for i := 0; i < count; i++ {
				packs = append(packs, Pack{Size: size})
			}
			order %= size
		}
	}

	// If there are remaining items that don't fit exactly into the pack sizes,
	// we need to add an additional pack of the smallest size.
	if order > 0 {
		packs = append(packs, Pack{Size: p.packSizes[len(p.packSizes)-1]})
	}

	return packs
}
