-- Create "access_tokens" table
CREATE TABLE `access_tokens` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime NULL,
  `token` varchar(100) NOT NULL,
  `usage` varchar(50) NOT NULL,
  `email` varchar(255) NOT NULL DEFAULT "",
  `expired_at` datetime NULL,
  `last_used_at` datetime NULL,
  `user_info` json NULL,
  PRIMARY KEY (`id`),
  INDEX `accesstoken_email` (`email`),
  UNIQUE INDEX `token` (`token`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "cache_locks" table
CREATE TABLE `cache_locks` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `key` varchar(255) NOT NULL,
  `owner` varchar(255) NOT NULL,
  `expired_at` datetime NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `key` (`key`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "db_cache" table
CREATE TABLE `db_cache` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `key` varchar(255) NOT NULL,
  `value` longtext NULL,
  `expired_at` datetime NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `key` (`key`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "namespaces" table
CREATE TABLE `namespaces` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime NULL,
  `name` varchar(100) NOT NULL COLLATE utf8mb4_general_ci,
  `image_pull_secrets` json NOT NULL,
  `description` text NULL,
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "repos" table
CREATE TABLE `repos` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime NULL,
  `name` varchar(255) NOT NULL COLLATE utf8mb4_general_ci,
  `default_branch` varchar(255) NULL,
  `git_project_name` varchar(255) NULL,
  `git_project_id` int NULL,
  `enabled` bool NOT NULL DEFAULT 0,
  `need_git_repo` bool NOT NULL DEFAULT 0,
  `mars_config` json NULL,
  `description` varchar(255) NOT NULL DEFAULT "",
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "projects" table
CREATE TABLE `projects` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime NULL,
  `name` varchar(100) NOT NULL DEFAULT "",
  `git_project_id` bigint NULL,
  `git_branch` varchar(255) NULL,
  `git_commit` varchar(255) NULL,
  `config` longtext NULL,
  `creator` varchar(255) NOT NULL,
  `override_values` longtext NULL,
  `docker_image` json NULL,
  `pod_selectors` json NULL,
  `atomic` bool NOT NULL DEFAULT 0,
  `deploy_status` int NOT NULL DEFAULT 0,
  `env_values` json NULL,
  `extra_values` json NULL,
  `final_extra_values` json NULL,
  `version` bigint NOT NULL DEFAULT 1,
  `config_type` varchar(255) NULL,
  `manifest` json NULL,
  `git_commit_web_url` varchar(255) NOT NULL DEFAULT "",
  `git_commit_title` varchar(255) NOT NULL DEFAULT "",
  `git_commit_author` varchar(255) NOT NULL DEFAULT "",
  `git_commit_date` datetime NULL,
  `namespace_id` bigint NULL,
  `repo_id` bigint NULL,
  PRIMARY KEY (`id`),
  INDEX `project_git_project_id` (`git_project_id`),
  INDEX `projects_namespaces_projects` (`namespace_id`),
  INDEX `projects_repos_projects` (`repo_id`),
  CONSTRAINT `projects_namespaces_projects` FOREIGN KEY (`namespace_id`) REFERENCES `namespaces` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL,
  CONSTRAINT `projects_repos_projects` FOREIGN KEY (`repo_id`) REFERENCES `repos` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "changelogs" table
CREATE TABLE `changelogs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime NULL,
  `version` bigint NOT NULL DEFAULT 1,
  `username` varchar(100) NOT NULL,
  `config` varchar(255) NULL,
  `git_branch` varchar(255) NULL,
  `git_commit` varchar(255) NULL,
  `docker_image` json NULL,
  `env_values` json NULL,
  `extra_values` json NULL,
  `final_extra_values` json NULL,
  `git_commit_web_url` varchar(255) NULL,
  `git_commit_title` varchar(255) NULL,
  `git_commit_author` varchar(255) NULL,
  `git_commit_date` datetime NULL,
  `config_changed` bool NOT NULL DEFAULT 0,
  `project_id` bigint NULL,
  PRIMARY KEY (`id`),
  INDEX `changelog_project_id_config_changed_deleted_at_version` (`project_id`, `config_changed`, `deleted_at`, `version`),
  CONSTRAINT `changelogs_projects_changelogs` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "files" table
CREATE TABLE `files` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime NULL,
  `upload_type` varchar(100) NOT NULL DEFAULT "local",
  `path` varchar(255) NOT NULL,
  `size` int NOT NULL DEFAULT 0,
  `username` varchar(255) NOT NULL DEFAULT "",
  `namespace` varchar(100) NOT NULL DEFAULT "",
  `pod` varchar(100) NOT NULL DEFAULT "",
  `container` varchar(100) NOT NULL DEFAULT "",
  `container_path` varchar(255) NOT NULL DEFAULT "",
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "events" table
CREATE TABLE `events` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `deleted_at` datetime NULL,
  `action` int NOT NULL DEFAULT 0,
  `username` varchar(255) NOT NULL DEFAULT "",
  `message` varchar(255) NOT NULL DEFAULT "",
  `old` longtext NULL,
  `new` longtext NULL,
  `has_diff` bool NOT NULL DEFAULT 0,
  `duration` varchar(255) NOT NULL DEFAULT "",
  `file_id` bigint NULL,
  PRIMARY KEY (`id`),
  INDEX `event_action` (`action`),
  INDEX `event_username_created_at` (`username`, `created_at`),
  INDEX `events_files_events` (`file_id`),
  CONSTRAINT `events_files_events` FOREIGN KEY (`file_id`) REFERENCES `files` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "favorites" table
CREATE TABLE `favorites` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `namespace_id` bigint NULL,
  PRIMARY KEY (`id`),
  INDEX `favorites_namespaces_favorites` (`namespace_id`),
  CONSTRAINT `favorites_namespaces_favorites` FOREIGN KEY (`namespace_id`) REFERENCES `namespaces` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
