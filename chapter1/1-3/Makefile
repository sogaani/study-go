PROG :=	echo3.exe
SUMMARY := "非効率なechoと効率的なechoの速度比較"

INPUT := $(shell cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 16 | head -n 1000 | sort | uniq)

.PHONY: all run clean

.SUFFIXES: .go .exe
all: $(PROG)
run: $(PROG)
	@echo "========サマリー========"
	@echo $(SUMMARY)
	@echo "========実行開始========"
	./$< $(INPUT)
	@echo ""
.go.exe:
	go build -o $@ $<
clean:
	rm -f *.exe
	rm -f *~
	rm -f \#*\#
