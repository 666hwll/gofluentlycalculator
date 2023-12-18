INSTALL_DIR = /usr/bin
StartDIR = /usr/share/applications/
ICODIR = /usr/share/pixmaps/
StartNAME = gocalc.desktop
INNAME = calculator64.go
ICONAME = gocalcico.svg
setdirc = /.config/gocalc

help:
	@echo "make install 	Install Gocalc"
	@echo "make uninstall 	Remove Gocalc"


install:
	@gccgo -o gocalc $(INNAME)
	@sudo cp gocalc $(INSTALL_DIR)
	@sudo cp $(ICONAME) $(ICODIR)
	@sudo cp $(StartNAME) $(StartDIR) 
	@mkdir -p $(HOME)/$(setdirc)
	@sudo cp settings.json $(HOME)/$(setdirc)
	@sudo cp Makefile $(HOME)/$(setdirc)
	@echo "finished installation process"
uninstall:
	@sudo rm $(INSTALL_DIR)/gocalc
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

deb:
	dpkg-deb --build --root-owner-group debianrel
	mv debianrel.deb gocalc.deb

debinstall:
	sudo dpkg -i gocalc.deb

debuninstall:
	sudo dpkg -r gocalc

crcptoarm64:
	GOARCH=arm64 CC=aarch64-linux-gnu-gccgo go build -o gocalc calculator64.go

