package model

type URL struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	OriginalURL string `gorm:"type:text;not null"`
	ShortCode   string `gorm:"type:varchar(10);uniqueIndex;not null"`
	Clicks      uint   `gorm:"default:0"`
}
