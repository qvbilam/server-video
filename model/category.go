package model

// Category 分类
type Category struct {
	IDModel
	ParentId int64  `gorm:"type:int not null default 0;comment:父级id;index:idx_parent_id"`
	Name     string `gorm:"type:varchar(255); not null default '';comment:名称"`
	Icon     string `gorm:"type:varchar(255); not null default '';comment:横版图标"`
	Level    int64  `gorm:"type:int(11);not null;default:0;comment:分类等级"`
	Visible
	DateModel
	DeletedModel
}
