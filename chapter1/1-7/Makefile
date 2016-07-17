PROG :=	fetch1.exe
SUMMARY := "io.Copyを使ったfetch"

.PHONY: all run clean

.SUFFIXES: .go .exe
all: $(PROG)
run: $(PROG)
	@echo "========サマリー========"
	@echo $(SUMMARY)
	@echo "========実行開始========"
	./$< http://google.com
	@echo ""
.go.exe:
	go build -o $@ $<
clean:
	rm -f *.exe
	rm -f *~
	rm -f \#*\#
