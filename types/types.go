package types

type Author struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	FullName       string `json:"fullName" gorm:"not null"`
	NickName       string `json:"nickName"`
	Specialization string `json:"specialization" gorm:"not null"`
}

type AuthorInterface struct {
	FullName       string `json:"fullName"`
	NickName       string `json:"nickName"`
	Specialization string `json:"specialization"`
}

type Book struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"title" gorm:"not null"`
	Genre    string `json:"genre" gorm:"not null"`
	ISBN     string `json:"isbn" gorm:"not null"`
	MemberID uint   `json:"memberID" gorm:"default:null"`
	AuthorID uint   `json:"authorID" gorm:"default:null"`
}

type BookInterface struct {
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	ISBN     string `json:"isbn"`
	MemberID uint   `json:"memberID"`
}

type Member struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	FullName string `json:"fullName" gorm:"not null"`
}

type MemberInterface struct {
	FullName string `json:"fullName" binding:"required"`
}
