package entity

import (
	//"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	database.AutoMigrate(
		&Patient{}, &Nurse{}, &Room{}, &Roomtypes{}, &Admission{}, &Right_Treatment{},
	)

	db = database

	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	//Nurse
	db.Model(&Nurse{}).Create(&Nurse{
		Name:     "Sumaree",
		Tel:      "0988888888",
		Pid:      "1155523456789",
		Password: string(password),
	})
	db.Model(&Nurse{}).Create(&Nurse{
		Name:     "Nattawee",
		Tel:      "0866663333",
		Pid:      "1166678901234",
		Password: string(password),
	})

	//Patient
	var Sumaree Nurse
	var Nattawee Nurse
	db.Raw("SELECT * FROM nurses WHERE pid =?", "1155523456789").Scan(&Sumaree)
	db.Raw("SELECT * FROM nurses WHERE pid =?", "1166678901234").Scan(&Nattawee)

	huy := Patient{
		Name:      "huy",
		Number_id: "12xxxxxxxxxxx",
		Date:      time.Now(),
		Record_by: Sumaree,
	}
	db.Model(&Patient{}).Create(&huy)

	guli := Patient{
		Name:      "guli",
		Number_id: "56xxxxxxxxxxx",
		Date:      time.Now(),
		Record_by: Nattawee,
	}
	db.Model(&Patient{}).Create(&guli)

	ball := Patient{
		Name:      "ball",
		Number_id: "79xxxxxxxxxxx",
		Date:      time.Now(),
		Record_by: Nattawee,
	}
	db.Model(&Patient{}).Create(&ball)

	//Roomtypes
	all := Roomtypes{
		Name:  "ห้องรวม",
		Price: 500,
	}
	db.Model(&Roomtypes{}).Create(&all)

	std := Roomtypes{
		Name:  "Standard Room",
		Price: 1500,
	}
	db.Model(&Roomtypes{}).Create(&std)

	pm := Roomtypes{
		Name:  "Premier Room",
		Price: 2500,
	}
	db.Model(&Roomtypes{}).Create(&pm)

	np := Roomtypes{
		Name:  "Negative pressure",
		Price: 4000,
	}
	db.Model(&Roomtypes{}).Create(&np)

	//Room
	A10001 := Room{
		Number:    "10001",
		Roomtypes: all,
	}
	db.Model(&Room{}).Create(&A10001)

	std20001 := Room{
		Number:    "STD-20001",
		Roomtypes: std,
	}
	db.Model(&Room{}).Create(&std20001)

	std20002 := Room{
		Number:    "STD-20002",
		Roomtypes: std,
	}
	db.Model(&Room{}).Create(&std20002)

	pm30001 := Room{
		Number:    "PM-30001",
		Roomtypes: pm,
	}
	db.Model(&Room{}).Create(&pm30001)

	np40001 := Room{
		Number:    "NP-40001",
		Roomtypes: np,
	}
	db.Model(&Room{}).Create(&np40001)

	//Right_Treatment
	NO0001 := Right_Treatment{
		Name:   "ไม่ใช้สิทธิการรักษา",
		Detail: "ไม่มีส่วนลด",
	}
	db.Model(&Right_Treatment{}).Create(&NO0001)

	GM0001 := Right_Treatment{
		Name:   "บัตร30",
		Detail: "ลดเหลือ 30 บาท",
	}
	db.Model(&Right_Treatment{}).Create(&GM0001)

	GM0002 := Right_Treatment{
		Name:   "ข้าราชการ",
		Detail: "ลดสูงสุด 20000 บาท",
	}
	db.Model(&Right_Treatment{}).Create(&GM0002)

	IV0001 := Right_Treatment{
		Name:   "ประกันชั้น1",
		Detail: "ลดสูงสุด 30000 บาท",
	}
	db.Model(&Right_Treatment{}).Create(&IV0001)

	IV0002 := Right_Treatment{
		Name:   "ประกันชั้น2",
		Detail: "ลดสูงสุด 15000 บาท",
	}
	db.Model(&Right_Treatment{}).Create(&IV0002)

	IV0003 := Right_Treatment{
		Name:   "ประกันชั้น3",
		Detail: "ลดสูงสุด 7000 บาท",
	}
	db.Model(&Right_Treatment{}).Create(&IV0003)

	//admission1
	/*db.Model(&Admission{}).Create(&Admission{
		AdmitTime:       time.Now(),
		Patient:         huy,
		Room:            std20001,
		Right_Treatment: GM0001,
	})
	//admission2
	db.Model(&Admission{}).Create(&Admission{
		AdmitTime:       time.Now(),
		Patient:         guli,
		Room:            np40001,
		Right_Treatment: IV0003,
	})
	//admission3
	db.Model(&Admission{}).Create(&Admission{
		AdmitTime:       time.Now(),
		Patient:         ball,
		Room:            pm30001,
		Right_Treatment: IV0002,
	})*/
}
