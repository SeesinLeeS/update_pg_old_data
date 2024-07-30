package handler

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
	"update_pg_old_data/db"
)

func ChangeSpaceDefaultAvatar() {
	var (
		count int
		err   error
	)
	tmpData := make([]map[string]interface{}, 0)
	err = db.InstanceDB().Transaction(func(tx *gorm.DB) error {
		tmpAvatarId := ""
		err = tx.Table("space").Select("avatar_id,id").Where("avatar_id is not null").Scan(&tmpData).Error
		if err != nil {
			return err
		}
		for _, v := range tmpData {
			if strings.HasPrefix(v["avatar_id"].(string), "https://cdn.openmind.cn/") {
				tmpAvatarId = v["avatar_id"].(string)
				tmpAvatarId = strings.Trim(tmpAvatarId, "https://cdn.openmind.cn/")
				err = tx.Table("space").Select("avatar_id").Where("id = ?", v["id"].(int64)).Update("avatar_id", tmpAvatarId).Error
				if err != nil {
					break
				}
				count++
			}
		}
		return err
	})
	fmt.Printf("change count %d\n", count)
}
