package model

// Style 风格
type Style struct {
	IDModel
	CategoryId int64  `gorm:"type:int not null default 0;comment:分类id;index:idx_category_id"`
	Name       string `gorm:"type:varchar(255); not null default '';comment:名称"`
	Icon       string `gorm:"type:varchar(255); not null default '';comment:横版图标"`
	Visible
	DateModel
	DeletedModel
}

// StyleVideo 风格视频关联
type StyleVideo struct {
	IDModel
	StyleId int64 `gorm:"type:int not null default 0;comment:风格id;index:idx_style_video"`
	VideoId int64 `gorm:"type:int not null default 0;comment:视频id;index:idx_style_video"`
	DateModel
}
