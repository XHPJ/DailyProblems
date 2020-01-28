package main

func LongestSubstring(k int, s string) string {
	if len(s) == 0 || k >= len(s) {
		return s
	} else if k == 1 {
		return string(s[0])
	}

	uniqueChar := 0
	knownChars := NewSet()

	remainingString := ""

	firstChar := s[0]
	secondIndex := 0

	for s[secondIndex] == firstChar {
		secondIndex++
	}

	candidate := s
	for index, char := range s {
		if !knownChars.Contains(string(char)) {
			knownChars.Add(string(char))
			uniqueChar++
		}
		if uniqueChar > k {
			candidate = string(s[:index])
			remainingString = string(s[secondIndex:])
			break
		}
	}

	longestRemaining := LongestSubstring(k, remainingString)

	if len(candidate) < len(longestRemaining) {
		return longestRemaining
	}
	return candidate
}

var exists = struct{}{}

type set struct {
	m map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(value string) {
	s.m[value] = exists
}

func (s *set) Remove(value string) {
	delete(s.m, value)
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}
