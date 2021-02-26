package service

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Cell struct {
	ID         int           `json:"id"`
	Img        string        `json:"img"`
	Text       string        `json:"text" gorm:"column:text;type:text"`
	Permission int           `json:"permission" gorm:"column:premission"`
	Cate       int           `json:"cate"`
	FromID     string        `json:"fromID" gorm:"column:from_id"`
	FromURL    string        `json:"fromURL" gorm:"column:from_url"`
	CreatedAt  time.Time     `json:"createdAt" gorm:"column:createdat"`
	CreatedBy  sql.NullInt64 `json:"createdBy" gorm:"column:createdby"`
	UpdatedAt  time.Time     `json:"updatedAt" gorm:"column:updatedat"`
	Content    string        `json:"content" gorm:"column:content;type:text"`
	Likes      int           `json:"likes"`
	Md5        string        `json:"md5"`
}

// 2020-04-12 接近 10w 条数据
const paginationSize = 200000

type withUid struct {
	UID string `gorm:"column:uid"`
}

func ListWeiboUsers(page int) (wbUserIDs []string, err error) {

	var uids []withUid
	// select DISTINCT(from_id) from cells where from_url like '%weibo.com%' and premission = 2
	err = db.
		Table("cells").
		Where("premission = ?", 2).
		Where("from_url like ?", "%weibo.com%").
		Not("from_id", blacklist).
		Select("DISTINCT(from_id) as uid").
		Find(&uids).
		Limit(paginationSize).
		Offset(page * paginationSize).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return wbUserIDs, nil
	}

	for _, u := range uids {
		wbUserIDs = append(wbUserIDs, u.UID)
	}

	return
}

func TempUpdateUserDomainToUid(domainName, uid string) error {
	return db.Table("cells").Where("from_id = ?", domainName).Update("from_id", uid).Error
}

func (c *Cell) Create() error {
	err := db.Table("cells").Create(c).Error

	if err != nil && strings.Contains(err.Error(), "invalid byte sequence for encoding") {
		logrus.Errorln("invalid", c)
	}
	return err
}
