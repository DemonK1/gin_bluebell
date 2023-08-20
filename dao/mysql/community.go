package mysql

import (
	"database/sql"
	"gin_bluebell/models"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := `select community_id,community_name from community`
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no db")
			err = nil
		}
	}
	return
}
