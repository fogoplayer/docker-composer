package main

// TYPES
type Kind int

type Token struct {
	kind   Kind
	name   string
	values []string
}

// ENUMS
const (
	literal Kind = iota
	variable
)

const (
	open  string = "{{"
	close string = "}}"
)

// FUNCTION
func tokenize(template string) []Token {
	tokenStart := 0
	tokens := []Token{}

	for i := 0; i < len(template)-2; i++ {
		// technically I should be keeping track of open and close and making sure they match
		// I'm not worried about that rn
		if template[i:i+2] == open {
			tokens = append(tokens, Token{
				kind:   literal,
				name:   "",
				values: []string{template[tokenStart:i]},
			})
		} else if template[i:i+2] == close {
			tokens = append(tokens, Token{
				kind:   variable,
				name:   template[tokenStart:i],
				values: []string{},
			})
		} else {
			continue
		}
		tokenStart = i + 2
		i++ // iterate once, for loop will iterate a second time to cover both characters
	}
	// add any final tokens
	if tokenStart < len(template)-1 {
		tokens = append(tokens, Token{
			kind:   literal,
			name:   "",
			values: []string{template[tokenStart:]},
		})
	}

	return tokens
}