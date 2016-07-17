PROG :=	fetchall2.exe
SUMMARY := "alexaの上位500件のサイトへアクセスするfetchall"
INPUT := inputs.csv

.PHONY: all run clean

.SUFFIXES: .go .exe
all: $(PROG)
run: $(PROG) $(INPUT)
	@echo "========サマリー========"
	@echo $(SUMMARY)
	@echo "========実行開始========"
	cut -d ',' -f 2 $(INPUT) | xargs ./$<
$(INPUT): getAlexaTopSites.py
	./$< > $@
.go.exe:
	go build -o $@ $<
clean:
	rm -f *.exe
	rm -f *~
	rm -f \#*\#
