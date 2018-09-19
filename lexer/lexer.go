package lexer

type Lexer struct {
	input string
	position int  // current position in the input which points to the current character
	readPosition int  // current reading position in input after current character
	ch byte  // current character under examination
}


func New(input string) *Lexer {
	l := &Lexer{input}
	return l
}