package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"

	"github.com/corpix/formats"
	"github.com/corpix/formats/compatibility"
)

var (
	Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "from",
			Value: "",
			Usage: "decode `from` format",
		},
		cli.StringFlag{
			Name:  "to",
			Value: "",
			Usage: "encode `to` format",
		},
		cli.BoolFlag{
			Name:  "n",
			Usage: "do not print line break",
		},
	}

	Commands = []cli.Command{
		cli.Command{
			Name:   "list",
			Usage:  "List supported formats",
			Action: ListAction,
		},
	}
)

func RootAction(ctx *cli.Context) error {
	var (
		v              = *new(interface{})
		fromFormatName = ctx.String("from")
		toFormatName   = ctx.String("to")
		from           formats.Format
		to             formats.Format
		buf            []byte
		err            error
	)

	if fromFormatName != "" {
		from, err = formats.New(fromFormatName)
		if err != nil {
			return err
		}
	}

	if toFormatName != "" {
		to, err = formats.New(toFormatName)
		if err != nil {
			return err
		}
	}

	buf, err = ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	if from == nil {
		v = formats.NewStringer(string(buf))
	} else {
		err = from.Unmarshal(buf, &v)
		if err != nil {
			return err
		}
	}

	if to == nil {
		fmt.Fprintf(os.Stdout, "%s", v)
	} else {
		switch toFormatName {
		case formats.JSON:
			v = compatibility.JSON(v)
		}

		buf, err = to.Marshal(v)
		if err != nil {
			return err
		}

		_, err = os.Stdout.Write(buf)
		if err != nil {
			return err
		}
	}

	if !ctx.Bool("n") {
		_, err = os.Stdout.Write([]byte{'\n'})
	}

	return err
}

func ListAction(ctx *cli.Context) error {
	for _, name := range formats.Names {
		fmt.Printf("%s - %s\n", name, formats.Descriptions[name])
	}
	return nil
}

func main() {
	app := cli.NewApp()

	app.Flags = Flags
	app.Commands = Commands
	app.Action = RootAction

	app.RunAndExitOnError()
}
