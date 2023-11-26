# gocalc
A simple cli calculator written in Go
### Goals:
- web server 
- the closest thing to minimalism while having a functioning calculator you could use (only stdlib) 🔥
- replacing the gnome calculator and my Mathprint TI
- c implementation with/out lua plugins; for portability and stability
- optimization of that mess lol
### Installation:
``` bash
make install
```

### new features of the normal calculator:
- cli flags: executing only one statement if the first given number isn't zero
- Tan/Tanh, Sin/Sinh, Cos/Cosh, Log, facility with the Gamma function
- confusion because of the flags usage lol
### what does not work intensionally because minimalism and format:
- Brackets
