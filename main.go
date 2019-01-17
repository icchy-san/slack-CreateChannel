package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/nlopes/slack"
)

// Env ... Variable for environment loading
var Env envConfig

func init() {
	loadEnvironment()
}

func main() {
	var fp *os.File
	var err error
	var faildChannels []string
	var randomTime int
	rand.Seed(time.Now().UnixNano())

	client := slack.New(Env.AccessToken)

	if len(os.Args) < 2 {
		os.Exit(1)
	} else {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	}

	reader := csv.NewReader(fp)
	reader.Comma = ','
	reader.LazyQuotes = true
	for {
		randomTime = rand.Intn(2)
		delay := time.Duration((randomTime + 1)) * time.Second
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		channelName := record[0]

		channel, err := client.CreateChannel(channelName)
		if err != nil {
			faildChannels = append(faildChannels, "#"+channelName)
			continue
		}
		log.Printf("CID: %#v\n", channel.ID)

		// 数秒くらい待ってあげましょうよ、という気持ちの現れ
		time.Sleep(delay)
	}

	params := setResultMessageParameters()

	faildChannelsText := strings.Join(faildChannels, ", ")
	fmt.Println(faildChannelsText)
	client.PostMessage(Env.AdminChannelID, faildChannelsText, params)
}

// setResultMessageParameters ... 実行結果ごのメッセージ設定関数
func setResultMessageParameters() slack.PostMessageParameters {
	attachment := slack.Attachment{
		Text: "プロセスが終了しました。チャネル名が表示された場合は、表示さているチャンネルが作成できませんでした。\nすでに存在していないか確認して下さい。",
	}
	return slack.PostMessageParameters{
		Attachments: []slack.Attachment{attachment},
	}
}
