package entity

import (
	"time"

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

// NewUserProfileWithUserID creates a new UserProfile with the given userID and initializes timestamps.
func NewUserProfileWithUserID(userID uuid.UUID) *UserProfile {
	now := time.Now()
	return &UserProfile{
		BaseEntityUUID: BaseEntityUUID{
			ID: uuid.New(), // override default gen_random_uuid()
			BaseEntity: BaseEntity{
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		UserID: userID,
	}
}

// NewUserProfile creates a new UserProfile with the given parameters and initializes timestamps.
func NewUserProfile(userID uuid.UUID, avatarURL, bio, timezone *string, preferences datatypes.JSON) *UserProfile {
	now := time.Now()
	return &UserProfile{
		BaseEntityUUID: BaseEntityUUID{
			ID: uuid.New(),
			BaseEntity: BaseEntity{
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
		UserID:      userID,
		AvatarURL:   avatarURL,
		Bio:         bio,
		Preferences: preferences,
		Timezone:    timezone,
	}
}
