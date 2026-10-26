import { AlumniProfile, Course, PaginatedAlumniResponse, PresignedURLResponse, User } from '../types';

const BASE_URL = '/api';

export class ApiError extends Error {
  status: number;
  constructor(message: string, status: number) {
    super(message);
    this.status = status;
  }
}

async function request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
  const token = localStorage.getItem('access_token');
  const headers = new Headers(options.headers || {});

  if (!headers.has('Content-Type') && !(options.body instanceof FormData)) {
    headers.set('Content-Type', 'application/json');
  }

  if (token) {
    headers.set('Authorization', `Bearer ${token}`);
  }

  const res = await fetch(`${BASE_URL}${endpoint}`, {
    ...options,
    headers,
    credentials: 'include',
  });

  if (!res.ok) {
    let errMsg = `Request failed (${res.status})`;
    try {
      const data = await res.json();
      if (data.error) errMsg = data.error;
    } catch {
      // ignore JSON parse fail
    }
    throw new ApiError(errMsg, res.status);
  }

  return res.json();
}

export const api = {
  // Courses
  getCourses: () => request<Course[]>('/courses'),

  // Alumni Directory
  getAlumniDirectory: (params: { course_id?: string; graduation_year?: number; search?: string; page?: number }) => {
    const q = new URLSearchParams();
    if (params.course_id) q.set('course_id', params.course_id);
    if (params.graduation_year) q.set('graduation_year', params.graduation_year.toString());
    if (params.search) q.set('search', params.search);
    if (params.page) q.set('page', params.page.toString());
    return request<PaginatedAlumniResponse>(`/alumni?${q.toString()}`);
  },

  getAlumniById: (id: string) => request<AlumniProfile>(`/alumni/${id}`),

  // Alumni Profile
  getMyProfile: () => request<AlumniProfile>('/alumni/profile/me'),

  updateMyProfile: (data: Partial<AlumniProfile>) =>
    request<AlumniProfile>('/alumni/profile/me', {
      method: 'PUT',
      body: JSON.stringify(data),
    }),

  // S3 Presigned Upload Ticket
  getUploadTicket: (fileName: string, contentType: string, category: 'avatar' | 'resume') =>
    request<PresignedURLResponse>('/alumni/upload-ticket', {
      method: 'POST',
      body: JSON.stringify({
        file_name: fileName,
        content_type: contentType,
        category,
      }),
    }),

  // Direct PUT to S3 / MinIO
  uploadDirectToS3: async (uploadUrl: string, file: File) => {
    const res = await fetch(uploadUrl, {
      method: 'PUT',
      headers: {
        'Content-Type': file.type,
      },
      body: file,
    });
    if (!res.ok) {
      throw new Error(`Direct upload failed with status ${res.status}`);
    }
  },

  // Admin
  getPendingUsers: () => request<User[]>('/admin/pending-users'),

  updateUserStatus: (userId: string, status: 'active' | 'suspended') =>
    request<{ message: string; status: string }>(`/admin/users/${userId}/status`, {
      method: 'PATCH',
      body: JSON.stringify({ status }),
    }),

  createCourse: (data: { code: string; name: string; department: string }) =>
    request<Course>('/admin/courses', {
      method: 'POST',
      body: JSON.stringify(data),
    }),

  deleteCourse: (id: string) =>
    request<{ message: string }>(`/admin/courses/${id}`, {
      method: 'DELETE',
    }),
};
