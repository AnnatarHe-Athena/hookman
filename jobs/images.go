package jobs

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func downloadAndSave(baseFolder string, obj tempKey) error {
	imageURL := fmt.Sprintf("https://wx2.sinaimg.cn/orj360/%s.jpg", obj.Image)

	resp, err := http.DefaultClient.Get(imageURL)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	out, err := os.Create(baseFolder + strconv.Itoa(obj.ID) + ".jpg")
	if err != nil {
		return err
	}

	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}

	return nil
}

func DownloadImages(p string) error {
	temps := make([]tempKey, 0)

	bf, err := ioutil.ReadFile(p + "/dump.json")

	if err != nil {
		return err
	}

	if err := json.Unmarshal(bf, &temps); err != nil {
		return err
	}

	if err := os.Mkdir(p+"/images/", 0777); err != nil {
		return err
	}

	failed := make([]tempKey, 0)

	for i, k := range temps {
		if strings.HasPrefix(k.Image, "http") {
			failed = append(failed, k)
			continue
		}

		if err := downloadAndSave(p+"/images/", k); err != nil {
			logrus.Errorln(err)
			failed = append(failed, k)
		}

		logrus.Println("%d/%d", i, len(temps))
	}

	logrus.Println("pass")

	return nil

}
