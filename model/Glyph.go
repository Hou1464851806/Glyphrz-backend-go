package model

type Glyph struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	FontID      uint   `json:"font_id"`                       // 所属字体ID
	Time        uint   `json:"time"`                          // 字形最新时间，用于云同步
	Unicode     string `json:"unicode"`                       // 字形unicode码
	Svg         string `json:"svg" gorm:"type:varchar(8192)"` // 字形描述
	XMin        uint   `json:"x_min"`
	XMax        uint   `json:"x_max"`
	YMin        uint   `json:"y_min"`
	YMax        uint   `json:"y_max"`
	MarginRight uint   `json:"margin_right"`
	IsGBK       uint   `json:"is_gbk"` // 统计GBKCount
}
