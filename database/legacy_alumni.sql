-- phpMyAdmin SQL Dump
-- version 4.5.1
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: Mar 21, 2019 at 08:56 AM
-- Server version: 10.1.19-MariaDB
-- PHP Version: 7.0.13

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Database: `alumni`
--

-- --------------------------------------------------------

--
-- Table structure for table `admin`
--

CREATE TABLE `admin` (
  `admin_ID` int(11) NOT NULL,
  `admin_profile_picture` text COLLATE utf8mb4_bin NOT NULL,
  `admin_fname` text COLLATE utf8mb4_bin NOT NULL,
  `admin_mname` text COLLATE utf8mb4_bin NOT NULL,
  `admin_lname` text COLLATE utf8mb4_bin NOT NULL,
  `admin_address` text COLLATE utf8mb4_bin,
  `admin_email` text COLLATE utf8mb4_bin NOT NULL,
  `admin_facebook` text COLLATE utf8mb4_bin,
  `admin_lno` text COLLATE utf8mb4_bin NOT NULL,
  `admin_cno` text COLLATE utf8mb4_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `admin`
--

INSERT INTO `admin` (`admin_ID`, `admin_profile_picture`, `admin_fname`, `admin_mname`, `admin_lname`, `admin_address`, `admin_email`, `admin_facebook`, `admin_lno`, `admin_cno`) VALUES
(1, '1.jpg?12341', 'James Emmanuel', 'Montealegre', 'Martinez', 'Philippines', 'amajamesmartinez@gmail.com', 'James2898', '123', '09381257028');

-- --------------------------------------------------------

--
-- Table structure for table `alumni`
--

CREATE TABLE `alumni` (
  `alumni_ID` int(11) NOT NULL,
  `alumni_student_ID` text COLLATE utf8mb4_bin,
  `alumni_profile_picture` text COLLATE utf8mb4_bin,
  `alumni_fname` text COLLATE utf8mb4_bin,
  `alumni_mname` text COLLATE utf8mb4_bin,
  `alumni_lname` text COLLATE utf8mb4_bin,
  `alumni_gender` text COLLATE utf8mb4_bin,
  `alumni_address` text COLLATE utf8mb4_bin,
  `alumni_cno` text COLLATE utf8mb4_bin,
  `alumni_lno` text COLLATE utf8mb4_bin,
  `alumni_email` text COLLATE utf8mb4_bin,
  `alumni_facebook` text COLLATE utf8mb4_bin,
  `alumni_linkedin` text COLLATE utf8mb4_bin,
  `alumni_website` text COLLATE utf8mb4_bin,
  `alumni_degree` text COLLATE utf8mb4_bin,
  `alumni_major` text COLLATE utf8mb4_bin,
  `alumni_grad_year` int(11) DEFAULT NULL,
  `alumni_register_date` text COLLATE utf8mb4_bin
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `alumni`
--

INSERT INTO `alumni` (`alumni_ID`, `alumni_student_ID`, `alumni_profile_picture`, `alumni_fname`, `alumni_mname`, `alumni_lname`, `alumni_gender`, `alumni_address`, `alumni_cno`, `alumni_lno`, `alumni_email`, `alumni_facebook`, `alumni_linkedin`, `alumni_website`, `alumni_degree`, `alumni_major`, `alumni_grad_year`, `alumni_register_date`) VALUES
(16, '123', '123.jpg?1234', 'Jiro', 'Cruz', 'Martinez', 'Male', '', '09774720937', '', 'zion.martin0888@gmail.com', '', '', '', '2', '6', 2000, '2019-03-15'),
(17, '1234', '1234.png?1234', 'Farah', 'Llesis', 'Eludo', 'Female', 'Philippines', '', '', 'farah.eludo2398@gmail.com', '', '', '', '1', '14', 2010, '2019-03-20'),
(18, '12345', '12345.jpg?1234', 'Jolan', 'Castillio', 'Clemente', 'Male', 'Philippines', '123', '123', 'amajamesmartinez@gmail.com', '123', '123', '123', '7', '2', 2000, '2019-03-20');

-- --------------------------------------------------------

--
-- Table structure for table `announcement`
--

CREATE TABLE `announcement` (
  `announcement_ID` int(11) NOT NULL,
  `announcement_title` text COLLATE utf8mb4_bin,
  `announcement_subject` text COLLATE utf8mb4_bin,
  `announcement_content` text COLLATE utf8mb4_bin,
  `announcement_date` text COLLATE utf8mb4_bin
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- Table structure for table `appointment`
--

CREATE TABLE `appointment` (
  `appointment_ID` int(11) NOT NULL,
  `appointment_alumni_ID` text COLLATE utf8mb4_bin,
  `appointment_details` text COLLATE utf8mb4_bin,
  `appointment_date_start` text COLLATE utf8mb4_bin,
  `appointment_date_end` text COLLATE utf8mb4_bin,
  `appointment_time_start` text COLLATE utf8mb4_bin,
  `appointment_time_end` text COLLATE utf8mb4_bin,
  `appointment_status` text COLLATE utf8mb4_bin
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- Table structure for table `ci_sessions`
--

CREATE TABLE `ci_sessions` (
  `id` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  `ip_address` varchar(100) COLLATE utf8mb4_bin NOT NULL,
  `timestamp` int(10) UNSIGNED NOT NULL DEFAULT '0',
  `data` blob NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- Table structure for table `courses`
--

CREATE TABLE `courses` (
  `course_ID` int(11) NOT NULL,
  `course_college` text COLLATE utf8mb4_bin NOT NULL,
  `course_title` text COLLATE utf8mb4_bin NOT NULL,
  `course_description` text COLLATE utf8mb4_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `courses`
--

INSERT INTO `courses` (`course_ID`, `course_college`, `course_title`, `course_description`) VALUES
(1, 'College of Computer Studies', 'BSCS', 'Bachelor of Science in Computer Science '),
(2, 'College of Computer Studies', 'BSIT', 'Bachelor of Science in Information Technology'),
(3, 'College of Computer Studies', 'BSIS', 'Bachelor of Science in Information Systems'),
(4, 'College of Engineering', 'BSCpE', 'Bachelor of Science in Computer Engineering'),
(5, 'College of Engineering', 'BSEE', 'Bachelor of Science in  Electronics Engineering'),
(6, 'College of Engineering', 'BSIE', 'Bachelor of Science in Industrial Engineering'),
(7, 'College of Business Administration', 'BSBA', 'Bachelor of Science in Business Administration'),
(8, 'College of Business Administration', 'BSA', 'Bachelor of Science in Accountancy'),
(9, 'College of Arts and Sciences', 'BAENGL', 'Bachelor of Arts in English'),
(10, 'College of Arts and Sciences', 'BAECON', 'Bachelor of Arts in Economics'),
(11, 'College of Arts and Sciences', 'BAP', 'Bachelor of Arts in Psychology'),
(12, 'College of Arts and Sciences', 'BAMC', 'Bachelor of Arts in Mass Communication '),
(13, 'College of Arts and Sciences', 'BAPS', 'Bachelor of Arts in Political Science');

-- --------------------------------------------------------

--
-- Table structure for table `industry`
--

CREATE TABLE `industry` (
  `industry_ID` int(11) NOT NULL,
  `industry_title` text COLLATE utf8mb4_bin
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `industry`
--

INSERT INTO `industry` (`industry_ID`, `industry_title`) VALUES
(1, 'Accounting'),
(2, 'Airlines/Aviation'),
(3, 'Alternative Dispute Resolution'),
(4, 'Alternative Medicine'),
(5, 'Animation'),
(6, 'Apparel & Fashion'),
(7, 'Architecture & Planning'),
(8, 'Arts & Crafts'),
(9, 'Automotive'),
(10, 'Aviation & Aerospace'),
(11, 'Banking'),
(12, 'Biotechnology'),
(13, 'Broadcast Media'),
(14, 'Building Materials'),
(15, 'Business Supplies & Equipment'),
(16, 'Capital Markets'),
(17, 'Chemicals'),
(18, 'Civic & Social Organization'),
(19, 'Civil Engineering'),
(20, 'Commercial Real Estate'),
(21, 'Computer & Network Security'),
(22, 'Computer Games'),
(23, 'Computer Hardware'),
(24, 'Computer Networking'),
(25, 'Computer Software'),
(26, 'Construction'),
(27, 'Consumer Electronics'),
(28, 'Consumer Goods'),
(29, 'Consumer Services'),
(30, 'Cosmetics'),
(31, 'Dairy'),
(32, 'Defense & Space'),
(33, 'Design'),
(34, 'E-learning'),
(35, 'Education Management'),
(36, 'Electrical & Electronic Manufacturing'),
(37, 'Entertainment'),
(38, 'Environmental Services'),
(39, 'Events Services'),
(40, 'Executive Office'),
(41, 'Facilities Services'),
(42, 'Farming'),
(43, 'Financial Services'),
(44, 'Fine Art'),
(45, 'Fishery'),
(46, 'Food & Beverages'),
(47, 'Food Production'),
(48, 'Fundraising'),
(49, 'Furniture'),
(50, 'Glass, Ceramics & Concrete'),
(51, 'Government Administration'),
(52, 'Government Relations'),
(53, 'Graphic Design'),
(54, 'Health, Wellness & Fitness'),
(55, 'Higher Education'),
(56, 'Hospital & Health Care'),
(57, 'Hospitality'),
(58, 'Human Resources'),
(59, 'Import & Export'),
(60, 'Individual & Family Services'),
(61, 'Industrial Automation'),
(62, 'Information Services'),
(63, 'Information Technology & Services'),
(64, 'Insurance'),
(65, 'International Affairs'),
(66, 'International Trade & Development'),
(67, 'Internet'),
(68, 'Investment Banking'),
(69, 'Investment Management'),
(70, 'Judiciary'),
(71, 'Law Enforcement'),
(72, 'Law Practice'),
(73, 'Legal Services'),
(74, 'Legislative Office'),
(75, 'Leisure, Travel & Tourism'),
(76, 'Libraries'),
(77, 'Logistics & Supply Chain'),
(78, 'Luxury Goods & Jewelry'),
(79, 'Machinery'),
(80, 'Management Consulting'),
(81, 'Maritime'),
(82, 'Market Research'),
(83, 'Marketing & Advertising'),
(84, 'Mechanical Or Industrial Engineering'),
(85, 'Media Production'),
(86, 'Medical Device'),
(87, 'Medical Practice'),
(88, 'Mental Health Care'),
(89, 'Military'),
(90, 'Mining & Metals'),
(91, 'Motion Pictures & Film'),
(92, 'Museums & Institutions'),
(93, 'Music'),
(94, 'Nanotechnology'),
(95, 'Newspapers'),
(96, 'Non-profit Organization Management'),
(97, 'Oil & Energy'),
(98, 'Online Media'),
(99, 'Outsourcing/Offshoring'),
(100, 'Package/Freight Delivery'),
(101, 'Packaging & Containers'),
(102, 'Paper & Forest Products'),
(103, 'Performing Arts'),
(104, 'Pharmaceuticals'),
(105, 'Philanthropy'),
(106, 'Photography'),
(107, 'Plastics'),
(108, 'Political Organization'),
(109, 'Primary/Secondary Education'),
(110, 'Printing'),
(111, 'Professional Training & Coaching'),
(112, 'Program Development'),
(113, 'Public Policy'),
(114, 'Public Relations & Communications'),
(115, 'Public Safety'),
(116, 'Publishing'),
(117, 'Railroad Manufacture'),
(118, 'Ranching'),
(119, 'Real Estate'),
(120, 'Recreational Facilities & Services'),
(121, 'Religious Institutions'),
(122, 'Renewables & Environment'),
(123, 'Research'),
(124, 'Restaurants'),
(125, 'Retail'),
(126, 'Security & Investigations'),
(127, 'Semiconductors'),
(128, 'Shipbuilding'),
(129, 'Sporting Goods'),
(130, 'Sports'),
(131, 'Staffing & Recruiting'),
(132, 'Supermarkets'),
(133, 'Telecommunications'),
(134, 'Textiles'),
(135, 'Think Tanks'),
(136, 'Tobacco'),
(137, 'Translation & Localization'),
(138, 'Transportation/Trucking/Railroad'),
(139, 'Utilities'),
(140, 'Venture Capital & Private Equity'),
(141, 'Veterinary'),
(142, 'Warehousing'),
(143, 'Wholesale'),
(144, 'Wine & Spirits'),
(145, 'Wireless'),
(146, 'Writing & Editing');

-- --------------------------------------------------------

--
-- Table structure for table `login`
--

CREATE TABLE `login` (
  `login_ID` int(11) NOT NULL,
  `login_user_ID` text COLLATE utf8mb4_bin NOT NULL,
  `login_password` text COLLATE utf8mb4_bin NOT NULL,
  `login_level` text COLLATE utf8mb4_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `login`
--

INSERT INTO `login` (`login_ID`, `login_user_ID`, `login_password`, `login_level`) VALUES
(1, '1', 'admin', '1'),
(10, '123', '123', '2'),
(11, '1234', '1234', '2'),
(12, '12345', '12345', '2');

-- --------------------------------------------------------

--
-- Table structure for table `major`
--

CREATE TABLE `major` (
  `major_ID` int(11) NOT NULL,
  `major_course_ID` text COLLATE utf8mb4_bin NOT NULL,
  `major_name` text COLLATE utf8mb4_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `major`
--

INSERT INTO `major` (`major_ID`, `major_course_ID`, `major_name`) VALUES
(1, '7', 'Financial Management'),
(2, '7', 'Human Resource Development Management'),
(3, '7', 'Marketing'),
(4, '2', 'Computer Animation'),
(5, '2', 'Multimedia Arts'),
(6, '2', 'Digital Arts'),
(7, '2', 'Service Management'),
(8, '2', 'Business Analytics'),
(9, '2', 'Network Administration Plan'),
(10, '2', 'Web Application Development'),
(11, '2', 'Game Development'),
(12, '2', 'Mobile Applications'),
(13, '2', 'Robotics'),
(14, '1', 'Systems and Software Egineering'),
(15, '1', 'Software Development Plan'),
(16, '1', 'Game Development'),
(17, '1', 'Mobile Applications Development'),
(18, '1', 'Service Management'),
(19, '1', 'Business Analytics'),
(20, '1', 'Digital Arts'),
(21, '1', 'Computer Animation'),
(22, '1', 'Multimedia Arts'),
(23, '1', 'Network Administration'),
(24, '1', 'Robotics'),
(25, '1', 'Artificial Intelligence'),
(26, '1', 'Web Application'),
(27, '1', 'Human Computer Interaction'),
(28, '1', 'Data Science');

-- --------------------------------------------------------

--
-- Table structure for table `notification`
--

CREATE TABLE `notification` (
  `notification_ID` int(11) NOT NULL,
  `notification_recieve_ID` text COLLATE utf8mb4_bin,
  `notification_sender_ID` text COLLATE utf8mb4_bin,
  `notification_unread` text COLLATE utf8mb4_bin,
  `notification_type` text COLLATE utf8mb4_bin,
  `notification_param` text COLLATE utf8mb4_bin,
  `notification_type_ID` text COLLATE utf8mb4_bin,
  `notification_datetime` text COLLATE utf8mb4_bin
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- Table structure for table `salary`
--

CREATE TABLE `salary` (
  `salary_ID` int(11) NOT NULL,
  `salary_range` text COLLATE utf8mb4_bin
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `salary`
--

INSERT INTO `salary` (`salary_ID`, `salary_range`) VALUES
(1, '< 10,000'),
(2, '10,001 - 20,000'),
(3, '20,001 - 30,000'),
(4, '30,001 - 40,000'),
(5, '40,001 - 50,000'),
(6, '> 50,000');

-- --------------------------------------------------------

--
-- Table structure for table `settings`
--

CREATE TABLE `settings` (
  `settings_id` int(11) NOT NULL,
  `settings_user_ID` text COLLATE utf8mb4_bin,
  `settings_type` text COLLATE utf8mb4_bin NOT NULL,
  `settings_description` text COLLATE utf8mb4_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `settings`
--

INSERT INTO `settings` (`settings_id`, `settings_user_ID`, `settings_type`, `settings_description`) VALUES
(1, '1', 'theme', '3'),
(2, '1', 'sms', 'off'),
(3, '1', 'email', 'off'),
(4, '123', 'theme', '2'),
(5, '123', 'sms', 'off'),
(6, '123', 'email', 'off'),
(13, '1234', 'theme', '1'),
(14, '1234', 'sms', 'on'),
(15, '1234', 'email', 'off'),
(16, '1234', 'theme', '1'),
(17, '1234', 'sms', 'on'),
(18, '1234', 'email', 'off'),
(19, '12345', 'theme', '1'),
(20, '12345', 'sms', 'on'),
(21, '12345', 'email', 'off');

-- --------------------------------------------------------

--
-- Table structure for table `theme`
--

CREATE TABLE `theme` (
  `theme_ID` int(11) NOT NULL,
  `theme_color` text COLLATE utf8mb4_bin NOT NULL,
  `theme_description` text COLLATE utf8mb4_bin NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `theme`
--

INSERT INTO `theme` (`theme_ID`, `theme_color`, `theme_description`) VALUES
(1, 'Red', 'danger'),
(2, 'Blue', 'info'),
(3, 'Orange', 'warning'),
(4, 'Violet', 'primary'),
(5, 'Green', 'success');

-- --------------------------------------------------------

--
-- Table structure for table `time`
--

CREATE TABLE `time` (
  `time_ID` int(11) NOT NULL,
  `time_hour` text COLLATE utf8mb4_bin,
  `time_description` text COLLATE utf8mb4_bin
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `time`
--

INSERT INTO `time` (`time_ID`, `time_hour`, `time_description`) VALUES
(8, '8:00:00', '8 AM'),
(9, '9:00:00', '9 AM'),
(10, '10:00:00', '10 AM'),
(11, '11:00:00', '11 AM'),
(13, '13:00:00', '1 PM'),
(14, '14:00:00', '2 PM'),
(15, '15:00:00', '3 PM'),
(16, '16:00:00', '4 PM');

-- --------------------------------------------------------

--
-- Table structure for table `work`
--

CREATE TABLE `work` (
  `work_ID` int(11) NOT NULL,
  `work_company_name` text COLLATE utf8mb4_bin,
  `work_company_address` text COLLATE utf8mb4_bin,
  `work_industry` text COLLATE utf8mb4_bin,
  `work_alumni_student_ID` text COLLATE utf8mb4_bin,
  `work_alumni_position` text COLLATE utf8mb4_bin,
  `work_alumni_salary_range` text COLLATE utf8mb4_bin,
  `work_alumni_start` text COLLATE utf8mb4_bin,
  `work_alumni_end` text COLLATE utf8mb4_bin
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

--
-- Dumping data for table `work`
--

INSERT INTO `work` (`work_ID`, `work_company_name`, `work_company_address`, `work_industry`, `work_alumni_student_ID`, `work_alumni_position`, `work_alumni_salary_range`, `work_alumni_start`, `work_alumni_end`) VALUES
(1, 'Sample Company', 'Sample Address', '5', '123', 'Sample Position', '5', '2019-01', ''),
(2, 'Sample Company', 'Sample Address', '28', '123', 'Sample Position', '5', '2019-01', ''),
(3, 'Sample Company', 'Sample Address', '4', '1234', 'Sample Position', '5', '2018-11', '2019-02');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`admin_ID`);

--
-- Indexes for table `alumni`
--
ALTER TABLE `alumni`
  ADD PRIMARY KEY (`alumni_ID`);

--
-- Indexes for table `announcement`
--
ALTER TABLE `announcement`
  ADD PRIMARY KEY (`announcement_ID`);

--
-- Indexes for table `appointment`
--
ALTER TABLE `appointment`
  ADD PRIMARY KEY (`appointment_ID`);

--
-- Indexes for table `ci_sessions`
--
ALTER TABLE `ci_sessions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `ci_sessions_timestamp` (`timestamp`);

--
-- Indexes for table `courses`
--
ALTER TABLE `courses`
  ADD PRIMARY KEY (`course_ID`);

--
-- Indexes for table `industry`
--
ALTER TABLE `industry`
  ADD PRIMARY KEY (`industry_ID`);

--
-- Indexes for table `login`
--
ALTER TABLE `login`
  ADD PRIMARY KEY (`login_ID`);

--
-- Indexes for table `major`
--
ALTER TABLE `major`
  ADD PRIMARY KEY (`major_ID`);

--
-- Indexes for table `notification`
--
ALTER TABLE `notification`
  ADD PRIMARY KEY (`notification_ID`);

--
-- Indexes for table `salary`
--
ALTER TABLE `salary`
  ADD PRIMARY KEY (`salary_ID`);

--
-- Indexes for table `settings`
--
ALTER TABLE `settings`
  ADD PRIMARY KEY (`settings_id`);

--
-- Indexes for table `theme`
--
ALTER TABLE `theme`
  ADD PRIMARY KEY (`theme_ID`);

--
-- Indexes for table `time`
--
ALTER TABLE `time`
  ADD PRIMARY KEY (`time_ID`);

--
-- Indexes for table `work`
--
ALTER TABLE `work`
  ADD PRIMARY KEY (`work_ID`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admin`
--
ALTER TABLE `admin`
  MODIFY `admin_ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
--
-- AUTO_INCREMENT for table `alumni`
--
ALTER TABLE `alumni`
  MODIFY `alumni_ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;
--
-- AUTO_INCREMENT for table `announcement`
--
ALTER TABLE `announcement`
  MODIFY `announcement_ID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `appointment`
--
ALTER TABLE `appointment`
  MODIFY `appointment_ID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `courses`
--
ALTER TABLE `courses`
  MODIFY `course_ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=14;
--
-- AUTO_INCREMENT for table `industry`
--
ALTER TABLE `industry`
  MODIFY `industry_ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=147;
--
-- AUTO_INCREMENT for table `login`
--
ALTER TABLE `login`
  MODIFY `login_ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
--
-- AUTO_INCREMENT for table `major`
--
ALTER TABLE `major`
  MODIFY `major_ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=29;
--
-- AUTO_INCREMENT for table `notification`
--
ALTER TABLE `notification`
  MODIFY `notification_ID` int(11) NOT NULL AUTO_INCREMENT;
--
-- AUTO_INCREMENT for table `salary`
--
ALTER TABLE `salary`
  MODIFY `salary_ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;
--
-- AUTO_INCREMENT for table `settings`
--
ALTER TABLE `settings`
  MODIFY `settings_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=22;
--
-- AUTO_INCREMENT for table `theme`
--
ALTER TABLE `theme`
  MODIFY `theme_ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
--
-- AUTO_INCREMENT for table `time`
--
ALTER TABLE `time`
  MODIFY `time_ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;
--
-- AUTO_INCREMENT for table `work`
--
ALTER TABLE `work`
  MODIFY `work_ID` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;