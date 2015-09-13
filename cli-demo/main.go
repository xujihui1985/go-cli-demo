package main

import (
	"github.com/codegangsta/cli"
	"github.com/ttacon/chalk"
	"os"
	"bytes"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello"
	app.Usage = "testtesttest"
	app.Flags = []cli.Flag {
		cli.StringFlag {
			Name: "lang, l",
			Value: "english",
			Usage: chalk.Magenta.Color("language for xxxxx"),
		},
		cli.BoolFlag {
			Name: "verbose, V",
			Usage: "show verbose message",
		},

	}
	app.Version = "0.1.0"
	app.Action = func (c *cli.Context) {
		name := "value"
		if len(c.Args()) == 0 {
			cli.ShowAppHelp(c);
		}
		if len(c.Args()) > 0 {
			name = c.Args()[0]
			//c.App.Writer.Write([]byte("hello world\n"))
		}
	    //c.App.Writer.Write([]byte(c.String("lang")+"\n"))
		if c.String("lang") == "chinese" {
			fmt.Printf("nihao %s\n", name)
		}
	}
	app.Commands = []cli.Command {
		cli.Command {
			Name: "add",
			Aliases: []string{"a"},
			Usage: "Add lang to",
			Description: "xxxxxxxxxxxxxxxxxxxxxx",
			Before: func (c *cli.Context) error {
				c.App.Writer.Write([]byte("before add\n"))
				return nil
			},
			After: func (c *cli.Context) error {
				c.App.Writer.Write([]byte("after add\n"))
				return nil
			},
			Action: func (c *cli.Context) {
				c.App.Writer.Write([]byte(c.Args().First()+"\n"))
			},
		},
		cli.Command {
			Name: "issue",
			Usage: "issue management",
			Description: "issue management Description",
			Subcommands: []cli.Command {
				cli.Command {
					Name: "add",
					Usage: "add issue",
					Description: "add issue to gitlab",
					Action: func(c *cli.Context) {
						buffer := bytes.NewBufferString("")
						buffer.WriteString("issue added\n")
						c.App.Writer.Write(buffer.Bytes())
					},
				},
				cli.Command {
					Name: "list",
					Usage: "list issue",
					Description: "list issue from gitlab",
					Action: func(c *cli.Context) {
						buffer := bytes.NewBufferString("issue listed \n")
						c.App.Writer.Write(buffer.Bytes())
					},
				},
			},
			Before: func (c *cli.Context) error {
				c.App.Writer.Write([]byte("before add\n"))
				return nil
			},
			After: func (c *cli.Context) error {
				c.App.Writer.Write([]byte("after add\n"))
				return nil
			},
			Action: func (c *cli.Context) {
				c.App.Writer.Write([]byte(c.Args().First()+"\n"))
			},
		},
	}
	app.Run(os.Args)
}


