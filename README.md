# gocalc
A simple cli calculator written in Go
### Goals:
- [x] web server 
- [x] the closest thing to minimalism while having a functioning calculator you could use (only stdlib) 🔥
- [] replacing the gnome calculator and my Mathprint TI
- [] c implementation with/out lua plugins; for portability and stability
- [] optimization of that mess lol
- [] debian package for pi-apps and the aptitude package manager
### Installation:
``` bash
make install
```

### new features of the normal calculator:
- [x] cli flags: executing only one statement if the first given number isn't zero
- [x] Tan/Tanh, Sin/Sinh, Cos/Cosh, Log, facility with the Gamma function
- [x] confusion because of the flags usage lol
- [x] a json for config
### what does not work intensionally because minimalism and format:
- [x] Brackets
- [x] variables
