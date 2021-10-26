# TimeTracking CLI Client
## Description

## Usage
```bash
$ tracking --help
NAME:
   tracking - tracking CLI Client

USAGE:
   tracking [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   list, l     Show task list
   add, a      Add task
   start, s    Start Task
   finish, f   Finish Task
   current, c  Show current tracking
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Install

To install, use `go get`:

```bash
$ go get github.com/netooo/tracking
```

## Config
Set config in `$HOME/.config/tracking/config.json`

It has following parameters:
```
{
  "spread_sheet_id": "xxxxxxxxxxxxxxxxxxxxxxx", # favorite google spread sheet id
  "origin_date" : "2021-06-27",                 # first date in your sheet
  "origin_row" : "D"                            # row for the date
}
```

Set config in `$HOME/.config/tracking/secret.json` 

Issue a google spreadsheet config and paste it:
```
{
  "type": "service_account",
  "project_id": "{your project id}",
  "private_key_id": "xxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "private_key": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  <snip> 
}
```