# Personal Golang CLI

Built to handle Networking Tasks along with Opening applications from the command
line

## Stack

To build this project I used the urface/cli package. The CLI commands where then
placed in an array with a function as action. Using net/http package for
networking.

## Description

NAME: My-Cli - General Commands to enhance the terminal

USAGE: my-cli.exe [global options] command [command options] [arguments...]

COMMANDS:

-   new clone a new template django, go, next
-   ns Looks up Nameservers
-   ip Looks up the specified ip address
-   cname Looks up CNAME for a given host
-   mx Looks up MX records for a host
-   host Looks up a specified hosts addresses
-   browser Opens Chrome to the url if specified
-   spotify Opens Spotify
-   help, h Shows a list of commands or help for one command

# Setting to environment variable

Open windows start and search advanced windows settings. Open environment
Variables Open Path in System Variables and click Edit Set the Directory the exe
file is in You can rename the exe fle to change the cmd prefix
