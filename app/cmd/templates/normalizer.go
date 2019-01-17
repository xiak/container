package templates

import "strings"

const IndentString = `  `

type normalizer struct {
	string
}

func (this normalizer) trim() normalizer {
	this.string = strings.TrimSpace(this.string)
	return this
}

func (this normalizer) trimAll() normalizer {
	indented := []string{}
	this.string = strings.TrimSpace(this.string)
	for _, line := range strings.Split(this.string, "\n") {
		t := strings.TrimSpace(line)
		indented = append(indented, t)
	}
	this.string = strings.Join(indented, "\n")
	return this
}

func (this normalizer) indent() normalizer {
	indented := []string{}
	for _, line := range strings.Split(this.string, "\n") {
		t := strings.TrimSpace(line)
		i := IndentString + t
		indented = append(indented, i)
	}
	this.string = strings.Join(indented, "\n")
	return this
}

func Long(s string) string {
	if len(s) == 0 {
		return s
	}
	return normalizer{s}.trimAll().string
}

func Example(s string) string {
	if len(s) == 0 {
		return s
	}
	return normalizer{s}.trim().indent().string
}

