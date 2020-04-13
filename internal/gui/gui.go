package gui

import (
	"context"

	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/text"
)

type Gui struct {
	quit func()
}

func (gui *Gui) keyHandler(k *terminalapi.Keyboard) {
	if k.Key == 'q' || k.Key == 'Q' {
		gui.quit()
	}
}

func (gui *Gui) Init() error {
	term, err := termbox.New()
	if err != nil {
		return err
	}
	defer term.Close()

	textWidget, err := text.New()
	if err != nil {
		return err
	}
	if err := textWidget.Write("Hello, World!"); err != nil {
		return err
	}

	mainContainer, err := container.New(
		term,
		container.SplitVertical(
			container.Left(
				container.PlaceWidget(textWidget),
				container.Border(linestyle.Light),
			),
			container.Right(),
			container.SplitPercent(30),
		),
	)
	if err != nil {
		panic(err)
	}

	ctx, quit := context.WithCancel(context.Background())
	gui.quit = quit // Why direct assignment doesn't work ??
	defer gui.quit()

	if err := termdash.Run(ctx, term, mainContainer, termdash.KeyboardSubscriber(gui.keyHandler)); err != nil {
		panic(err)
	}

	return nil
}

func NewGui() error {
	gui := Gui{}
	return gui.Init()
}
