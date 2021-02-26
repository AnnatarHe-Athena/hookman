package service

import (
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type Tag struct {
	ID          int
	Name        string
	Description string
	CreatedAt   time.Time `gorm:"column:createdat"`
	UpdatedAt   time.Time `gorm:"column:updatedat"`
}

type TagCell struct {
	ID     int
	CellID int `gorm:"column:cell_id"`
	TagID  int `gorm:"column:tag_id"`
}

func (TagCell) TableName() string {
	return "tags_girls"
}

var wordsTagNameMapping map[int][]string = map[int][]string{
	1:  []string{"可爱", "可可爱爱", "洛丽塔", "白丝"},
	2:  []string{"性感", "黑丝", "酒吧", "夜店", "HM", "BM"},
	3:  []string{"黑丝", "白丝", "丝袜"},
	4:  []string{"jk"},
	5:  []string{},
	6:  []string{},
	7:  []string{},
	8:  []string{},
	9:  []string{},
	10: []string{"肌肉", "健身", "腹肌", "💪"},
	11: []string{"男朋友", "帅哥"},
	12: []string{"脸", "睫毛", "眼睛", "👀", "👄", "嘴巴", "耳朵", "👂", "鼻子", "👃", "化妆", "眼影", "妆容", "发色", "刘海", "宿舍"},
	13: []string{"🐻", "胸", "胖"},
	14: []string{"臀"},
	15: []string{"腿", "🦵", "黑丝", "jk"},
	16: []string{"好吃的", "餐厅", "鸡腿", "午饭", "晚饭", "夜宵", "火锅", "烧烤", "🍲"},
	17: []string{"景色", "风景", "山", "海", "湾"},
	18: []string{},
}

var tagsMapping map[string]int = map[string]int{
	"kawai":    1,
	"sexy":     2,
	"silk":     3,
	"jk":       4,
	"weibo":    5,
	"zhihu":    6,
	"jike":     7,
	"red":      8,
	"ig":       9,
	"muscle":   10,
	"boy":      11,
	"face":     12,
	"chest":    13,
	"buttocks": 14,
	"leg":      15,
	"food":     16,
	"view":     17,
	"others":   18,
}

func _SetupTags() (map[string]int, error) {
	var tags []Tag
	result := map[string]int{}
	if err := db.Table("tags").Find(&tags).Error; err != nil {
		return result, err
	}

	for _, v := range tags {
		result[v.Name] = v.ID
	}

	return result, nil
}

func WalkCells(lastId, limit int) (result []Cell, err error) {
	err = db.Table("cells").Where("id < ?", lastId).Order("id DESC").Limit(limit).Find(&result).Error

	return result, err
}

func AnalysisCell(c Cell) []TagCell {
	connections := make([]TagCell, 0)
	for id, words := range wordsTagNameMapping {
		if len(words) == 0 {
			continue
		}

		for _, w := range words {
			if !strings.Contains(strings.ToLower(c.Text), strings.ToLower(w)) {
				continue
			}
			connections = append(connections, TagCell{
				CellID: c.ID,
				TagID:  id,
			})
		}
	}

	return connections
}

func SaveCellTags(tags []TagCell) error {
	logrus.Println(len(tags), tags)
	return db.Table("tags_girls").Create(&tags).Error
}
