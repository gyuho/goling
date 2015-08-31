package segment

import "math"

type split struct {
	Head string
	Tail string
}

func doSplit(txt string) []split {
	splits := []split{}
	for i := range txt {
		splits = append(splits, split{txt[:i+1], txt[i+1:]})
	}
	return splits
}

func mostPlausible(chunks [][]string, probFunc func(string) float64) []string {
	chunk := []string{}
	min := -1 * math.MaxFloat64
	for _, words := range chunks {
		score := 1.0
		for _, elem := range words {
			score *= probFunc(elem)
		}
		if min < score {
			min = score
			chunk = words
		}
	}
	return chunk
}

// prev stores previously found segmentations.
var prev = map[string][]string{}

// Get returns the highest-scoring segmentation.
func Get(txt string, probFunc func(string) float64) []string {
	if len(txt) == 0 {
		return []string{}
	}
	if result, ok := prev[txt]; ok {
		return result
	}
	chunks := [][]string{}
	for _, split := range doSplit(txt) {
		chunks = append(chunks,
			append([]string{split.Head},
				Get(split.Tail, probFunc)...,
			),
		)
	}
	rs := mostPlausible(chunks, probFunc)
	prev[txt] = rs
	return rs
}
