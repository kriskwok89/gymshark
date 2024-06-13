package config

import (
	"strconv"
	"strings"

	"github.com/magiconair/properties"
)

type Config struct {
	PackSizes []int
}

func LoadConfig(filePath string) Config {
	p := properties.MustLoadFile(filePath, properties.UTF8)
	packSizesStr := p.GetString("packSizes", "250,500,1000,2000,5000")
	packSizesStrArr := strings.Split(packSizesStr, ",")
	packSizes := make([]int, len(packSizesStrArr))
	for i, sizeStr := range packSizesStrArr {
		size, err := strconv.Atoi(sizeStr)
		if err != nil {
			packSizes = []int{250, 500, 1000, 2000, 5000}
			break
		}
		packSizes[i] = size
	}
	return Config{PackSizes: packSizes}
}
