package entity

type BaseModel struct {
	Id        int   `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	CreatedAt int64 `gorm:"column:created_at" json:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt int64 `gorm:"column:deleted_at" json:"deleted_at"`
}
