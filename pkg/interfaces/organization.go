package interfaces

import (
	"time"

	"github.com/google/uuid"
)

// Organization represents a company or organization
type Organization struct {
	// Primary key
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null;default:now()" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`

	// Organization details
	Name        string  `gorm:"type:varchar(255);not null" json:"name"`
	DisplayName *string `gorm:"type:varchar(255)" json:"display_name"`
	Description *string `gorm:"type:text" json:"description"`
	Website     *string `gorm:"type:varchar(255)" json:"website"`
	LogoURL     *string `gorm:"type:text" json:"logo_url"`

	// Contact information
	Email   *string `gorm:"type:varchar(255)" json:"email"`
	Phone   *string `gorm:"type:varchar(50)" json:"phone"`
	Address *string `gorm:"type:text" json:"address"`
	City    *string `gorm:"type:varchar(100)" json:"city"`
	Country *string `gorm:"type:varchar(100)" json:"country"`

	// Organization settings
	IsActive   bool `gorm:"default:true" json:"is_active"`
	IsVerified bool `gorm:"default:false" json:"is_verified"`

	// Relations
	Users []User `gorm:"foreignKey:OrganizationID" json:"users,omitempty"`
}

// TableName specifies the table name for Organization
func (Organization) TableName() string {
	return "organizations"
}
