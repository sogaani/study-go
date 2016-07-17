PROG :=	fetchall1.exe
SUMMARY := "ファイルへ保存するfetchall"

.PHONY: all run clean

.SUFFIXES: .go .exe
all: $(PROG)
run: $(PROG)
	@echo "========サマリー========"
	@echo $(SUMMARY)
	@echo "========実行開始1回目========"
	./$< https://www.ricoh.co.jp/pdf/34211620/MP_C8002_C6502.pdf
	mv MP_C8002_C6502.pdf first
	@echo ""
	@echo "========実行開始2回目========"
	./$< https://www.ricoh.co.jp/pdf/34211620/MP_C8002_C6502.pdf
	mv MP_C8002_C6502.pdf second
	@echo ""
	diff first second
.go.exe:
	go build -o $@ $<
clean:
	rm -f *.exe
	rm -f *~
	rm -f \#*\#
	rm -f first second
