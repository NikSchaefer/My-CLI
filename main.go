package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "My-Cli",
		Usage: "General Commmands to enhance the terminal",
	}

	app.Commands = []*cli.Command{
		{
			Name:  "new",
			Usage: "clone a new template, django, go, next",
			Action: func(c *cli.Context) error {
				arg := c.Args().First()
				switch arg {
				case "next":
					cloneRepo("NikSchaefer/nextjs-boilerplate")
					fmt.Println("Sucessfully Created")
				case "django":
					cloneRepo("NikSchaefer/Django-backend")
					fmt.Println("Sucessfully Created")
				case "go":
					cloneRepo("NikSchaefer/go-fiber")
					fmt.Println("Sucessfully Created")
				default:
					fmt.Println("Specify a template, (django, go, or next)")
				}
				return nil
			},
		},
		{
			Name:  "ns",
			Usage: "Looks up Nameservers",
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.Args().First())
				must(err)
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the specified ip address",
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.Args().First())
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
			Usage: "Looks up CNAME for a given host",
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.Args().First())
				must(err)
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up MX records for a host",
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				must(err)
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i])
				}
				return nil
			},
		},
		{
			Name:  "host",
			Usage: "Looks up a specified hosts addresses",
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
			Usage: "Opens Chrome to the url if specified",
			Action: func(c *cli.Context) error {
				url := c.Args().First()
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
			Name:  "discord",
			Usage: "Opens Discord",
			Action: func(c *cli.Context) error {
				path := fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Discord\\Update.exe", os.Getenv("USERPROFILE"))
				exec.Command(path, "--processStart", "Discord.exe").Output()
				return nil
			},
		},
		{
			Name:  "spotify",
			Usage: "Opens Spoitfy",
			Action: func(c *cli.Context) error {
				path := fmt.Sprintf("C:\\Users\\%s\\AppData\\Roaming\\Spotify\\Spotify.exe", os.Getenv("USERPROFILE"))
				_, err := exec.Command(path).Output()
				must(err)
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	must(err)
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
func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

func cloneRepo(repo string) error {
	url := fmt.Sprintf("https://github.com/%s", repo)
	name := strings.Split(repo, "/")[1]
	_, err := exec.Command("git", "clone", url).Output()
	if err != nil {
		log.Fatal("Folder already exists")
	}
	git := fmt.Sprintf("./%s/.git", name)
	err = RemoveContents(git)
	must(err)
	err = os.Remove(git)
	must(err)
	fmt.Println("Sucessfully Created Boilerplate at")
	wd, _ := os.Getwd()
	fmt.Printf("%s\nextjs-boilerplate", wd)

	return nil
}
