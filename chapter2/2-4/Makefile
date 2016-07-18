PROG :=	test2.exe
SUMMARY := "ビットシフトを使ったPopCountと速度比較"

.PHONY: all run clean

.SUFFIXES: .go .exe
all: $(PROG)
run: $(PROG)
	@echo "========サマリー========"
	@echo $(SUMMARY)
	@echo "========実行開始========"
	./$<
	@echo ""
.go.exe:
	go build -o $@ $<
clean:
	rm -f *.exe
	find . -name "*~" | xargs rm -f
	find . -name "\#*\#" | xargs rm -f
