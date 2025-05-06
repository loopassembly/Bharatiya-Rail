package models


import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}



type CouplerDetails struct {
	CV  *string `json:"cv" gorm:"column:cv"`
	NCV *string `json:"ncv" gorm:"column:ncv"`
}

type KnuckleDetails struct {
	CV  *string `json:"cv" gorm:"column:cv"`
	NCV *string `json:"ncv" gorm:"column:ncv"`
}

type DraftGearDetails struct {
	CV  *string `json:"cv" gorm:"column:cv"`
	NCV *string `json:"ncv" gorm:"column:ncv"`
}

type StiffenerPlateDetails struct {
	CV  *string `json:"cv" gorm:"column:cv"`
	NCV *string `json:"ncv" gorm:"column:ncv"`
}

type SilentBlockDetails struct {
	CV  *string `json:"cv" gorm:"column:cv"`
	NCV *string `json:"ncv" gorm:"column:ncv"`
}

type MUCDDetails struct {
	CV  *string `json:"cv" gorm:"column:cv"`
	NCV *string `json:"ncv" gorm:"column:ncv"`
}

type CBCMaterialRecord struct {
	gorm.Model // Add this to include ID, CreatedAt, UpdatedAt, and DeletedAt fields

	UserID               uint                  `json:"user_id" gorm:"not null"` // Reference to User ID
	Date                 string                `json:"date" gorm:"column:date"`
	CoachNo              string                `json:"coach_no" gorm:"column:coach_no"`
	Division             string                `json:"division" gorm:"column:division"`
	Make                 string                `json:"make" gorm:"column:make"`
	Code                 string                `json:"code" gorm:"column:code"`
	SSIIOrSSIII          string                `json:"SSII/SSIII" gorm:"column:ssii_ssiii"`

	CouplerDetails       *CouplerDetails       `json:"coupler_details" gorm:"embedded;embeddedPrefix:coupler_"`
	KnuckleDetails       *KnuckleDetails       `json:"knuckle_details" gorm:"embedded;embeddedPrefix:knuckle_"`
	DraftGearDetails     *DraftGearDetails     `json:"draft_gear_details" gorm:"embedded;embeddedPrefix:draft_gear_"`
	StiffenerPlateDetails *StiffenerPlateDetails `json:"stiffener_plate_details" gorm:"embedded;embeddedPrefix:stiffener_plate_"`
	SilentBlockDetails   *SilentBlockDetails   `json:"silent_block_details" gorm:"embedded;embeddedPrefix:silent_block_"`
	MUCDDetails          *MUCDDetails          `json:"mucd_details" gorm:"embedded;embeddedPrefix:mucd_"`
}

