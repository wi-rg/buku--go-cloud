BUILD = build
BOOKNAME = pemrograman-go
TITLE = title.txt
METADATA = metadata.xml
CHAPTERS = awal.md bab-01.md bab-02.md bab-03.md bab-04.md bab-05.md bab-06.md bab-07.md bab-08.md bab-09.md bab-10.md bab-11.md bab-12.md bab-13.md bab-14.md bab-15.md bab-16.md bab-17.md bab-18.md bab-19.md bab-20.md bab-21.md bab-22.md bab-23.md bab-24.md sumberdaya.md daftar-pustaka.md

TOC = --toc --toc-depth=2
COVER_IMAGE = images/cover.jpg
LATEX_CLASS = report

all: book

book: epub html pdf

clean:
	rm -r $(BUILD)

epub: $(BUILD)/epub/$(BOOKNAME).epub

html: $(BUILD)/html/$(BOOKNAME).html

pdf: $(BUILD)/pdf/$(BOOKNAME).pdf

$(BUILD)/epub/$(BOOKNAME).epub: $(TITLE) $(CHAPTERS)
	mkdir -p $(BUILD)/epub
	pandoc $(TOC) -S --epub-metadata=$(METADATA) --epub-cover-image=$(COVER_IMAGE) -o $@ $^

$(BUILD)/html/$(BOOKNAME).html: $(CHAPTERS)
	mkdir -p $(BUILD)/html
	pandoc $(TOC) --standalone --to=html5 -o $@ $^

$(BUILD)/pdf/$(BOOKNAME).pdf: $(TITLE) $(CHAPTERS)
	mkdir -p $(BUILD)/pdf
	pandoc $(TOC) --latex-engine=xelatex -V documentclass=$(LATEX_CLASS) -o $@ $^

.PHONY: all book clean epub html pdf
