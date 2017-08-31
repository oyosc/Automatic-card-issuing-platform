-- ************************************************************
--
-- https://github.com/oyosc/Automatic-card-issuing-platform
--
-- Host: 127.0.0.1 (psql (PostgreSQL) 9.6.3)
-- Port:5432
-- Database: automatic
-- Generation Time: 2017-08-31 11:52:01 +0000
-- ************************************************************

-- Dump of table article_category
-- -----------------------------------------------------------------------

DROP TABLE IF EXISTS "user_role";

CREATE TABLE "user_role"(
	"id" int PRIMARY KEY NOT NULL,
	"user_id" int NOT NULL,
	"type" int NOT NULL DEFAULT 3,
	"updated_at" varchar(50) NOT NULL,
	"created_at" varchar(50) NOT NULL
);


INSERT INTO "user_role" ("id","user_id","type", "updated_at", "created_at") VALUES (1,1,1,'2017-08-31 13:40:54', '2017-08-31 13:40:54');

-- Dump of table user
-- ---------------------------------------------------------------------

DROP TABLE IF EXISTS "users";

CREATE TABLE "users"(
	"id" int PRIMARY KEY NOT NULL,
	"name" varchar(50) NOT NULL,
	"passwd" varchar(50) NOT NULL,
	"phone" varchar(50) NOT NULL,
	"email" varchar(50) NOT NULL,
	"status" int NOT NULL DEFAULT 0,
	"created_at" varchar(50) NOT NULL,
	"updated_at" varchar(50) NOT NULL
);


INSERT INTO "users" ("id","name", "passwd", "phone", "email", "status", "updated_at", "created_at") VALUES(1, 'admin', 'admin', 'null', 'chenlei@nicefilm.com', '1', '2017-08-31 13:40:54', '2017-08-31 13:40:54');

-- Dump of table assortment
-- --------------------------------------------------------------------

DROP TABLE IF EXISTS "assortment";

CREATE TABLE "assortment"(
	"id" int PRIMARY KEY NOT NULL,
	"name" varchar(50) NOT NULL,
	"created_at" varchar(50) NOT NULL,
	"updated_at" varchar(50) NOT NULL,
	"status" int NOT NULL DEFAULT 1,
	"user_id" int NOT NULL
);

-- Dump of table assortment
-- -------------------------------------------------------------------

DROP TABLE IF EXISTS "goods";

CREATE TABLE "goods" (
	"id" int PRIMARY KEY NOT NULL,
	"title" varchar(50) NOT NULL,
	"stock" int NOT NULL,
	"price" varchar(20) NOT NULL,
	"details" varchar(100) NOT NULL,
	"type" int NOT NULL DEFAULT 0,
	"assortment_id" int NOT NULL,
	"user_id" int NOT NULL,
	"created_at" varchar(50) NOT NULL,
	"updated_at" varchar(50) NOT NULL,
	"status" int NOT NULL DEFAULT 1
);

-- Dump of table cdkey
-- ---------------------------------------------------------------------

DROP TABLE IF EXISTS "cdkey";

CREATE TABLE "cdkey" (
	"id" int PRIMARY KEY NOT NULL,
	"cdk" varchar(100) NOT NULL,
	"goods_id" int NOT NULL,
	"user_id" int NOT NULL,
	"status" int NOT NULL DEFAULT 1,
	"created_at" varchar(50) NOT NULL,
	"updated_at" varchar(50) NOT NULL
);

-- Dump of table order
-- ----------------------------------------------------------------------

DROP TABLE IF EXISTS "order";

CREATE TABLE "order" (
	"id" int PRIMARY KEY NOT NULL,
	"good_name" varchar(100) NOT NULL,
	"count" int NOT NULL,
	"total_price" varchar(20) NOT NULL,
	"contact" varchar(100) NOT NULL,
	"goods_id" int NOT NULL,
	"user_id" int NOT NULL,
	"status" int NOT NULL,
	"order_number" varchar(100) NOT NULL,
	"created_at" varchar(50) NOT NULL,
	"success_at" varchar(50) NOT NULL
);













