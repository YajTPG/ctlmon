# ctlmom - Systemctl Services Monitor

## Getting Started

- Download latest binary from [releases](https://github.com/YajTPG/ctlmon/releases) page,  
- Run the binary once to create config file  
- Edit config file located in `/etc/ctlmon/config.yaml` to your desired choice.

## Sample config file in [config.example.yml](/config.example.yml)

- `nodename` => For reference in the discord webhook, default is hostname
- `roleid` => Discord RoleID to be mentioned, to notify people about the downtime of the service
- `services` => List of services to check
- `version` => Current version of the config file, __do not touch__
- `webhookenabled` => Acts as a killswitch to turn off discord notifications
- `webhookurl` => URL of the webhook to send the POST request to

## Adding Cron

Open your cron file via `sudo crontab -e` and add in the following:

```c
* * * * * /PATH/TO/ctlmon // Runs ctlmon commmand every minute
```
