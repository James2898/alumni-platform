import React, { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { api } from '../api/client';
import { ShieldCheck, Check, X, BookOpen, Plus, Trash2 } from 'lucide-react';

export const AdminPage: React.FC = () => {
  const queryClient = useQueryClient();
  const [newCourseCode, setNewCourseCode] = useState('');
  const [newCourseName, setNewCourseName] = useState('');
  const [newCourseDept, setNewCourseDept] = useState('');
  const [msg, setMsg] = useState<string | null>(null);

  const { data: pendingUsers = [], isLoading: loadingPending } = useQuery({
    queryKey: ['pending-users'],
    queryFn: api.getPendingUsers,
  });

  const { data: courses = [], isLoading: loadingCourses } = useQuery({
    queryKey: ['courses'],
    queryFn: api.getCourses,
  });

  const updateStatusMutation = useMutation({
    mutationFn: ({ userId, status }: { userId: string; status: 'active' | 'suspended' }) =>
      api.updateUserStatus(userId, status),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['pending-users'] });
      queryClient.invalidateQueries({ queryKey: ['alumni'] });
      setMsg('User status updated successfully.');
      setTimeout(() => setMsg(null), 3000);
    },
  });

  const createCourseMutation = useMutation({
    mutationFn: (data: { code: string; name: string; department: string }) =>
      api.createCourse(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['courses'] });
      setNewCourseCode('');
      setNewCourseName('');
      setNewCourseDept('');
      setMsg('Course added successfully.');
      setTimeout(() => setMsg(null), 3000);
    },
  });

  const deleteCourseMutation = useMutation({
    mutationFn: (id: string) => api.deleteCourse(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['courses'] });
      setMsg('Course removed.');
      setTimeout(() => setMsg(null), 3000);
    },
  });

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8 space-y-8">
      <div>
        <h1 className="text-3xl font-bold tracking-tight text-slate-900">Admin Control Center</h1>
        <p className="mt-2 text-sm text-slate-600">
          Verify pending alumni accounts and maintain academic course offerings.
        </p>
      </div>

      {msg && (
        <div className="p-3 bg-green-50 border border-green-200 text-green-700 rounded-xl text-xs font-medium">
          {msg}
        </div>
      )}

      {/* Section 1: Pending Verifications */}
      <div className="bg-white rounded-3xl border border-slate-200 p-6 sm:p-8 shadow-xs">
        <div className="flex items-center gap-2 mb-6">
          <ShieldCheck className="w-5 h-5 text-blue-600" />
          <h2 className="text-lg font-bold text-slate-900">
            Pending Alumni Verifications ({pendingUsers.length})
          </h2>
        </div>

        {loadingPending ? (
          <div className="h-32 bg-slate-50 rounded-2xl animate-pulse" />
        ) : pendingUsers.length === 0 ? (
          <div className="text-center py-8 text-slate-500 text-sm">
            All registered alumni accounts have been verified. No pending approvals.
          </div>
        ) : (
          <div className="overflow-x-auto">
            <table className="w-full text-left text-xs">
              <thead className="bg-slate-50 border-b border-slate-200 text-slate-500 font-semibold uppercase">
                <tr>
                  <th className="py-3 px-4">Name & Email</th>
                  <th className="py-3 px-4">Student ID</th>
                  <th className="py-3 px-4">Degree / Batch</th>
                  <th className="py-3 px-4 text-right">Actions</th>
                </tr>
              </thead>
              <tbody className="divide-y divide-slate-100">
                {pendingUsers.map((u) => (
                  <tr key={u.id} className="hover:bg-slate-50">
                    <td className="py-3 px-4">
                      <div className="font-semibold text-slate-900">
                        {u.alumni_profile ? `${u.alumni_profile.first_name} ${u.alumni_profile.last_name}` : 'Alumnus'}
                      </div>
                      <div className="text-slate-400">{u.email}</div>
                    </td>
                    <td className="py-3 px-4 font-mono font-medium text-slate-700">
                      {u.alumni_profile?.student_number}
                    </td>
                    <td className="py-3 px-4">
                      <span className="font-semibold text-blue-600">
                        {u.alumni_profile?.course?.code || 'Course'}
                      </span>
                      <span className="text-slate-500"> • Class of {u.alumni_profile?.graduation_year}</span>
                    </td>
                    <td className="py-3 px-4 text-right space-x-2">
                      <button
                        onClick={() => updateStatusMutation.mutate({ userId: u.id, status: 'active' })}
                        className="inline-flex items-center gap-1 px-3 py-1.5 bg-green-600 hover:bg-green-700 text-white rounded-lg font-semibold transition-colors shadow-xs"
                      >
                        <Check className="w-3.5 h-3.5" />
                        <span>Approve</span>
                      </button>
                      <button
                        onClick={() => updateStatusMutation.mutate({ userId: u.id, status: 'suspended' })}
                        className="inline-flex items-center gap-1 px-3 py-1.5 bg-red-50 hover:bg-red-100 text-red-600 rounded-lg font-semibold transition-colors"
                      >
                        <X className="w-3.5 h-3.5" />
                        <span>Decline</span>
                      </button>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>

      {/* Section 2: Manage Academic Programs */}
      <div className="bg-white rounded-3xl border border-slate-200 p-6 sm:p-8 shadow-xs">
        <div className="flex items-center gap-2 mb-6">
          <BookOpen className="w-5 h-5 text-blue-600" />
          <h2 className="text-lg font-bold text-slate-900">Academic Programs Management</h2>
        </div>

        {/* Add Course Form */}
        <form
          onSubmit={(e) => {
            e.preventDefault();
            if (!newCourseCode || !newCourseName || !newCourseDept) return;
            createCourseMutation.mutate({
              code: newCourseCode,
              name: newCourseName,
              department: newCourseDept,
            });
          }}
          className="grid grid-cols-1 sm:grid-cols-4 gap-3 mb-6 p-4 bg-slate-50 rounded-2xl border border-slate-200"
        >
          <div>
            <input
              type="text"
              placeholder="Code (e.g. BSCS)"
              value={newCourseCode}
              onChange={(e) => setNewCourseCode(e.target.value)}
              className="w-full px-3 py-2 bg-white border border-slate-200 rounded-xl text-xs focus:ring-2 focus:ring-blue-500 focus:outline-hidden"
              required
            />
          </div>
          <div className="sm:col-span-2">
            <input
              type="text"
              placeholder="Full Degree Name"
              value={newCourseName}
              onChange={(e) => setNewCourseName(e.target.value)}
              className="w-full px-3 py-2 bg-white border border-slate-200 rounded-xl text-xs focus:ring-2 focus:ring-blue-500 focus:outline-hidden"
              required
            />
          </div>
          <div className="flex gap-2">
            <input
              type="text"
              placeholder="College / Department"
              value={newCourseDept}
              onChange={(e) => setNewCourseDept(e.target.value)}
              className="w-full px-3 py-2 bg-white border border-slate-200 rounded-xl text-xs focus:ring-2 focus:ring-blue-500 focus:outline-hidden"
              required
            />
            <button
              type="submit"
              disabled={createCourseMutation.isPending}
              className="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-xl text-xs font-semibold shrink-0 flex items-center gap-1 shadow-xs"
            >
              <Plus className="w-3.5 h-3.5" />
              <span>Add</span>
            </button>
          </div>
        </form>

        {/* Course List */}
        {loadingCourses ? (
          <div className="h-32 bg-slate-50 rounded-2xl animate-pulse" />
        ) : (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-3">
            {courses.map((c) => (
              <div
                key={c.id}
                className="p-3.5 border border-slate-200 rounded-xl flex items-center justify-between hover:border-slate-300 transition-colors"
              >
                <div>
                  <span className="font-bold text-xs text-blue-700 bg-blue-50 px-2 py-0.5 rounded-md mr-2">
                    {c.code}
                  </span>
                  <span className="text-xs font-medium text-slate-800">{c.name}</span>
                  <div className="text-[10px] text-slate-400 mt-0.5">{c.department}</div>
                </div>
                <button
                  onClick={() => deleteCourseMutation.mutate(c.id)}
                  title="Delete Course"
                  className="p-1.5 text-slate-400 hover:text-red-600 rounded-lg transition-colors"
                >
                  <Trash2 className="w-3.5 h-3.5" />
                </button>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};
