package interfaces

import (
	"time"

	"github.com/google/uuid"
)

// Deal maps to the `deals` table
type Deal struct {
	ID              int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	AuthorID        uuid.UUID  `gorm:"type:uuid;not null;index" json:"author_id"`
	MerchantID      uuid.UUID  `gorm:"type:uuid;not null;index" json:"merchant_id"`
	Title           string     `gorm:"type:text;not null" json:"title"`
	Slug            string     `gorm:"type:text;uniqueIndex;not null" json:"slug"`
	DescriptionMD   *string    `gorm:"type:text" json:"description_md"`
	DealType        string     `gorm:"type:text;not null" json:"deal_type"`
	OriginalPrice   *float64   `gorm:"type:numeric(12,2)" json:"original_price"`
	Price           *float64   `gorm:"type:numeric(12,2)" json:"price"`
	Currency        string     `gorm:"type:char(3);not null;default:EUR" json:"currency"`
	DiscountPercent *int16     `json:"discount_percent"`
	StartAt         time.Time  `gorm:"type:timestamptz;not null;default:now()" json:"start_at"`
	EndAt           *time.Time `gorm:"type:timestamptz" json:"end_at"`
	ExpiresAt       *time.Time `gorm:"type:timestamptz" json:"expires_at"`
	IsLocal         bool       `gorm:"not null;default:false" json:"is_local"`
	LocationType    string     `gorm:"type:text;not null;default:'online'" json:"location_type"`
	Status          string     `gorm:"type:text;not null;default:'draft'" json:"status"`
	CreatedAt       time.Time  `gorm:"type:timestamptz;not null;default:now()" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"type:timestamptz;not null;default:now()" json:"updated_at"`
	PublishedAt     *time.Time `gorm:"type:timestamptz" json:"published_at"`
}
