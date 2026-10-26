-- Initial Seed Data Sanitized from Legacy System (James2898/PHP-CI-Alumni-Management-System)
-- Modernized with UUID identifiers and BCrypt password hashes

-- 1. Insert Courses
INSERT INTO `courses` (`id`, `code`, `name`, `department`) VALUES
('018e5f2a-7c30-7000-8000-000000000001', 'BSCS', 'Bachelor of Science in Computer Science', 'College of Computer Studies'),
('018e5f2a-7c30-7000-8000-000000000002', 'BSIT', 'Bachelor of Science in Information Technology', 'College of Computer Studies'),
('018e5f2a-7c30-7000-8000-000000000003', 'BSIS', 'Bachelor of Science in Information Systems', 'College of Computer Studies'),
('018e5f2a-7c30-7000-8000-000000000004', 'BSCpE', 'Bachelor of Science in Computer Engineering', 'College of Engineering'),
('018e5f2a-7c30-7000-8000-000000000005', 'BSEE', 'Bachelor of Science in Electronics Engineering', 'College of Engineering'),
('018e5f2a-7c30-7000-8000-000000000006', 'BSIE', 'Bachelor of Science in Industrial Engineering', 'College of Engineering'),
('018e5f2a-7c30-7000-8000-000000000007', 'BSBA', 'Bachelor of Science in Business Administration', 'College of Business Administration'),
('018e5f2a-7c30-7000-8000-000000000008', 'BSA', 'Bachelor of Science in Accountancy', 'College of Business Administration'),
('018e5f2a-7c30-7000-8000-000000000009', 'BAENGL', 'Bachelor of Arts in English', 'College of Arts and Sciences'),
('018e5f2a-7c30-7000-8000-000000000010', 'BAECON', 'Bachelor of Arts in Economics', 'College of Arts and Sciences'),
('018e5f2a-7c30-7000-8000-000000000011', 'BAP', 'Bachelor of Arts in Psychology', 'College of Arts and Sciences'),
('018e5f2a-7c30-7000-8000-000000000012', 'BAMC', 'Bachelor of Arts in Mass Communication', 'College of Arts and Sciences'),
('018e5f2a-7c30-7000-8000-000000000013', 'BAPS', 'Bachelor of Arts in Political Science', 'College of Arts and Sciences');

-- 2. Insert Users (Password for admin is 'admin123', password for alumni is 'password123')
INSERT INTO `users` (`id`, `email`, `password_hash`, `role`, `status`) VALUES
-- Admin account
('018e5f2a-7c30-7000-8001-000000000001', 'admin@alumni.edu', '$2a$10$Wzyf4ukClTYDA890L4a03.5zbeQvRinQ2AHp7RDMztiDL4qqjtx3W', 'admin', 'active'),
-- Alumni 1: Active
('018e5f2a-7c30-7000-8002-000000000001', 'jiro.martinez@alumni.edu', '$2a$10$.OjlAFsrppVc9E4pqvufsOkiahLmEK5E1yLzpbFnJVuRPeKmFxHQe', 'alumni', 'active'),
-- Alumni 2: Active
('018e5f2a-7c30-7000-8002-000000000002', 'farah.eludo@alumni.edu', '$2a$10$.OjlAFsrppVc9E4pqvufsOkiahLmEK5E1yLzpbFnJVuRPeKmFxHQe', 'alumni', 'active'),
-- Alumni 3: Pending verification (for admin dashboard testing)
('018e5f2a-7c30-7000-8002-000000000003', 'jolan.clemente@alumni.edu', '$2a$10$.OjlAFsrppVc9E4pqvufsOkiahLmEK5E1yLzpbFnJVuRPeKmFxHQe', 'alumni', 'pending');

-- 3. Insert Admin Profile
INSERT INTO `admin_profiles` (`id`, `user_id`, `first_name`, `last_name`, `phone`) VALUES
('018e5f2a-7c30-7000-8003-000000000001', '018e5f2a-7c30-7000-8001-000000000001', 'James Emmanuel', 'Martinez', '+63 938 125 7028');

-- 4. Insert Alumni Profiles
INSERT INTO `alumni_profiles` (`id`, `user_id`, `student_number`, `first_name`, `last_name`, `course_id`, `graduation_year`, `phone`, `linkedin_url`, `bio`) VALUES
('018e5f2a-7c30-7000-8004-000000000001', '018e5f2a-7c30-7000-8002-000000000001', '2015-10023', 'Jiro', 'Martinez', '018e5f2a-7c30-7000-8000-000000000002', 2019, '+63 977 472 0937', 'https://linkedin.com/in/jiro-martinez', 'Senior Software Engineer specializing in Go microservices and distributed systems.'),
('018e5f2a-7c30-7000-8004-000000000002', '018e5f2a-7c30-7000-8002-000000000002', '2016-10452', 'Farah', 'Eludo', '018e5f2a-7c30-7000-8000-000000000001', 2020, '+63 917 555 4321', 'https://linkedin.com/in/farah-eludo', 'Frontend Specialist and UX Engineer with a passion for accessible web apps.'),
('018e5f2a-7c30-7000-8004-000000000003', '018e5f2a-7c30-7000-8002-000000000003', '2017-10889', 'Jolan', 'Clemente', '018e5f2a-7c30-7000-8000-000000000007', 2021, '+63 920 123 9876', 'https://linkedin.com/in/jolan-clemente', 'Financial Analyst and Operations Manager in global logistics.');

-- 5. Insert Work Experiences
INSERT INTO `work_experiences` (`id`, `alumni_id`, `company_name`, `position`, `start_date`, `end_date`, `is_current`) VALUES
('018e5f2a-7c30-7000-8005-000000000001', '018e5f2a-7c30-7000-8004-000000000001', 'CloudScale Technologies', 'Senior Backend Engineer', '2022-01-15', NULL, TRUE),
('018e5f2a-7c30-7000-8005-000000000002', '018e5f2a-7c30-7000-8004-000000000001', 'NexGen Solutions', 'Go Developer', '2019-06-01', '2021-12-31', FALSE),
('018e5f2a-7c30-7000-8005-000000000003', '018e5f2a-7c30-7000-8004-000000000002', 'DesignStudio Global', 'Lead Frontend Engineer', '2021-03-01', NULL, TRUE);
