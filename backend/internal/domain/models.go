package domain

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleAlumni UserRole = "alumni"
)

type UserStatus string

const (
	StatusActive    UserStatus = "active"
	StatusPending   UserStatus = "pending"
	StatusSuspended UserStatus = "suspended"
)

type User struct {
	ID           string         `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Email        string         `gorm:"uniqueIndex;type:varchar(255);not null" json:"email"`
	PasswordHash string         `gorm:"type:varchar(255);not null" json:"-"`
	Role         UserRole       `gorm:"type:enum('admin','alumni');default:'alumni';not null" json:"role"`
	Status       UserStatus     `gorm:"type:enum('active','pending','suspended');default:'pending';not null" json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	AlumniProfile *AlumniProfile `gorm:"foreignKey:UserID" json:"alumni_profile,omitempty"`
	AdminProfile  *AdminProfile  `gorm:"foreignKey:UserID" json:"admin_profile,omitempty"`
}

type Course struct {
	ID         string    `gorm:"primaryKey;type:varchar(36)" json:"id"`
	Code       string    `gorm:"uniqueIndex;type:varchar(50);not null" json:"code"`
	Name       string    `gorm:"type:varchar(150);not null" json:"name"`
	Department string    `gorm:"type:varchar(150);not null" json:"department"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type AdminProfile struct {
	ID        string    `gorm:"primaryKey;type:varchar(36)" json:"id"`
	UserID    string    `gorm:"uniqueIndex;type:varchar(36);not null" json:"user_id"`
	FirstName string    `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(100);not null" json:"last_name"`
	Phone     string    `gorm:"type:varchar(50)" json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AlumniProfile struct {
	ID             string    `gorm:"primaryKey;type:varchar(36)" json:"id"`
	UserID         string    `gorm:"uniqueIndex;type:varchar(36);not null" json:"user_id"`
	StudentNumber  string    `gorm:"uniqueIndex;type:varchar(64);not null" json:"student_number"`
	FirstName      string    `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName       string    `gorm:"type:varchar(100);not null" json:"last_name"`
	CourseID       string    `gorm:"type:varchar(36);not null" json:"course_id"`
	GraduationYear int       `gorm:"not null" json:"graduation_year"`
	AvatarKey      string    `gorm:"type:varchar(255)" json:"avatar_key"`
	AvatarURL      string    `gorm:"-" json:"avatar_url,omitempty"`
	ResumeKey      string    `gorm:"type:varchar(255)" json:"resume_key"`
	Phone          string    `gorm:"type:varchar(50)" json:"phone"`
	LinkedinURL    string    `gorm:"type:varchar(255)" json:"linkedin_url"`
	Bio            string    `gorm:"type:text" json:"bio"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	Course           *Course          `gorm:"foreignKey:CourseID" json:"course,omitempty"`
	WorkExperiences  []WorkExperience `gorm:"foreignKey:AlumniID" json:"work_experiences,omitempty"`
	User             *User            `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

type WorkExperience struct {
	ID          string     `gorm:"primaryKey;type:varchar(36)" json:"id"`
	AlumniID    string     `gorm:"type:varchar(36);not null;index" json:"alumni_id"`
	CompanyName string     `gorm:"type:varchar(150);not null" json:"company_name"`
	Position    string     `gorm:"type:varchar(100);not null" json:"position"`
	StartDate   time.Time  `gorm:"type:date;not null" json:"start_date"`
	EndDate     *time.Time `gorm:"type:date" json:"end_date"`
	IsCurrent   bool       `gorm:"default:false;not null" json:"is_current"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
