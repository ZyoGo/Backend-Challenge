CREATE TABLE
    "products" (
        "id" varchar(26) PRIMARY KEY NOT NULL,
        "category_id" varchar(26) NOT NULL,
        "name" varchar(255) NOT NULL,
        "description" text,
        "stock" integer NOT NULL,
        "price" numeric(10, 2) NOT NULL,
        "created_at" timestamp NOT NULL DEFAULT (now ()),
        "updated_at" timestamp NOT NULL DEFAULT (now ()),
        "deleted_at" timestamp
    );

CREATE TABLE
    "categories" (
        "id" varchar(26) PRIMARY KEY NOT NULL,
        "name" varchar(100) NOT NULL,
        "created_at" timestamp NOT NULL DEFAULT (now ()),
        "updated_at" timestamp NOT NULL DEFAULT (now ()),
        "deleted_at" timestamp
    );

CREATE TABLE
    "cart_items" (
        "id" varchar(26) PRIMARY KEY NOT NULL,
        "cart_id" varchar(26) NOT NULL,
        "product_id" varchar(26) NOT NULL,
        "quantity" integer NOT NULL,
        "created_at" timestamp NOT NULL DEFAULT (now ()),
        "updated_at" timestamp NOT NULL DEFAULT (now ()),
        "deleted_at" timestamp
    );

CREATE TABLE
    "carts" (
        "id" varchar(26) PRIMARY KEY NOT NULL,
        "user_id" varchar(26) NOT NULL,
        "created_at" timestamp NOT NULL DEFAULT (now ()),
        "updated_at" timestamp NOT NULL DEFAULT (now ()),
        "deleted_at" timestamp
    );

CREATE TABLE
    "users" (
        "id" varchar(26) PRIMARY KEY NOT NULL,
        "username" varchar(255) NOT NULL,
        "email" varchar(255) NOT NULL,
        "password" text NOT NULL,
        "address" text NOT NULL,
        "phone_number" varchar(155) NOT NULL,
        "created_at" timestamp NOT NULL DEFAULT (now ()),
        "updated_at" timestamp NOT NULL DEFAULT (now ()),
        "deleted_at" timestamp
    );

CREATE TABLE
    "orders" (
        "id" varchar(26) PRIMARY KEY NOT NULL,
        "user_id" varchar(26) NOT NULL,
        "payment_status" varchar(30),
        "created_at" timestamp NOT NULL DEFAULT (now ()),
        "updated_at" timestamp NOT NULL DEFAULT (now ()),
        "deleted_at" timestamp
    );

CREATE TABLE
    "order_items" (
        "id" varchar(26) PRIMARY KEY NOT NULL,
        "order_id" varchar(26) NOT NULL,
        "product_id" varchar(26) NOT NULL,
        "quantity" smallint NOT NULL,
        "price" numeric(10, 2) NOT NULL,
        "created_at" timestamp NOT NULL DEFAULT (now ()),
        "updated_at" timestamp NOT NULL DEFAULT (now ()),
        "deleted_at" timestamp
    );

ALTER TABLE "products" ADD CONSTRAINT "products_category_id_fkey" FOREIGN KEY ("category_id") REFERENCES "categories" ("id");

ALTER TABLE "cart_items" ADD CONSTRAINT "cart_items_cart_id_fkey" FOREIGN KEY ("cart_id") REFERENCES "carts" ("id");

ALTER TABLE "cart_items" ADD CONSTRAINT "cart_items_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "carts" ADD CONSTRAINT "carts_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "order_items" ADD CONSTRAINT "order_items_order_id_fkey" FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "order_items" ADD CONSTRAINT "order_items_product_id_fkey" FOREIGN KEY ("product_id") REFERENCES "products" ("id");

ALTER TABLE "orders" ADD CONSTRAINT "orders_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id");