package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"strconv"
	"unsafe"
)

func main() {
	gtk.Init(&os.Args)

	dialog := gtk.Dialog()
	dialog.SetTitle("number input")

	vbox := dialog.GetVBox()

	label := gtk.Label("Numnber:")
	vbox.Add(label)

	input := gtk.Entry()
	input.SetEditable(true)
	vbox.Add(input)

	input.Connect("insert-text", func(ctx *glib.CallbackContext) {
		a := (*[2000]uint8)(unsafe.Pointer(ctx.Args(0)))
		p := (*int)(unsafe.Pointer(ctx.Args(2)))
		i := 0
		for a[i] != 0 {
			i++
		}
		s := string(a[0:i])
		if s == "." {
			if *p == 0 {
				input.StopEmission("insert-text")
			}
		} else {
			_, err := strconv.ParseFloat(s, 64)
			if err != nil {
				input.StopEmission("insert-text")
			}
		}
	})

	button := gtk.ButtonWithLabel("OK")
	button.Connect("clicked", func() {
		println(input.GetText())
		gtk.MainQuit()
	})
	vbox.Add(button)

	dialog.ShowAll()
	gtk.Main()
}
