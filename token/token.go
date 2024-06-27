package token

type TokenType string

/*
Structure of a Token:
- Type
- Literal
*/
type Token struct {
	Type    TokenType // Type of the token (e.g., identifier, operators, delimeters, etc).
	Literal string    // Literal value of the token from the source code.
}

const (
	ILLEGAL = "ILLEGAL" // The token which is not recongnized is assigned as illegal.
	EOF     = "EOF"     // The end of the file.

	// Identifiers and literals.
	IDENT = "IDENT" // Identifiers (e.g., num, foo, add, etc).
	INT   = "INT"   // Integer type literals (e.g., 1, -19, 0, etc).

	// Operators.
	ASSIGN  = "=" // Assignment Operator (use to assign new variables).
	PLUS    = "+" // Plus Operator (use to perform arithmetic addition).
	MINUS   = "-" // Minus Operator (use to perform arithmetic subtraction).
	ASTERIK = "*" // Asterik Operator (use to perform arithmetic multiplication).
	SLASH   = "/" // Slash Operator (use to perform arithmetic divison).
	BANG    = "!" // Bang Operator (Logical Operator).

	LT = "<" // Lesser Than Operator (Logical Operator).
	GT = ">" // Greater Than Operator (Logical Operator).

	EQ = "==" // Equality Operator (Logical Operator).
	NQ = "!=" // Not Equality Operator (Logical Operator).

	// Delimeters.
	COMMA     = "," // Comma (use in arrays, sets, etc to add distinction between the values).
	SEMICOLON = ";" // Semicolon (use to break the line of code).
	LPAREN    = "(" // Left Parenthesis
	RPAREN    = ")" // Right Parenthesis
	LBRACE    = "{" // Left Braces
	RBRACE    = "}" // Right Braces

	// Keywords
	FUNCTION = "FUNCTION" // To create new functions.
	LET      = "LET"      // To create new variables.
	TRUE     = "TRUE"     // Boolean True.
	FALSE    = "FALSE"    // Boolean False.
	RETURN   = "RETURN"   // Return statement (use to return a value).
	IF       = "IF"       // Conditional if.
	ELSE     = "ELSE"     // Conditional else.
)

/*
Reserved keywords by programming language.
*/
var Keywords = map[string]TokenType{
	"let":    LET,
	"fn":     FUNCTION,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
}

/*
Helper function use to look up that the given string (input) is reserved keyword or indentifier (output).
*/
func LookupIdent(ident string) TokenType {
	if tokType, ok := Keywords[ident]; ok {
		return tokType
	}

	return IDENT
}
