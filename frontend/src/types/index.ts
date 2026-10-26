export type UserRole = 'admin' | 'alumni';
export type UserStatus = 'active' | 'pending' | 'suspended';

export interface User {
  id: string;
  email: string;
  role: UserRole;
  status: UserStatus;
  created_at: string;
  updated_at: string;
  alumni_profile?: AlumniProfile;
  admin_profile?: AdminProfile;
}

export interface Course {
  id: string;
  code: string;
  name: string;
  department: string;
}

export interface AdminProfile {
  id: string;
  user_id: string;
  first_name: string;
  last_name: string;
  phone?: string;
}

export interface WorkExperience {
  id: string;
  alumni_id: string;
  company_name: string;
  position: string;
  start_date: string;
  end_date?: string | null;
  is_current: boolean;
}

export interface AlumniProfile {
  id: string;
  user_id: string;
  student_number: string;
  first_name: string;
  last_name: string;
  course_id: string;
  graduation_year: number;
  avatar_key?: string;
  avatar_url?: string;
  resume_key?: string;
  phone?: string;
  linkedin_url?: string;
  bio?: string;
  course?: Course;
  work_experiences?: WorkExperience[];
}

export interface PaginatedAlumniResponse {
  data: AlumniProfile[];
  total: number;
  page: number;
  page_size: number;
  total_pages: number;
}

export interface PresignedURLResponse {
  upload_url: string;
  object_key: string;
  file_uri: string;
}
