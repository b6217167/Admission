package entity

import (
	"time"

	"gorm.io/gorm"
)

type Nurse struct {
	gorm.Model
	Name     string
	Tel      string `gorm:"uniqueIndex"`
	Pid      string `gorm:"uniqueIndex"`
	Password string

	Patients []Patient `gorm:"foreignKey:Record_byID"`
}

type Patient struct {
	gorm.Model
	Date      time.Time
	Name      string
	Number_id string `gorm:"uniqueIndex"`

	Record_byID *uint
	Record_by   Nurse
	Admissions  []Admission `gorm:"foreignKey:PatientID"`
}

type Roomtypes struct {
	gorm.Model
	Name  string
	Price int

	Rooms []Room `gorm:"foreignKey:RoomtypesID"`
}

type Room struct {
	gorm.Model
	Number string

	RoomtypesID *uint
	Roomtypes   Roomtypes

	Admissions []Admission `gorm:"foreignKey:RoomID"`
}

type Right_Treatment struct {
	gorm.Model
	Name   string
	Detail string

	Admissions []Admission `gorm:"foreignKey:Right_TreatmentID"`
}

type Admission struct {
	gorm.Model
	AdmitTime time.Time

	Right_TreatmentID *uint
	Right_Treatment   Right_Treatment

	RoomID *uint
	Room   Room

	PatientID *uint
	Patient   Patient
}
