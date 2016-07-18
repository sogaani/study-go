PROG :=	cf1.exe
SUMMARY := "絶対温度も出力するcf1"

.PHONY: all run clean

.SUFFIXES: .go .exe
all: $(PROG)
run: $(PROG)
	@echo "========サマリー========"
	@echo $(SUMMARY)
	@echo "========実行開始========"
	./$< 23.5
	@echo ""
.go.exe:
	go build -o $@ $<
clean:
	rm -f *.exe
	find . -name "*~" | xargs rm -f
	find . -name "\#*\#" | xargs rm -f
