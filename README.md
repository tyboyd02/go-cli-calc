# Calculator Program

This calculator program evaluates mathematical expressions by following the PEMDAS order of operations. 
It utilizes the Shunting Yard algorithm to convert [Infix notation](https://en.wikipedia.org/wiki/Infix_notation) `4.2+18/(9-3)` to [Reverse Polish notation](https://en.wikipedia.org/wiki/Reverse_Polish_notation) `4.2 18 9 3 - / +` before solving.

## Supported Symbols

- Parentheses: `(`, `)`
- Operators: 
	- `^` (exponentiation)
	- `*` (multiplication)
	- `/` (division)
	- `+` (addition)
	- `-` (subtraction)

## Usage Examples
```plaintext 
go-cli-calc 2+2*2
	OUTPUT: 6
	
go-cli-calc "4.2 + 18 / (9 - 3)"
	OUTPUT: 7.2

go-cli-calc "(1-.5)+2*9+2"
	OUTPUT: 20.666666666666668
```

## NOTES

- Whitespace between numbers and operators will not affect the calculation
- You *may* be able to omit the quotes "" if there is no white space between the numbers and operators

## Known issues 
 
- Can't use negative numbers `2 + -1` or `2^-2`
