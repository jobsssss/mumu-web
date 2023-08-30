package bookHotWord

import "mumu/pkg/model"

func All() (wanted []BookHotWord) {
	model.DB.Where("deleted_at", 0).Find(&wanted)
	return
}
