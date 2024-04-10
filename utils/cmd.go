package utils

import (
	"fmt"

	"github.com/hellflame/argparse"
)

type args struct {
	Host      string
	Port      int
	NoBrowser bool
	AfterHelp bool
}

func (a *args) GetAddress() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

func ParseArgs() *args {
	parser := argparse.NewParser("", "start arithmetic server", &argparse.ParserConfig{
		WithHint: true, WithColor: true, DisableDefaultShowHelp: true,
	})

	host := parser.String("", "host", &argparse.Option{Default: "127.0.0.1", Help: "host address"})
	port := parser.Int("p", "port", &argparse.Option{Default: "8080", Help: "listen port"})

	noBrowser := parser.Flag("", "no-browser", &argparse.Option{Help: "don't open browser"})

	if e := parser.Parse(nil); e != nil {
		fmt.Println(e.Error())
		return nil
	}
	return &args{
		Host: *host,
		Port: *port,

		NoBrowser: *noBrowser,
	}
}
