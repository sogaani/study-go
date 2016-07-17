PROG :=	server1.exe
SUMMARY := "パラメータによって様々なリサージュ曲線を表示するserver"

.PHONY: all run clean

.SUFFIXES: .go .exe
all: $(PROG)
run: $(PROG)
	@echo "========サマリー========"
	@echo $(SUMMARY)
	@echo "========実行開始========"
	./$< &
	curl http://localhost:8000/?cycle=20 > cycle20.gif
	curl http://localhost:8000/?size=200 > size200.gif
	pkill $<
.go.exe:
	go build -o $@ $<
clean:
	rm -f *.exe
	rm -f *~
	rm -f \#*\#
