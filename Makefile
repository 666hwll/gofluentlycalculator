INSTALL_DIR = /usr/bin
sourcemirror = https://github.com/666hwll/gocalc
StartDIR = /usr/share/applications/
ICODIR = /usr/share/pixmaps/
StartNAME = gocalc.desktop
INNAME = calculator64.go
ICONAME = gocalcico.svg
OUTNAME = gocalc
pref = settings.json
setdirc = /.config/gocalc

help:
	@echo "make install 	Install Gocalc"
	@echo "make uninstall 	Remove Gocalc"

install:
	@gccgo -o $(OUTNAME) $(INNAME)
	@sudo cp $(OUTNAME) $(INSTALL_DIR)
	@sudo cp $(ICONAME) $(ICODIR)
	@sudo cp $(StartNAME) $(StartDIR) 
	@mkdir -p $(HOME)/$(setdirc)
	@cp $(pref) $(HOME)/$(setdirc)
	@cp Makefile $(HOME)/$(setdirc)
	@echo "finished installation process"
uninstall:
	@sudo rm $(INSTALL_DIR)/$(OUTNAME)
	@sudo rm $(StartDIR)/$(StartNAME)
	@sudo rm $(ICODIR)/$(ICONAME)
	@rm -r $(HOME)/$(setdirc)

dependencies:
	@read -p  "works only on debian-based systems with APT; continue? [y/n]" answer; \
	if [ "$$answer" = "y" ]; then \
		sudo apt install gccgo; \
	else \
		echo "Aborting..."; \
		exit 1; \
	fi


upgrade:
	uninstall
	current_dir := $(shell pwd)
	cd ..
	rm -r $(current_dir)
	git clone $(sourcemirror)
