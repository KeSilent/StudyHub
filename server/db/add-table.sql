CREATE TABLE
    `edu_organization` (
        `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '组织ID',
        `name` varchar(255) NOT NULL COMMENT '组织名称',
        `brief` varchar(255) DEFAULT NULL COMMENT '组织简介',
        `details` longtext DEFAULT NULL COMMENT '组织详情（HTML字符串）',
        `contact_info` varchar(255) DEFAULT NULL COMMENT '联系方式',
        `address` varchar(255) DEFAULT NULL COMMENT '地址',
        `created_at` datetime NOT NULL COMMENT '组织创建时间',
        `updated_at` datetime NOT NULL COMMENT '组织更新时间',
        PRIMARY KEY (`id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

ALTER TABLE `sys_users`
ADD
    COLUMN `edu_organization_id` int(11) DEFAULT NULL COMMENT '组织ID';

ALTER TABLE `sys_users`
ADD
    COLUMN `wx_openid` VARCHAR(255) DEFAULT NULL COMMENT '微信ID';

ALTER TABLE `sys_users`
ADD
    COLUMN `wx_session_key` VARCHAR(255) DEFAULT NULL COMMENT '微信session_key';

CREATE TABLE
    `edu_course` (
        `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '课程ID',
        `organization_id` int(11) NOT NULL COMMENT '组织ID',
        `admin_id` int(11) NOT NULL COMMENT '管理员ID',
        `course_name` varchar(255) NOT NULL COMMENT '课程名称',
        `description` varchar(255) DEFAULT NULL COMMENT '课程描述',
        `image_url` varchar(255) DEFAULT NULL COMMENT '课程展示图URL',
        `created_at` datetime NOT NULL COMMENT '课程创建时间',
        `updated_at` datetime NOT NULL COMMENT '课程更新时间',
        PRIMARY KEY (`id`),
        KEY `admin_id` (`admin_id`),
        KEY `organization_id` (`organization_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE
    `edu_enrollment` (
        `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '报名ID',
        `user_id` int(11) NOT NULL COMMENT '用户ID',
        `course_id` int(11) NOT NULL COMMENT '课程ID',
        `total_sessions` int(11) NOT NULL COMMENT '总课时数',
        `remaining_sessions` int(11) NOT NULL COMMENT '剩余课时数',
        `created_at` datetime NOT NULL COMMENT '报名创建时间',
        `updated_at` datetime NOT NULL COMMENT '报名更新时间',
        PRIMARY KEY (`id`),
        KEY `user_id` (`user_id`),
        KEY `course_id` (`course_id`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE
    `edu_class_session` (
        `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '课时ID',
        `enrollment_id` int(11) NOT NULL COMMENT '报名ID',
        `action` ENUM('add', 'subtract') NOT NULL COMMENT '操作类型（增加或扣除）',
        `reason` varchar(255) DEFAULT NULL COMMENT '操作原因',
        `num_sessions` int(11) NOT NULL COMMENT '课时数量',
        `created_at` datetime NOT NULL COMMENT '课时操作创建时间',
        `updated_at` datetime NOT NULL COMMENT '课时操作更新时间',
        PRIMARY KEY (id),
        KEY enrollment_id (enrollment_id)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;