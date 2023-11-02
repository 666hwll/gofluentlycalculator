INSTALL_DIR = /usr/bin
StartDIR = /usr/share/applications/
StartNAME = gocalc.desktop
INNAME = calculator64.go
OUTNAME = gocalc

help:
	@echo "make install 	Install Gocalc"
	@echo "make uninstall 	Remove Gocalc"

install:
	gccgo -o $(OUTNAME) $(INNAME)
	sudo cp $(OUTNAME) $(INSTALL_DIR)
	sudo cp $(StartNAME) $(StartDIR) 
uninstall:
	sudo rm $(INSTALL_DIR)/$(OUTNAME)
	sudo rm $(StartDIR)/$(StartNAME)
