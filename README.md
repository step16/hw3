# Sample Go Code for HW3

calc.go is intended to be a "same flavor" of lexing/parsing calculator as
in
https://github.com/xharaken/step2015/blob/master/calculator_modularize_2.py

**Please note that this is not a very canonical "Go" style.**

This intentionally follows
https://github.com/xharaken/step2015/blob/master/calculator_modularize_2.py
nearly line-by-line and avoids concepts like interfaces, functors, or
slices of strings.

If you're curious, there's a very good talk by Rob Pike (one of the
authors of Go) about a much cleaner way to write lexers like this,
explaining how https://golang.org/pkg/text/template/ works internally
at https://www.youtube.com/watch?v=HxaD_trXwRE (with slides at
https://talks.golang.org/2011/lex.slide ).

## Want to learn Go?

Start by checking out ["A Tour of Go"](https://tour.golang.org/)!
