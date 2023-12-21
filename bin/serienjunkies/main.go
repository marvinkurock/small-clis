package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Item struct {
	CreatedAt  string `json:"createdAt"`
	Encoding   string `json:"encoding"`
	Language   string `json:"language"`
	Name       string `json:"name"`
	Resolution string `json:"resolution"`

	Type    string
	Season  int
	Episode int
}

var filter string

func main() {
	var cmd = &cobra.Command{
		Use:   "sj",
		Short: "sj is a simple serienjunkies query cli",
		Run: func(cmd *cobra.Command, args []string) {
			println("cmd: ", filter)
			getData()
		},
	}
	cmd.Flags().StringVarP(&filter, "filter", "f", "", "filter for shows")
	cmd.Execute()
}

func getData() {
	url := "https://serienjunkies.org/api/releases/latest/0"
	resp, err := http.Get(url)
	if err != nil {
		println("error getting data")
		panic(err)
	}
	defer resp.Body.Close()

	var data []Item
	buffer, err := io.ReadAll(resp.Body)
	if err != nil {
		println("error reading body")
		panic(err)
	}
	err = json.Unmarshal(buffer, &data)
	if err != nil {
		println("error unmarshalling")
		panic(err)
	}
	var result []Item
	for _, i := range data {
		t, s, e := parseName(i.Name)
		i.Season = s
		i.Episode = e
		i.Type = t
		result = append(result, i)
	}
	for _, i := range result {
		prettyPrintItem(i)
	}
}

func parseName(name string) (string, int, int) {
	season := -1
	episode := -1
	t := "Movie"
	reg := regexp.MustCompile(`[Ss]{1}(\d{1,2})[Eex]{0,1}(\d{0,2})`)
	matches := reg.FindStringSubmatch(name)
	if len(matches) == 3 {
		season, _ = strconv.Atoi(matches[1])
		episode, _ = strconv.Atoi(matches[2])
		t = "Show"
	}
	return t, season, episode
}

func prettyPrintItem(i Item) {
	if i.Type == "Show" {
		// 31 = red
		// 32 = green
		// 33 = yellow
		// 34 = blue
		// 35 = magenta
		// 0 = reset
		resolutionFix := 5 - len(i.Resolution)
		padding := strings.Repeat(" ", resolutionFix)
		if i.Episode == 0 {
			fmt.Printf("[\033[35m%v\x1b[0m]  [\033[31m%v\x1b[0m]%v (Season \033[33m%v\x1b[0m)\t %v\n", "Season", i.Resolution, padding, i.Season, i.Name)
		} else {
			fmt.Printf("[\033[32m%v\x1b[0m] [\033[31m%v\x1b[0m]%v (S\033[34m%v\x1b[0m E\033[34m%v\x1b[0m)\t %v\n", "Episode", i.Resolution, padding, i.Season, i.Episode, i.Name)
		}
	} else {
		fmt.Printf("[%v] %v\n", i.Type, i.Name)
	}
}
