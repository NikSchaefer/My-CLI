package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/urfave/cli"
)

func main() {

	userPath := "schaefer"

	app := cli.NewApp()
	app.Name = "My Personal CLI"
	app.Usage = "Boosting Productivity and opening apps from the Command Line"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "nikschaefer.tech",
		},
		&cli.StringFlag{
			Name:  "url",
			Value: "",
		},
		&cli.StringFlag{
			Name:  "addr",
			Value: "",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up Nameservers",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up Nameservers",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Looks up CNAME for a host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					return err
				}
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "ping",
			Usage: "Ping Address",
			Flags: myFlags,
			Action: func(c *cli.Context) error {

				exec.Command("ping", c.String("addr"))
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up MX records for a host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i])
				}
				return nil
			},
		},
		{
			Name:  "host",
			Usage: "Looks up a hosts addresses",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				hosts, err := net.LookupHost(c.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(hosts); i++ {
					fmt.Println(hosts)
				}
				return nil
			},
		},
		{
			Name:  "browser",
			Usage: "Opens Chrome",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				url := c.String("url")
				if url == "" {
					openbrowser(" ")
				}
				if strings.Contains(url, "https://") {
					openbrowser(url)
				} else {
					openbrowser("https://" + url)
				}
				return nil
			},
		},
		{
			Name:  "dis",
			Usage: "Opens Discord",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				path := fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Discord\\Update.exe", userPath)
				exec.Command(path, "--processStart", "Discord.exe").Output()
				return nil
			},
		},
		{
			Name:  "spotify",
			Usage: "Opens Spoitfy",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				path := fmt.Sprintf("C:\\Users\\%s\\AppData\\Roaming\\Spotify\\Spotify.exe", userPath)
				exec.Command(path).Output()
				return nil
			},
		},
	}
	err := app.Run((os.Args))
	if err != nil {
		log.Fatal(err)
	}
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
