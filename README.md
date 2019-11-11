## Building, testing etc

```
$ make help
Usage:
     bin  ................ build the binary (goes to ./bitmap)
     dep  ................ update dependencies
     test ................ run all tests (requires 'ginkgo')
     fake ................ regenerate interface mocks for testing
```

### Alternative design

I had a last minute change of heart about many things, an earlier version of the code can be found at commit [7d2fc40](https://github.com/mo-work/go-technical-test-for-claudia/tree/7d2fc402a6f339a6037d97b6796c79aae60d58ea).


# Go Technical Assignment

This assignment is meant to evaluate the golang proficiency of software engineers.
Your code should follow best practices and our evaluation will focus primarily on correctness and completeness of implementation. During the face to face interview you will have the opportunity to explain your design choices.

## Technical challenge

Go program that simulates a basic interactive bitmap editor. Bitmaps are represented as a matrix of pixels with each element representing a colour.

### Commands
There are 6 supported commands:

- I M N : Creates a new M x N image with all pixels coloured white (O).
- C : Clears the table, setting all pixels to white (O).
- L X Y C : Colours the pixel (X,Y) with colour C.
- V X Y1 Y2 C : Draws a vertical segment of colour C in column X between rows Y1 and Y2 (inclusive).
- H X1 X2 Y C : Draws a horizontal segment of colour C in row Y between columns X1 and X2 (inclusive).
- S : Shows the contents of the current image.

### Conditions

- Please show relevant error message for error
- M and N can be between 1 <= M,N <= 1024

### Example

*Input:*
```
I 5 5
L 1 3 A
V 2 3 5 W
H 3 5 2 Z
S
```

*Output:*
```
OOOOO
OOZZZ
AWOOO
OWOOO
OWOOO
```

## Evaluation points in order of importance

- Use of clean code which is self documenting
- Use of golang idiomatic principles
- Tests for logic
- Use of code quality checkers such as linters and build tools
- Use of git with appropriate commit messages
- Write re-usable code
- Commit as often as you can, we want to see how you develop and implement your idea