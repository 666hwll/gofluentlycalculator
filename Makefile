INSTALL_DIR = /usr/bin
INNAME = calculator64.go
OUTNAME = gocalc

help:
	@echo "make install 	Install Gocalc"
	@echo "make uninstall 	Remove Gocalc"

install:
	gccgo -o $(OUTNAME) $(INNAME)
	sudo cp $(OUTNAME) $(INSTALL_DIR)

uninstall:
	sudo rm $(INSTALL_DIR)/$(OUTNAME)