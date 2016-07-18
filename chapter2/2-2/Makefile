PROG :=	cf2.exe
SUMMARY := "温度、長さ、重さを変換できるcf"

.PHONY: all run clean

.SUFFIXES: .go .exe
all: $(PROG)
run: $(PROG)
	@echo "========サマリー========"
	@echo $(SUMMARY)
	@echo "========実行開始========"
	./$< temp 23.5
	./$< weight 60
	./$< length 120
	@echo ""
.go.exe:
	go build -o $@ $<
clean:
	rm -f *.exe
	find . -name "*~" | xargs rm -f
	find . -name "\#*\#" | xargs rm -f
