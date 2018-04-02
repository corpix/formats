package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/corpix/formats"
	"github.com/corpix/formats/compatibility"
	"github.com/urfave/cli"
)

func Action(ctx *cli.Context) error {
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

	switch toFormatName {
	case formats.JSON:
		v = compatibility.JSON(v)
	}

	if to == nil {
		fmt.Fprintf(
			os.Stdout,
			"%s",
			v,
		)
	} else {
		buf, err = to.Marshal(v)
		if err != nil {
			return err
		}
	}

	_, err = os.Stdout.Write(buf)
	if err != nil {
		return err
	}

	if !ctx.Bool("n") {
		_, err = os.Stdout.Write([]byte{'\n'})
	}

	return err
}

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "from",
			Value: "",
			Usage: "Decode `from` format",
		},
		cli.StringFlag{
			Name:  "to",
			Value: "",
			Usage: "Encode `to` format",
		},
		cli.BoolFlag{
			Name:  "n",
			Usage: "Do not print line break",
		},
	}

	app.Action = Action

	app.RunAndExitOnError()
}
