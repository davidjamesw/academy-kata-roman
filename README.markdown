# Go Roman Numerals to Arabic Converter

My first proper go at some Go. 

Splits a given Roman Numeral and loops through the resulting slice in reverse order, adding or subtracting valid next numbers in the sequence.

To run:
* Clone this repository
* Run `go build romanNumerals.go`
* Run `./romanNumerals <INSERT ROMAN NUMERAL OF YOUR CHOICE>`

The result will be printed to your console.

Done as part of a BBC coding challenge, see below for details:

# Academy Coding Challenge #1

## What is this?

The first in a series of occasional coding challenges from Academy Tech Training & D+E! This is intended to be fun, offer an opportunity for deliberate practice, and give a variety of opportunities to engage with it (brief self-study when convenient for you, light competition with your peers, or scheduled collaborative mob-programming events).

Though the puzzle spec may suggest an approach, you can submit a solution using any programming language, paradigm, style, and testing methodology/library! Feel free to set any constraints you like, or challenge your colleagues.

## The challenge

Create a function that converts Roman numerals to Arabic.

### Basics

| I, II, III                | 1, 2, 3          |
| ------------------------- | ---------------- |
| X, XX, XXX                | 10, 20, 30       |
| C, CC, CCC                | 100, 200, 300    |
| M, MM, MMM                | 1000, 2000, 3000 |
| V<br>VI<br>VII<br>VIII    | 5<br>6<br>7<br>8 |
| IV                        | 4                |
| IX<br>XI                  | 9<br>11          |
| L<br>XL                   | 50<br>40         |
| D<br>CD                   | 500<br>400       |
| CM                        | 900              |
| MCMXCIX                   | 1999             |
| No need to go beyond 3999 |                  |

### Additional constraints

- You can’t repeat a Unit character more than 3 times (e.g. IIII, XXXX, etc. are disallowed)
- You can’t repeat a Five character (e.g. VV, LL, DD are disallowed)
- You can only subtract a unit from next highest five/unit, (e.g. IV and IX are allowed, but IL and IM aren’t)
- You can only subtract a single unit (e.g. IIX is disallowed)
- Except for subtractions, all digits are in nonincreasing order
- What happens in the case of an invalid roman numeral like “ABCDE”?
- What happens in the case of a malformed numeral like “ICXXXXIIVV”?
