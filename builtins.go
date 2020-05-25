package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

var (
	Aliased = make(map[string]string)
)

func Alias(keyword string, macro string) {
	Aliased[keyword] = macro
}

func ChangeDir(location string) error {
	if strings.HasPrefix(location, "~") {
		location = strings.Replace(location, "~", os.Getenv("HOMEPATH"), 1)
	} else if strings.HasPrefix(location, "..") {
		location = strings.Replace(location, "..", WorkingDir()[0:strings.LastIndex(WorkingDir(), "/")], 1)
	} else if strings.HasPrefix(location, "@") {
		location = strings.Replace(location, "@", os.Getenv("TURT_PASTDIR"), 1)
	}

	location = strings.ReplaceAll(location, "\\", "/")
	location = strings.ReplaceAll(location, "//", "/")

	_ = os.Setenv("TURT_PASTDIR", WorkingDir())
	_ = os.Setenv("TURT_CURDIR", location)
	return os.Chdir(location)
}

func formatName(name string, isDir bool) string {
	if len(name) > 50 {
		name = name[:47] + "..."
	}

	if isDir {
		name = name + "/"
	}

	return name
}

func ListDir() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Size", "Last Modified"})

	files, err := ioutil.ReadDir(WorkingDir())
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		t.AppendRow([]interface{}{formatName(f.Name(), f.IsDir()), f.Size(), f.ModTime()})
	}

	t.SetStyle(table.Style{
		Name: "myNewStyle",
		Box: table.BoxStyle{
			BottomLeft:       " ",
			BottomRight:      " ",
			BottomSeparator:  " ",
			Left:             " ",
			LeftSeparator:    " ",
			MiddleHorizontal: " ",
			MiddleSeparator:  " ",
			MiddleVertical:   " ",
			PaddingLeft:      " ",
			PaddingRight:     " ",
			Right:            " ",
			RightSeparator:   " ",
			TopLeft:          " ",
			TopRight:         " ",
			TopSeparator:     " ",
			UnfinishedRow:    " ",
		},
		Format: table.FormatOptions{
			Footer: text.FormatUpper,
			Header: text.FormatUpper,
			Row:    text.FormatDefault,
		},
		Options: table.Options{
			DrawBorder:      true,
			SeparateColumns: false,
			SeparateFooter:  false,
			SeparateHeader:  false,
			SeparateRows:    false,
		},
	})
	t.Render()
}

func WorkingDir() string {
	return os.Getenv("TURT_CURDIR")
}

func Unalias(keyword string) {
	delete(Aliased, keyword)
}
