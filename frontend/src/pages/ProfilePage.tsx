import React, { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { api } from '../api/client';
import { useAuth } from '../context/AuthContext';
import { Camera, Upload, AlertCircle, Linkedin } from 'lucide-react';

export const ProfilePage: React.FC = () => {
  const { user } = useAuth();
  const queryClient = useQueryClient();

  const { data: profile, isLoading } = useQuery({
    queryKey: ['my-profile'],
    queryFn: api.getMyProfile,
  });

  const [formData, setFormData] = useState({
    first_name: '',
    last_name: '',
    phone: '',
    linkedin_url: '',
    bio: '',
  });

  const [isEditing, setIsEditing] = useState(false);
  const [uploadStatus, setUploadStatus] = useState<string | null>(null);
  const [errorMsg, setErrorMsg] = useState<string | null>(null);

  // Sync initial state when profile loads
  React.useEffect(() => {
    if (profile) {
      setFormData({
        first_name: profile.first_name || '',
        last_name: profile.last_name || '',
        phone: profile.phone || '',
        linkedin_url: profile.linkedin_url || '',
        bio: profile.bio || '',
      });
    }
  }, [profile]);

  const updateProfileMutation = useMutation({
    mutationFn: (data: Partial<typeof formData & { avatar_key?: string }>) =>
      api.updateMyProfile(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['my-profile'] });
      queryClient.invalidateQueries({ queryKey: ['alumni'] });
      setIsEditing(false);
    },
    onError: (err: Error) => {
      setErrorMsg(err.message);
    },
  });

  // Handle direct S3 Presigned URL upload
  const handleAvatarFileChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    try {
      setUploadStatus('Requesting S3 presigned ticket...');
      setErrorMsg(null);

      // 1. Get presigned PUT URL from backend
      const ticket = await api.getUploadTicket(file.name, file.type, 'avatar');

      // 2. Upload file directly to S3/MinIO
      setUploadStatus('Uploading directly to S3 storage bucket...');
      await api.uploadDirectToS3(ticket.upload_url, file);

      // 3. Update alumni profile record with object key
      setUploadStatus('Saving avatar reference...');
      await updateProfileMutation.mutateAsync({ avatar_key: ticket.object_key });

      setUploadStatus('Avatar updated successfully!');
      setTimeout(() => setUploadStatus(null), 3000);
    } catch (err: unknown) {
      setErrorMsg(err instanceof Error ? err.message : 'Upload failed');
      setUploadStatus(null);
    }
  };

  if (isLoading) {
    return (
      <div className="max-w-4xl mx-auto py-12 px-4">
        <div className="bg-white rounded-2xl p-8 border border-slate-200 animate-pulse h-96" />
      </div>
    );
  }

  if (!profile) {
    return (
      <div className="max-w-md mx-auto py-12 px-4 text-center">
        <p className="text-slate-600">Alumni profile record not found.</p>
      </div>
    );
  }

  return (
    <div className="max-w-4xl mx-auto py-8 px-4 sm:px-6">
      {/* Header card */}
      <div className="bg-white rounded-3xl border border-slate-200 p-6 sm:p-8 shadow-xs mb-6">
        <div className="flex flex-col sm:flex-row items-center gap-6">
          {/* Avatar with S3 upload button */}
          <div className="relative group">
            {profile.avatar_url ? (
              <img
                src={profile.avatar_url}
                alt="Avatar"
                className="w-24 h-24 sm:w-28 sm:h-28 rounded-full object-cover border-4 border-white shadow-md"
              />
            ) : (
              <div className="w-24 h-24 sm:w-28 sm:h-28 rounded-full bg-blue-100 text-blue-700 flex items-center justify-center font-bold text-3xl shadow-xs">
                {profile.first_name[0]}
                {profile.last_name[0]}
              </div>
            )}

            <label className="absolute bottom-0 right-0 p-2 bg-blue-600 hover:bg-blue-700 text-white rounded-full shadow-md cursor-pointer transition-transform hover:scale-110">
              <Camera className="w-4 h-4" />
              <input
                type="file"
                accept="image/*"
                onChange={handleAvatarFileChange}
                className="hidden"
              />
            </label>
          </div>

          <div className="text-center sm:text-left flex-1">
            <h1 className="text-2xl sm:text-3xl font-bold text-slate-900">
              {profile.first_name} {profile.last_name}
            </h1>
            <p className="text-sm font-semibold text-blue-600 mt-1">
              {profile.course?.name} ({profile.course?.code}) • Class of {profile.graduation_year}
            </p>
            <p className="text-xs text-slate-400 mt-0.5">Student ID: {profile.student_number}</p>
          </div>

          <div className="shrink-0">
            <button
              onClick={() => setIsEditing(!isEditing)}
              className="px-4 py-2 border border-slate-200 hover:border-slate-300 bg-slate-50 hover:bg-slate-100 rounded-xl text-xs font-semibold text-slate-700 transition-colors"
            >
              {isEditing ? 'Cancel Editing' : 'Edit Profile'}
            </button>
          </div>
        </div>

        {uploadStatus && (
          <div className="mt-4 p-3 bg-blue-50 border border-blue-200 text-blue-700 rounded-xl text-xs flex items-center gap-2">
            <Upload className="w-4 h-4 animate-bounce" />
            <span>{uploadStatus}</span>
          </div>
        )}

        {errorMsg && (
          <div className="mt-4 p-3 bg-red-50 border border-red-200 text-red-700 rounded-xl text-xs flex items-center gap-2">
            <AlertCircle className="w-4 h-4 shrink-0" />
            <span>{errorMsg}</span>
          </div>
        )}
      </div>

      {/* Main Content */}
      <div className="bg-white rounded-3xl border border-slate-200 p-6 sm:p-8 shadow-xs">
        {isEditing ? (
          <form
            onSubmit={(e) => {
              e.preventDefault();
              updateProfileMutation.mutate(formData);
            }}
            className="space-y-4"
          >
            <h3 className="text-lg font-bold text-slate-900 mb-4">Edit Profile Information</h3>

            <div className="grid grid-cols-2 gap-4">
              <div>
                <label className="block text-xs font-semibold text-slate-700 mb-1">First Name</label>
                <input
                  type="text"
                  value={formData.first_name}
                  onChange={(e) => setFormData({ ...formData, first_name: e.target.value })}
                  className="w-full px-3 py-2 bg-slate-50 border border-slate-200 rounded-xl text-sm focus:outline-hidden focus:ring-2 focus:ring-blue-500"
                />
              </div>
              <div>
                <label className="block text-xs font-semibold text-slate-700 mb-1">Last Name</label>
                <input
                  type="text"
                  value={formData.last_name}
                  onChange={(e) => setFormData({ ...formData, last_name: e.target.value })}
                  className="w-full px-3 py-2 bg-slate-50 border border-slate-200 rounded-xl text-sm focus:outline-hidden focus:ring-2 focus:ring-blue-500"
                />
              </div>
            </div>

            <div>
              <label className="block text-xs font-semibold text-slate-700 mb-1">Contact Phone</label>
              <input
                type="text"
                value={formData.phone}
                onChange={(e) => setFormData({ ...formData, phone: e.target.value })}
                className="w-full px-3 py-2 bg-slate-50 border border-slate-200 rounded-xl text-sm focus:outline-hidden focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label className="block text-xs font-semibold text-slate-700 mb-1">LinkedIn Profile URL</label>
              <input
                type="url"
                value={formData.linkedin_url}
                onChange={(e) => setFormData({ ...formData, linkedin_url: e.target.value })}
                placeholder="https://linkedin.com/in/username"
                className="w-full px-3 py-2 bg-slate-50 border border-slate-200 rounded-xl text-sm focus:outline-hidden focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div>
              <label className="block text-xs font-semibold text-slate-700 mb-1">Professional Bio</label>
              <textarea
                rows={4}
                value={formData.bio}
                onChange={(e) => setFormData({ ...formData, bio: e.target.value })}
                placeholder="Share your current career focus, skills, or achievements..."
                className="w-full px-3 py-2 bg-slate-50 border border-slate-200 rounded-xl text-sm focus:outline-hidden focus:ring-2 focus:ring-blue-500"
              />
            </div>

            <div className="flex justify-end gap-3 pt-4 border-t border-slate-100">
              <button
                type="button"
                onClick={() => setIsEditing(false)}
                className="px-4 py-2 border border-slate-200 rounded-xl text-sm font-semibold text-slate-600 hover:bg-slate-50"
              >
                Cancel
              </button>
              <button
                type="submit"
                disabled={updateProfileMutation.isPending}
                className="px-5 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-xl text-sm font-semibold disabled:opacity-50"
              >
                {updateProfileMutation.isPending ? 'Saving...' : 'Save Changes'}
              </button>
            </div>
          </form>
        ) : (
          <div className="space-y-6">
            <div>
              <h3 className="text-xs font-bold text-slate-400 uppercase tracking-wider mb-2">About Me</h3>
              <p className="text-sm text-slate-700 leading-relaxed">
                {profile.bio || 'No professional bio added yet. Click "Edit Profile" to write a summary.'}
              </p>
            </div>

            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 pt-6 border-t border-slate-100">
              <div>
                <h4 className="text-xs font-bold text-slate-400 uppercase tracking-wider mb-1">Email</h4>
                <p className="text-sm font-medium text-slate-800">{user?.email}</p>
              </div>
              <div>
                <h4 className="text-xs font-bold text-slate-400 uppercase tracking-wider mb-1">Phone</h4>
                <p className="text-sm font-medium text-slate-800">{profile.phone || 'Not provided'}</p>
              </div>
              <div>
                <h4 className="text-xs font-bold text-slate-400 uppercase tracking-wider mb-1">LinkedIn</h4>
                {profile.linkedin_url ? (
                  <a
                    href={profile.linkedin_url}
                    target="_blank"
                    rel="noreferrer"
                    className="text-sm font-medium text-blue-600 hover:underline flex items-center gap-1"
                  >
                    <Linkedin className="w-4 h-4" />
                    <span>View Profile</span>
                  </a>
                ) : (
                  <p className="text-sm font-medium text-slate-400">None linked</p>
                )}
              </div>
              <div>
                <h4 className="text-xs font-bold text-slate-400 uppercase tracking-wider mb-1">Account Status</h4>
                <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-semibold bg-green-100 text-green-800 capitalize">
                  {user?.status}
                </span>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};
