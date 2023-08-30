package category

import "mumu/pkg/model"

func All(channelId int)(category []BookCategory)  {
	builder := model.DB.Order("id DESC")
	if channelId > 0 {
		builder.Where("channel_id = ?", channelId)
	}
	builder.Where("deleted_at", 0).Find(&category)
	return
}
