PROG :=	dup1.exe
SUMMARY := "重複した行毎にその行が含まれていたファイル名をすべて表示するdup"

testfiles := test1 test2 test3 test4

.PHONY: all run clean

.SUFFIXES: .go .exe
all: $(PROG)
run: $(PROG) $(testfiles)
	@echo "========サマリー========"
	@echo $(SUMMARY)
	@echo "========実行開始========"
	./$< $(testfiles)
	@echo ""
$(testfiles):
	cat /dev/urandom | tr -dc 'A-Z' | fold -w 1 | head -n 20 > $@
.go.exe:
	go build -o $@ $<
clean:
	rm $(testfiles)
	rm -f *.exe
	rm -f *~
	rm -f \#*\#
