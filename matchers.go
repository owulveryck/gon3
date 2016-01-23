package gon3

import (
	"github.com/rychipman/easylex"
)

var (
	matchPeriod       = easylex.NewMatcher().AcceptRunes(".")
	matchComma        = easylex.NewMatcher().AcceptRunes(",")
	matchSemicolon    = easylex.NewMatcher().AcceptRunes(";")
	matchOpenBracket  = easylex.NewMatcher().AcceptRunes("[")
	matchCloseBracket = easylex.NewMatcher().AcceptRunes("]")
	matchOpenParens   = easylex.NewMatcher().AcceptRunes("(")
	matchCloseParens  = easylex.NewMatcher().AcceptRunes(")")
	matchWhitespace   = easylex.NewMatcher().AcceptRunes("\u0020\u0009\u000D\u000A")

	matchTrue  = easylex.NewMatcher().AcceptString("true")
	matchFalse = easylex.NewMatcher().AcceptString("false")
	matchA     = easylex.NewMatcher().AcceptRunes("a")

	matchEscapable    = easylex.NewMatcher().AcceptRunes("_~.-!$&'()*+,;=/?#@%")
	matchHex          = easylex.NewMatcher().AcceptRunes("0123456789abcdefABCDEF")
	matchDigits       = easylex.NewMatcher().AcceptRunes("0123456789")
	matchAlphabet     = easylex.NewMatcher().AcceptRunes("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	matchAlphaNumeric = easylex.NewMatcher().AcceptRunes("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

	matchPNCharsBase = easylex.NewMatcher().AcceptRunes("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz").AcceptUnicodeRange(rune(0x00C0), rune(0x00D6)).AcceptUnicodeRange(rune(0x00D8), rune(0x00F6)).AcceptUnicodeRange(rune(0x00F8), rune(0x02FF)).AcceptUnicodeRange(rune(0x0370), rune(0x037D)).AcceptUnicodeRange(rune(0x037F), rune(0x1FFF)).AcceptUnicodeRange(rune(0x200C), rune(0x200D)).AcceptUnicodeRange(rune(0x2170), rune(0x218F)).AcceptUnicodeRange(rune(0x2C00), rune(0x2FEF)).AcceptUnicodeRange(rune(0x3001), rune(0xD7FF)).AcceptUnicodeRange(rune(0xF900), rune(0xFDCF)).AcceptUnicodeRange(rune(0xFDF0), rune(0xFFFD)).AcceptUnicodeRange(rune(0x10000), rune(0xEFFFF))
	matchPNCharsU    = easylex.NewMatcher().Union(matchPNCharsBase).AcceptRunes("_")
	matchPNChars     = easylex.NewMatcher().Union(matchPNCharsU).Union(matchDigits).AcceptRunes("-\u00B7\u203F\u2040").AcceptUnicodeRange(rune(0x0300), rune(0x036F))
)