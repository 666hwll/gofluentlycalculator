INSTALL_DIR = /usr/bin
StartDIR = /usr/share/applications/
ICODIR = /usr/share/pixmaps/
StartNAME = gocalc.desktop
INNAME = calculator64.go
ICONAME = gocalcico.svg
OUTNAME = gocalc
pref = settings.json
setdirc = .gocalc

help:
	@echo "make install 	Install Gocalc"
	@echo "make uninstall 	Remove Gocalc"

install:
	gccgo -o $(OUTNAME) $(INNAME)
	sudo cp $(OUTNAME) $(INSTALL_DIR)
	sudo cp $(ICONAME) $(ICODIR)
	sudo cp $(StartNAME) $(StartDIR) 
	mkdir -p $(HOME)/$(setdirc)
	cp $(pref) $(HOME)/$(setdirc)
uninstall:
	sudo rm $(INSTALL_DIR)/$(OUTNAME)
	sudo rm $(StartDIR)/$(StartNAME)
	sudo rm $(ICODIR)/$(ICONAME)
	rm -r $(HOME)/$(setdirc)
