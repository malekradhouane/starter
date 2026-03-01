package interfaces

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	// Primary key
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time  `gorm:"not null;default:now()" json:"created_at"`
	UpdatedAt time.Time  `gorm:"not null;default:now()" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`

	// Authentication
	Email        string     `gorm:"type:varchar(255);uniqueIndex:idx_users_email_lower,expression:LOWER(email);not null" json:"email"`
	PasswordHash string     `gorm:"type:varchar(255)" json:"-"`
	Provider     string     `gorm:"type:varchar(32);not null;default:'email'" json:"provider"`
	ProviderID   *string    `gorm:"type:varchar(255)" json:"provider_id"`
	LastLoginAt  *time.Time `json:"last_login_at"`
	Role         string     `gorm:"type:varchar(50);not null;default:'user'" json:"role"`

	// User profile
	Username    string  `gorm:"type:varchar(50);uniqueIndex:idx_users_username_lower,expression:LOWER(username);not null" json:"username"`
	FirstName   string  `gorm:"type:varchar(100)" json:"first_name"`
	LastName    string  `gorm:"type:varchar(100)" json:"last_name"`
	AvatarURL   string  `gorm:"type:text" json:"avatar_url"`
	PhoneNumber string  `gorm:"type:varchar(50)" json:"phone_number"`
	DateOfBirth *string `gorm:"type:date" json:"date_of_birth"` // Using string pointer for optional date
	Gender      string  `gorm:"type:varchar(20)" json:"gender"`
	Locale      string  `gorm:"type:varchar(10)" json:"locale"`

	// Account status
	EmailVerified bool `gorm:"default:false" json:"email_verified"`
	PhoneVerified bool `gorm:"default:false" json:"phone_verified"`
	IsActive      bool `gorm:"default:true;index:idx_users_is_active" json:"is_active"`
	IsSuperuser   bool `gorm:"default:false" json:"is_superuser"`

	// Security
	MFAEnabled          bool       `gorm:"default:false" json:"mfa_enabled"`
	MFASecret           string     `gorm:"type:varchar(100)" json:"-"`
	LastPasswordChange  *time.Time `json:"last_password_change"`
	FailedLoginAttempts int        `gorm:"default:0" json:"-"`
	LockedUntil         *time.Time `json:"-"`

	// Metadata
	Metadata *map[string]interface{} `gorm:"type:jsonb" json:"metadata,omitempty"`
}

// BeforeCreate is a hook that runs before creating a user
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return
}

// ValidationToken
type ValidationToken struct {
	UserID    uuid.UUID `gorm:"type:uuid;not null" json:"userID"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	Token     string    `gorm:"size:120;not null;primaryKey" json:"token"`
	TokenType string    `gorm:"size:20;not null;default:'activation'" json:"token_type"`
	ExpiredAt time.Time `gorm:"not null" json:"expired_at"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

// TableName specifies the table name for the ValidationToken model
func (ValidationToken) TableName() string {
	return "validation_tokens"
}
