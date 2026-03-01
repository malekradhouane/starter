package interfaces

import (
	"time"

	"github.com/google/uuid"
)

// Merchant maps to the `merchants` table
type Merchant struct {
	ID         uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CreatedAt  time.Time       `gorm:"type:timestamptz;not null;default:now()" json:"created_at"`
	UpdatedAt  time.Time       `gorm:"type:timestamptz;not null;default:now()" json:"updated_at"`
	DeletedAt  *time.Time      `gorm:"index" json:"deleted_at,omitempty"`
	Name       string          `gorm:"type:text;not null" json:"name"`
	Slug       string          `gorm:"type:text;uniqueIndex;not null" json:"slug"`
	WebsiteURL *string         `gorm:"type:text" json:"website_url,omitempty"`
	LogoURL    *string         `gorm:"type:text" json:"logo_url,omitempty"`
	IsActive   bool            `gorm:"not null;default:true;index" json:"is_active"`
	Metadata   *map[string]any `gorm:"type:jsonb" json:"metadata,omitempty"`
}
