Program that simulates a basic interactive bitmap editor.

### Commands

- I M N : Creates a new M x N image with all pixels coloured white (O).
- C : Clears the table, setting all pixels to white (O).
- L X Y C : Colours the pixel (X,Y) with colour C.
- V X Y1 Y2 C : Draws a vertical segment of colour C in column X between rows Y1 and Y2 (inclusive).
- H X1 X2 Y C : Draws a horizontal segment of colour C in row Y between columns X1 and X2 (inclusive).
- S : Shows the contents of the current image.

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

### Building, testing etc

```
$ make help
Usage:
     bin  ................ build the binary (goes to ./bitmap)
     dep  ................ update dependencies
     test ................ run all tests (requires 'ginkgo')
     fake ................ regenerate interface mocks for testing
```
