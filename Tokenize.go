package main

// TYPES
type TokenKind string

type Token struct {
	kind   TokenKind
	name   string
	values []string
}

// ENUMS
const (
	LITERAL  TokenKind = "literal"
	VARIABLE TokenKind = "variable"
)

const (
	open  string = "{{"
	close string = "}}"
)

// FUNCTION
func tokenize(template string) ([]Token, *map[string][]string) {
	tokenStart := 0
	tokens := []Token{}
	variables := make(map[string][]string)

	for i := 0; i <= len(template)-2; i++ {
		// technically I should be keeping track of open and close and making sure they match
		// I'm not worried about that rn
		if template[i:i+2] == open {
			tokens = append(tokens, Token{
				kind:   LITERAL,
				name:   "",
				values: []string{template[tokenStart:i]},
			})
		} else if template[i:i+2] == close {
			newToken := Token{
				kind:   VARIABLE,
				name:   template[tokenStart:i],
				values: []string{},
			}
			tokens = append(tokens, newToken)
			variables[newToken.name] = []string{}
		} else {
			continue
		}
		tokenStart = i + 2
		i++ // iterate once, for loop will iterate a second time to cover both characters
	}
	// add any final tokens
	if tokenStart < len(template)-1 {
		tokens = append(tokens, Token{
			kind:   LITERAL,
			name:   "",
			values: []string{template[tokenStart:]},
		})
	}

	return tokens, &variables
}
