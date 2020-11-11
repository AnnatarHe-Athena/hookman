package jobs

import (
	"github.com/AnnatarHe-Athena/hookman/service"
	"github.com/sirupsen/logrus"
)

func TagAll() error {
	latestID := 1 << 30
	for latestID > 0 {
		cells, err := service.WalkCells(latestID, 1000)

		defer func() {
			logrus.Println("LastID: ", latestID)
			latestID -= 1000
		}()

		if err != nil {
			return err
		}

		tags := make([]service.TagCell, 0)

		for _, c := range cells {

			ts := service.AnalysisCell(c)
			for _, t := range ts {
				tags = append(tags, t)
			}
		}

		logrus.Println(len(tags), tags)

		// if err := service.SaveCellTags(tags); err != nil {
		// 	return err
		// }
	}

	return nil
}
