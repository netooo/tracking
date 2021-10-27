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
![Add task](https://user-images.githubusercontent.com/46105888/138989198-7935c1ca-330e-4a26-bdd0-39ccd5972b46.gif)

### Start task
![Start task](https://user-images.githubusercontent.com/46105888/138989560-c84e7d6a-d5f4-4cde-bf3a-e13722412932.gif)

### Finish task
![Finish task](https://user-images.githubusercontent.com/46105888/138989571-4f01322d-4dbb-4690-ac5b-69278c60f632.gif)

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