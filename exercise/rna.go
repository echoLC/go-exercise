package exercise

import (
	"go-exercise/utils"
	"strings"
)

func toRNAItem (key string) (string) {
	dnaMap := map[string]string{"G": "C", "C": "G", "T": "A", "A": "U"}
	return dnaMap[key] 
}

func ToRNA(dna string) string {

	dnsSlice := strings.Split(dna, "")

	if dna == "" {
			return dna
	}

	toSlice := utils.Map(dnsSlice, toRNAItem)

	return strings.Join(toSlice, "")
}