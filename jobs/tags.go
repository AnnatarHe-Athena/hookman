package jobs

import (
	"encoding/json"
	"io/ioutil"

	"github.com/AnnatarHe-Athena/hookman/service"
	"github.com/sirupsen/logrus"
)

// 三万条数据
// 实际上： 17454
const startAt = 26955237
const endAt = 26985237

func TagAll() error {
	latestID := 1 << 30
	for latestID > startAt {
		logrus.Println("LastID: ", latestID)
		cells, err := service.WalkCells(latestID, 1000)
		if err != nil {
			return err
		}
		latestID = cells[len(cells)-1].ID
		tags := make([]service.TagCell, 0)

		for _, c := range cells {
			ts := service.AnalysisCell(c)
			for _, t := range ts {
				tags = append(tags, t)
			}
		}
		// logrus.Println(len(tags), tags)

		if err := service.SaveCellTags(tags); err != nil {
			return err
		}
	}

	return nil
}

type tempKey struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
	Tags  []int  `json:"tags"`
}

func DumpTrainData() error {
	latestID := 1 << 30
	temps := make([]tempKey, 0)
	for latestID > startAt-(1<<16) {
		logrus.Println("LastID: ", latestID)
		cells, err := service.WalkCells(latestID, 1000)
		if err != nil {
			return err
		}
		latestID = cells[len(cells)-1].ID
		for _, c := range cells {
			ts := service.AnalysisCell(c)

			tidList := make([]int, 0)
			for _, t := range ts {
				tidList = append(tidList, t.TagID)
			}

			temps = append(temps, tempKey{
				ID:    c.ID,
				Image: c.Img,
				Tags:  tidList,
			})
		}
	}

	b, _ := json.Marshal(temps)
	ioutil.WriteFile("./dump.json", b, 0777)
	return nil

}
