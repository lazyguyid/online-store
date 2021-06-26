/*
 Navicat Premium Data Transfer

 Source Server         : local-shar
 Source Server Type    : PostgreSQL
 Source Server Version : 100001
 Source Host           : localhost:5432
 Source Catalog        : online_store
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 100001
 File Encoding         : 65001

 Date: 26/06/2021 22:58:54
*/


-- Create SEQUENCES
CREATE SEQUENCE orderdetailid_seq;
CREATE SEQUENCE userid_seq;
CREATE SEQUENCE productid_seq;
CREATE SEQUENCE orderid_seq;

-- ----------------------------
-- Table structure for order_details
-- ----------------------------
DROP TABLE IF EXISTS "order_details";
CREATE TABLE "order_details" (
  "id" int4 NOT NULL DEFAULT nextval('orderdetailid_seq'::regclass),
  "product_id" int4 NOT NULL,
  "price" float4 NOT NULL,
  "qty" int4 NOT NULL
)
;
ALTER TABLE "order_details" OWNER TO "root";

-- ----------------------------
-- Records of order_details
-- ----------------------------
BEGIN;
INSERT INTO "order_details" VALUES (2, 1, 1.1e+07, 10);
COMMIT;

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS "orders";
CREATE TABLE "orders" (
  "id" int4 NOT NULL DEFAULT nextval('orderid_seq'::regclass),
  "user_id" int4,
  "total" float4
)
;
ALTER TABLE "orders" OWNER TO "root";

-- ----------------------------
-- Records of orders
-- ----------------------------
BEGIN;
INSERT INTO "orders" VALUES (2, 1, 0);
COMMIT;

-- ----------------------------
-- Table structure for products
-- ----------------------------
DROP TABLE IF EXISTS "products";
CREATE TABLE "products" (
  "id" int4 NOT NULL DEFAULT nextval('productid_seq'::regclass),
  "name" varchar(50) COLLATE "pg_catalog"."default",
  "qty" int4,
  "created" timestamptz(0) NOT NULL,
  "updated" timestamptz(0) NOT NULL,
  "price" float4
)
;
ALTER TABLE "products" OWNER TO "root";

-- ----------------------------
-- Records of products
-- ----------------------------
BEGIN;
INSERT INTO "products" VALUES (2, 'Samsung S20', 5, '2021-06-26 18:26:47+07', '2021-06-26 18:26:52+07', 1.2e+07);
INSERT INTO "products" VALUES (3, 'Xiaomi', 2, '2021-06-26 18:27:11+07', '2021-06-26 18:27:18+07', 8e+06);
INSERT INTO "products" VALUES (1, 'Iphone x', 10, '2021-06-26 18:26:23+07', '2021-06-26 18:26:28+07', 1.1e+07);
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS "users";
CREATE TABLE "users" (
  "id" int4 NOT NULL DEFAULT nextval('userid_seq'::regclass),
  "name" varchar(50) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "users" OWNER TO "root";

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO "users" VALUES (1, 'buyer A');
INSERT INTO "users" VALUES (2, 'buyer B');
INSERT INTO "users" VALUES (3, 'buyer C');
COMMIT;

-- ----------------------------
-- Primary Key structure for table order_details
-- ----------------------------
ALTER TABLE "order_details" ADD CONSTRAINT "order_detail_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table orders
-- ----------------------------
ALTER TABLE "orders" ADD CONSTRAINT "order_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table products
-- ----------------------------
ALTER TABLE "products" ADD CONSTRAINT "product_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table users
-- ----------------------------
ALTER TABLE "users" ADD CONSTRAINT "user_pkey" PRIMARY KEY ("id");
