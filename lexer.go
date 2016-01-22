package gon3

type lexer interface {
	nextToken() token
}

func charMatchLexer(name, input string) *charMatchLexer {
	l := &charMatchLexer{
		name:   name,
		input:  input,
		state:  lexDocument,
		tokens: make(chan tokens, 2),
	}
	return l
}

func naiveRegexLexer(name, input string) *naiveRegexLexer {
	l := &naiveRegexLexer{
		name:  name,
		input: input,
	}
	return l
}
