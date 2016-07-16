PROG := lissajous2.exe
SUMMARY := "背景が黒　線の色が虹色のリサージュ図形を作成する"

.PHONY: all run clean

.SUFFIXES: .go .exe
all: $(PROG)
run: $(PROG)
	./$< > lissajous.gif
.go.exe:
	go build -o $@ $<
clean:
	rm -f *.exe
	rm -f *~
	rm -f \#*\#
