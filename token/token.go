package token 

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

// Tokens
const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifers and literals
	IDENT = "IDENT"
	INT = "INT"

	//Operators
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	BANG = "!"
	ASTERISK = "*"
	SLASH = "/"
	LT = "<"
	GT = ">"
	EQ = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"

)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType{
	// this takes 2 values from the map tok which is the value of the map
	// and a bool of whether the value is in the map (ok). So this statement
	// is saying tok and ok from keywords[ident] and if ok is true, continue
	if tok, ok := keywords[ident]; ok{
		return tok
	}
	return IDENT
}