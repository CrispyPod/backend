package helpers

import (
	"crispypod.com/crispypod-backend/graph/model"
	"gorm.io/gorm"
)

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetPageInfo(pageIndex int, pageSize int, total int) model.PageInfo {
	rt := model.PageInfo{
		HasNextPage:     true,
		HasPreviousPage: true,
	}
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageIndex == 0 || pageIndex == 1 {
		rt.HasPreviousPage = false
		pageIndex = 1
	}

	curEnd := pageIndex * pageSize
	if curEnd >= total {
		rt.HasNextPage = false
	}

	return rt
}
