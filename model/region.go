package model

// Region 区域
type Region struct {
	IDModel
	Name string `gorm:"type:varchar(255); not null default '';comment:名称"`
	Icon string `gorm:"type:varchar(255); not null default '';comment:横版图标"`
	Visible
	DateModel
	DeletedModel
}
