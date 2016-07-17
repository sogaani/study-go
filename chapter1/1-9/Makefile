PROG :=	fetch3.exe
SUMMARY := "httpのステータスコードも出力するfetch"

.PHONY: all run clean

.SUFFIXES: .go .exe
all: $(PROG)
run: $(PROG)
	@echo "========サマリー========"
	@echo $(SUMMARY)
	@echo "========実行開始========"
	./$< google.com
	@echo ""
.go.exe:
	go build -o $@ $<
clean:
	rm -f *.exe
	rm -f *~
	rm -f \#*\#
