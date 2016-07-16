subdirs := $(shell ls chapter*)
.PHONY: all run clean

all:
	@LIST="$(subdirs)";\
	for x in $$LIST; do\
	 $(MAKE) -C $$x all;\
	done
run:
	@LIST="$(subdirs)";\
	for x in $$LIST; do\
	 $(MAKE) -C $$x run;\
	done
clean:
	@LIST="$(subdirs)";\
	for x in $$LIST; do\
	 $(MAKE) -C $$x clean;\
	done
