CREATE DATABASE IF NOT EXISTS drifting;

USE drifting;

CREATE TABLE IF NOT EXISTS users(
    id BIGINT NOT NULL AUTO_INCREMENT,
    created_at DATETIME NULL,
    update_at DATETIME NULL,
    deleted_at DATETIME NULL,
    student_id BIGINT NOT NULL,
    name LONGTEXT NOT NULL,
    sex LONGTEXT NULL,
    avatar LONGTEXT NULL,
    PRIMARY KEY (id)
    )ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS drifting_notes(
    id BIGINT NOT NULL AUTO_INCREMENT,
    created_at DATETIME NULL,
    update_at DATETIME NULL,
    deleted_at DATETIME NULL,
    name LONGTEXT NOT NULL,
    contact LONGTEXT NULL,
    cover LONGTEXT NULL,
    owner_id BIGINT NOT NULL,
    )

CREATE TABLE IF NOT EXISTS drifting_pictures(
    id BIGINT NOT NULL AUTO_INCREMENT,
    created_at DATETIME NULL,
    update_at DATETIME NULL,
    deleted_at DATETIME NULL,
    name LONGTEXT NOT NULL,
    contact lONGTEXT NULL,
    cover lONGTEXT NULL,
    owener_id BIGINT NOT NULL,
    )