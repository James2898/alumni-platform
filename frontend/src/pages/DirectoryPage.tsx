import React, { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { api } from '../api/client';
import { Search, GraduationCap, Briefcase, Linkedin, Phone } from 'lucide-react';

export const DirectoryPage: React.FC = () => {
  const [courseFilter, setCourseFilter] = useState('');
  const [yearFilter, setYearFilter] = useState<number | undefined>(undefined);
  const [searchQuery, setSearchQuery] = useState('');
  const [page, setPage] = useState(1);

  const { data: courses = [] } = useQuery({
    queryKey: ['courses'],
    queryFn: api.getCourses,
  });

  const { data: directory, isLoading, isError } = useQuery({
    queryKey: ['alumni', courseFilter, yearFilter, searchQuery, page],
    queryFn: () =>
      api.getAlumniDirectory({
        course_id: courseFilter || undefined,
        graduation_year: yearFilter,
        search: searchQuery || undefined,
        page,
      }),
  });

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      {/* Hero / Header */}
      <div className="mb-8">
        <h1 className="text-3xl font-bold tracking-tight text-slate-900 sm:text-4xl">
          Alumni Directory
        </h1>
        <p className="mt-2 text-base text-slate-600">
          Connect with graduates, explore career paths, and build your professional collegiate network.
        </p>
      </div>

      {/* Filter / Search Bar */}
      <div className="bg-white p-4 sm:p-6 rounded-2xl border border-slate-200 shadow-xs mb-8">
        <div className="grid grid-cols-1 sm:grid-cols-12 gap-4">
          {/* Keyword Search */}
          <div className="sm:col-span-5 relative">
            <Search className="w-5 h-5 absolute left-3.5 top-1/2 -translate-y-1/2 text-slate-400" />
            <input
              type="text"
              placeholder="Search by name or student ID..."
              value={searchQuery}
              onChange={(e) => {
                setSearchQuery(e.target.value);
                setPage(1);
              }}
              className="w-full pl-10 pr-4 py-2.5 bg-slate-50 border border-slate-200 rounded-xl text-sm focus:outline-hidden focus:ring-2 focus:ring-blue-500 focus:bg-white transition-all"
            />
          </div>

          {/* Course Filter */}
          <div className="sm:col-span-4">
            <select
              value={courseFilter}
              onChange={(e) => {
                setCourseFilter(e.target.value);
                setPage(1);
              }}
              className="w-full px-3.5 py-2.5 bg-slate-50 border border-slate-200 rounded-xl text-sm focus:outline-hidden focus:ring-2 focus:ring-blue-500 focus:bg-white transition-all text-slate-700"
            >
              <option value="">All Academic Programs</option>
              {courses.map((c) => (
                <option key={c.id} value={c.id}>
                  {c.code} - {c.name}
                </option>
              ))}
            </select>
          </div>

          {/* Graduation Year */}
          <div className="sm:col-span-3">
            <select
              value={yearFilter || ''}
              onChange={(e) => {
                setYearFilter(e.target.value ? parseInt(e.target.value, 10) : undefined);
                setPage(1);
              }}
              className="w-full px-3.5 py-2.5 bg-slate-50 border border-slate-200 rounded-xl text-sm focus:outline-hidden focus:ring-2 focus:ring-blue-500 focus:bg-white transition-all text-slate-700"
            >
              <option value="">All Batch Years</option>
              {Array.from({ length: 30 }, (_, i) => 2026 - i).map((yr) => (
                <option key={yr} value={yr}>
                  Class of {yr}
                </option>
              ))}
            </select>
          </div>
        </div>
      </div>

      {/* Directory Grid */}
      {isLoading ? (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {[1, 2, 3, 4, 5, 6].map((n) => (
            <div key={n} className="bg-white rounded-2xl p-6 border border-slate-200 animate-pulse h-64" />
          ))}
        </div>
      ) : isError ? (
        <div className="text-center py-12 bg-white rounded-2xl border border-red-200 p-8">
          <p className="text-red-600 font-medium">Failed to load alumni directory. Please verify API connection.</p>
        </div>
      ) : directory?.data.length === 0 ? (
        <div className="text-center py-16 bg-white rounded-2xl border border-slate-200 p-8">
          <GraduationCap className="w-12 h-12 text-slate-400 mx-auto mb-3" />
          <h3 className="text-lg font-semibold text-slate-800">No alumni records found</h3>
          <p className="text-sm text-slate-500 mt-1">Try resetting search filters or keywords.</p>
        </div>
      ) : (
        <>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {directory?.data.map((alumni) => (
              <div
                key={alumni.id}
                className="bg-white rounded-2xl border border-slate-200 p-6 shadow-xs hover:shadow-md transition-shadow flex flex-col justify-between"
              >
                <div>
                  <div className="flex items-start gap-4">
                    {alumni.avatar_url ? (
                      <img
                        src={alumni.avatar_url}
                        alt={`${alumni.first_name} ${alumni.last_name}`}
                        className="w-14 h-14 rounded-full object-cover border-2 border-blue-100"
                      />
                    ) : (
                      <div className="w-14 h-14 rounded-full bg-blue-100 text-blue-700 flex items-center justify-center font-bold text-lg">
                        {alumni.first_name[0]}
                        {alumni.last_name[0]}
                      </div>
                    )}
                    <div className="flex-1 min-w-0">
                      <h3 className="text-lg font-bold text-slate-900 truncate">
                        {alumni.first_name} {alumni.last_name}
                      </h3>
                      <p className="text-xs font-semibold text-blue-600 uppercase tracking-wide">
                        {alumni.course?.code || 'Alumni'} • Batch {alumni.graduation_year}
                      </p>
                      <p className="text-xs text-slate-400 mt-0.5">ID: {alumni.student_number}</p>
                    </div>
                  </div>

                  {/* Bio */}
                  {alumni.bio && (
                    <p className="mt-4 text-xs text-slate-600 line-clamp-2 leading-relaxed">
                      {alumni.bio}
                    </p>
                  )}

                  {/* Work Experience Tag */}
                  {alumni.work_experiences && alumni.work_experiences.length > 0 && (
                    <div className="mt-4 pt-4 border-t border-slate-100 space-y-2">
                      <div className="flex items-center gap-1.5 text-xs text-slate-500 font-medium">
                        <Briefcase className="w-3.5 h-3.5 text-slate-400" />
                        <span>Recent Experience</span>
                      </div>
                      {alumni.work_experiences.slice(0, 2).map((exp) => (
                        <div key={exp.id} className="text-xs flex items-center justify-between">
                          <span className="font-medium text-slate-700 truncate">{exp.position}</span>
                          <span className="text-slate-400 text-[11px] truncate ml-2">@{exp.company_name}</span>
                        </div>
                      ))}
                    </div>
                  )}
                </div>

                {/* Footer Links */}
                <div className="mt-6 pt-4 border-t border-slate-100 flex items-center justify-between text-xs text-slate-500">
                  <div className="flex items-center gap-2">
                    {alumni.phone && (
                      <span title={alumni.phone} className="p-1.5 bg-slate-100 rounded-md hover:bg-slate-200">
                        <Phone className="w-3.5 h-3.5 text-slate-600" />
                      </span>
                    )}
                    {alumni.linkedin_url && (
                      <a
                        href={alumni.linkedin_url}
                        target="_blank"
                        rel="noreferrer"
                        className="p-1.5 bg-blue-50 text-blue-600 rounded-md hover:bg-blue-100"
                      >
                        <Linkedin className="w-3.5 h-3.5" />
                      </a>
                    )}
                  </div>
                  <span className="text-[11px] font-medium text-slate-400">
                    {alumni.course?.department || 'Alumnus'}
                  </span>
                </div>
              </div>
            ))}
          </div>

          {/* Pagination */}
          {directory && directory.total_pages > 1 && (
            <div className="mt-8 flex items-center justify-center gap-2">
              <button
                disabled={page <= 1}
                onClick={() => setPage((p) => Math.max(1, p - 1))}
                className="px-4 py-2 border border-slate-200 rounded-xl text-sm font-medium bg-white disabled:opacity-40"
              >
                Previous
              </button>
              <span className="text-sm font-medium text-slate-600 px-2">
                Page {page} of {directory.total_pages}
              </span>
              <button
                disabled={page >= directory.total_pages}
                onClick={() => setPage((p) => p + 1)}
                className="px-4 py-2 border border-slate-200 rounded-xl text-sm font-medium bg-white disabled:opacity-40"
              >
                Next
              </button>
            </div>
          )}
        </>
      )}
    </div>
  );
};
