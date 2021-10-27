# TimeTracking CLI Client
## Description
This program will let you use the time tracking in CLI.

And, the tracking data will be automatically entered into the Google Spread Sheet.

## Usage
```
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
   start, s    Start task
   finish, f   Finish task
   current, c  Show current tracking
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Demo (with [fzf](https://github.com/junegunn/fzf))
### Add task
![Add_task](https://user-images.githubusercontent.com/46105888/138991428-e272d358-a340-46de-8692-19884e2ed80e.gif)

### Start task
![Start_task](https://user-images.githubusercontent.com/46105888/138991566-5950085a-05af-4639-bfba-0051c5e99ec4.gif)

### Finish task
![Finish_task](https://user-images.githubusercontent.com/46105888/138991571-b54c8900-c94a-42a0-8e72-ed2055f7637f.gif)

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