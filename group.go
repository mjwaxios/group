package group

import (
	"sort"
)

// Group holds a list of names
type Group struct {
	str []string
}

// String will return the group as a string
func (g Group) String() (s string) {
	for i := range g.str {
		s += g.str[i]
		if i < len(g.str)-1 {
			s += ", "
		}
	}
	return
}

// Sort will sort the name slice,  this will modify the struct it self
func (g Group) Sort() Group {
	sort.Strings(g.str)
	return g
}

// Randomizer takes a list of names and randomizes the order based on hash of pre + name text
func (g Group) Randomizer(hashFun func(string) string, pre string) (out Group) {
	hashMap := make(map[string]string)
	for _, s := range g.str {
		hashMap[hashFun(pre+s)] = s
	}

	out.str = sortMapByKey(hashMap)
	return
}

// New Creates a new Group of strings
func New(s []string) Group {
	return Group{str: s}
}

// sortMapByKey takes a map[string]string and returns the values sorted by key, without the keys
func sortMapByKey(hm map[string]string) []string {
	var s []string
	for key := range hm {
		s = append(s, key)
	}

	sort.Strings(s)

	out := make([]string, len(hm))
	for i, hash := range s {
		out[i] = hm[hash]
	}

	return out
}
