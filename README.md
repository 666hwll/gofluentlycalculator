# gocalc
A simple cli calculator written in Go,
that only takes 3 args in one statement as input;
to improve readability in the terminal with proactive simplicity
### Goals:
- [x] web server 
- [x] the closest thing to minimalism while having a functioning calculator you could use (only stdlib) 🔥
- [x] prevent someone from using google.com as calculator 
- replacing the gnome calculator and my Mathprint TI
- c implementation for portability and stability or somekind of portable version
- optimization of that mess lol
- debian package for pi-apps and new releases; planning to make a windows version 
### Installation:
``` bash
make install
```
or 
``` bash
sudo dpkg -i gocalc_(whateverarchitecture).deb
```

### new features of the normal calculator:
- [x] cli flags: executing only one statement if the first given number isn't zero
- [x] Tan/Tanh, Sin/Sinh, Cos/Cosh, Log, facility with the Gamma function
- [x] confusion because of the flags usage lol
- [x] a json for config
### what does not work intensionally because minimalism and format:
- [x] Brackets
- [x] variable

