package service

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type Cell struct {
	ID         int           `json:"id"`
	Img        string        `json:"img"`
	Text       string        `json:"text"`
	Permission int           `json:"permission" gorm:"column:premission"`
	Cate       int           `json:"cate"`
	FromID     string        `json:"fromID" gorm:"column:from_id"`
	FromURL    string        `json:"fromURL" gorm:"column:from_url"`
	CreatedAt  time.Time     `json:"createdAt" gorm:"column:createdat"`
	CreatedBy  sql.NullInt64 `json:"createdBy" gorm:"column:createdby"`
	UpdatedAt  time.Time     `json:"updatedAt" gorm:"column:updatedat"`
	Content    string        `json:"content"`
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
		Select("DISTINCT(from_id) as uid").
		Find(&uids).
		Limit(paginationSize).
		Offset(page * paginationSize).
		Error

	if gorm.IsRecordNotFoundError(err) {
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
	return db.Table("cells").Create(c).Error
}
