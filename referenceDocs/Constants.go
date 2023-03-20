package main

/* Integet literals
212         /* Legal
215u        /* Legal
0xFeeL      /* Legal
078         /* Illegal: 8 is not an octal digit
032UU       /* Illegal: cannot repeat a suffix

85         /* decimal
0213       /* octal
0x4b       /* hexadecimal
30         /* int
30u        /* unsigned int
30l        /* long
30ul       /* unsigned long

floating point literals
3.14159       /* Legal
314159E-5L    /* Legal
510E          /* Illegal: incomplete exponent
210f          /* Illegal: no decimal or exponent
.e55          /* Illegal: missing integer or fraction

Escape sequences
\\	\ character
\'	' character
\"	" character
\?	? character
\a	Alert or bell
\b	Backspace
\f	Form feed
\n	Newline
\r	Carriage return
\t	Horizontal tab
\v	Vertical tab
\ooo	Octal number of one to three digits
\xhh . . .	Hexadecimal number of one or more digits
*/

import "fmt"

func main() {
	fmt.Printf("Hello\tWorld!\n")

	//const
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("Length: %d, Width: %d, value of area: %d", LENGTH, WIDTH, area)

}
