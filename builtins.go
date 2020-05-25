package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

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
	location = strings.ReplaceAll(location, "\\", "/")
	location = strings.ReplaceAll(location, "//", "/")

	if strings.HasPrefix(location, "~") {
		location = strings.Replace(location, "~", os.Getenv("HOMEPATH")+"/", 1)
	} else if strings.HasPrefix(location, "..") {
		location = strings.Replace(location, "..", WorkingDir()[0:strings.LastIndex(WorkingDir(), "/")], 1)
	} else if strings.HasPrefix(location, "@") {
		location = strings.Replace(location, "@", os.Getenv("TURT_PASTDIR")+"/", 1)
	} else if !strings.HasPrefix(location, "/") {
		location = WorkingDir() + "/" + location
	}

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

// https://programming.guide/go/formatting-byte-size-to-human-readable-format.html
func formatMemory(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%d %cB", b/div, "KMGTPE"[exp])
}

func formatTime(time time.Time) string {
	return time.Format("01/02/2006 15:04")
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
		t.AppendRow([]interface{}{formatName(f.Name(), f.IsDir()), formatMemory(f.Size()), formatTime(f.ModTime())})
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
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 2, Align: text.AlignRight},
	})
	t.Render()
}

func WorkingDir() string {
	return os.Getenv("TURT_CURDIR")
}

func Unalias(keyword string) {
	delete(Aliased, keyword)
}
