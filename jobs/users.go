package jobs

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/AnnatarHe-Athena/hookman/config"
	"github.com/AnnatarHe-Athena/hookman/service"
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
)

func FetchUser() {
}

func fetchUserFeed(page int, uid string) (response GetWeiboFeedResponse, err error) {
	ua := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 13_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/81.0.4044.62 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (iPad; CPU OS 13_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/81.0.4044.62 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (iPod; CPU iPhone OS 13_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/81.0.4044.62 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (Linux; Android 10; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.96 Mobile Safari/537.36",
		"Mozilla/5.0 (Linux; Android 10; SM-N960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.96 Mobile Safari/537.36",
	}

	url := fmt.Sprintf("https://m.weibo.cn/api/container/getIndex?count=%d&page=%d&containerid=107603%s", 20, page, uid)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", ua[rand.Intn(len(ua))])
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		// continue
		return
	}

	defer res.Body.Close()

	responseData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return
	}

	err = json.Unmarshal(responseData, &response)

	if err != nil {
		logrus.Println(url)
		// continue
		return
	}

	logrus.Println(response.Ok)
	if err != nil {
		return response, err
	}

	if response.Ok != 0 && response.Ok != 1 {
		return response, errors.New("no data")
	}

	if len(response.Data.Cards) < 1 {
		return response, errors.New("no data")
	}

	return
}

func UpdateUsersFor1week() error {
	userIDs, err := service.ListWeiboUsers(0)
	logrus.Println(userIDs)
	// return nil

	if err != nil {
		return err
	}

	for _, userID := range userIDs {
		page := 0
		for page > -1 {
			// 防止过多的请求
			time.Sleep(time.Millisecond * 100)
			f, err := fetchUserFeed(page, userID)
			if err != nil {
				panic(err)
			}

			for _, card := range f.Data.Cards {
				if card.CardType != 9 {
					continue
				}

				for _, pic := range card.Mblog.Pics {
					doc, err := goquery.NewDocumentFromReader(bytes.NewReader([]byte("<div>" + card.Mblog.Text + "</div>")))
					if err != nil {
						logrus.Errorln(err)
						continue
					}

					maxLen := len(doc.Text())

					if maxLen > 255 {
						maxLen = 255
					}
					content := doc.Text()[0:maxLen]

					cell := &service.Cell{
						Img:        pic.Pid,
						Cate:       177,
						Text:       content,
						Permission: 2,
						FromID:     userID,
						FromURL:    fmt.Sprintf("https://weibo.com/%s/%s", userID, card.Mblog.Bid),
						CreatedAt:  time.Now(),
						CreatedBy:  sql.NullInt64{Int64: int64(2), Valid: true},
						UpdatedAt:  time.Now(),
						Content:    content,
						Likes:      0,
						Md5:        sha256Encode(pic.Pid),
					}

					if err := cell.Create(); err != nil {
						if !strings.Contains(err.Error(), "duplicate key value violates") {
							logrus.Errorln(err)
						}
						continue
					}
				}
			}

			// 下一页
			page++
			// 只拉 5 页
			if page > config.MaxPage {
				page = -1
			}
		}
	}

	return nil
}
