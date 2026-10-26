package domain

type RegisterRequest struct {
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required,min=8"`
	StudentNumber  string `json:"student_number" binding:"required"`
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	CourseID       string `json:"course_id" binding:"required"`
	GraduationYear int    `json:"graduation_year" binding:"required,gte=1950,lte=2030"`
	Phone          string `json:"phone"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	User         User   `json:"user"`
}

type UpdateProfileRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone"`
	LinkedinURL string `json:"linkedin_url"`
	Bio         string `json:"bio"`
	AvatarKey   string `json:"avatar_key"`
}

type PresignedURLRequest struct {
	FileName    string `json:"file_name" binding:"required"`
	ContentType string `json:"content_type" binding:"required"`
	Category    string `json:"category" binding:"required,oneof=avatar resume"`
}

type PresignedURLResponse struct {
	UploadURL string `json:"upload_url"`
	ObjectKey string `json:"object_key"`
	FileURI   string `json:"file_uri"`
}

type AlumniFilter struct {
	CourseID       string `form:"course_id"`
	GraduationYear int    `form:"graduation_year"`
	Search         string `form:"search"`
	Company        string `form:"company"`
	Page           int    `form:"page,default=1"`
	PageSize       int    `form:"page_size,default=12"`
}

type PaginatedAlumniResponse struct {
	Data       []AlumniProfile `json:"data"`
	Total      int64           `json:"total"`
	Page       int             `json:"page"`
	PageSize   int             `json:"page_size"`
	TotalPages int             `json:"total_pages"`
}

type UpdateStatusRequest struct {
	Status UserStatus `json:"status" binding:"required,oneof=active pending suspended"`
}

type CourseRequest struct {
	Code       string `json:"code" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Department string `json:"department" binding:"required"`
}
