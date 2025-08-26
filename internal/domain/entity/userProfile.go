package entity

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// UserProfile represents additional profile information for a user.
type UserProfile struct {
	BaseEntityUUID

	// information User (one-to-one)
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null;uniqueIndex"`

	// Extended information User
	AvatarURL   *string        `json:"avatar_url,omitempty" gorm:"type:varchar(255)"`
	Bio         *string        `json:"bio,omitempty" gorm:"type:text"`
	Preferences datatypes.JSON `json:"preferences,omitempty"`
	Timezone    *string        `json:"timezone,omitempty" gorm:"type:varchar(50)"`

	// Relations:
	// One-to-One với User
	User User `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;<-:false"`
}
