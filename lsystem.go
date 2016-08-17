package main

import "bytes"

type Productions map[rune]string

func (p Productions) Apply(input string) string {
	buffer := bytes.Buffer{}

	for i := range input {
		current := rune(input[i])
		production, ok := p[current]
		if !ok {
			buffer.WriteRune(current)
			continue
		}

		buffer.WriteString(production)
	}
	return buffer.String()
}

func (p Productions) ApplyTimes(input string, n int) string {
	result := input
	for i := 0; i < n; i++ {
		result = p.Apply(result)
	}

	return result
}
