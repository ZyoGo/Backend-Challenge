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
        "amount" numeric(10, 2) NOT NULL,
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

-- Add seeder for user: 
-- email: johndoe@gmail.com
-- password: Password123
INSERT INTO
    users (
        id,
        username,
        email,
        password,
        address,
        phone_number,
        created_at,
        updated_at
    )
VALUES
    (
        '01G65Z755AFWAKHE12NY0CQ9FH',
        'John Doe',
        'johndoe@gmail.com',
        '$2a$10$2GiBuRdqVaiyo2kNmspIt.14ktje..HyG3BiG1X8gqMvDBQKknhHC',
        'BTN',
        '08*****',
        now (),
        now ()
    );

-- Add seeder for category
INSERT INTO
    categories (id, name, created_at, updated_at)
VALUES
    (
        '01HWS9JHBNEZ286AF1AT58V3AD',
        'Keyboard',
        now (),
        now ()
    ),
    (
        '01HWS9M3GRQ3N4781W1MKRJGKB',
        'Mouse',
        now (),
        now ()
    ),
    (
        '01HWS9MTPZSK8YYV5BFKF96DZ4',
        'In Ear Monitor',
        now (),
        now ()
    );

-- Add seeder for products
INSERT INTO
    products (
        id,
        category_id,
        name,
        description,
        stock,
        price,
        created_at,
        updated_at
    )
VALUES
    (
        '01HWS9JHBNEZ286AF1AT58V3AD',
        '01HWS9JHBNEZ286AF1AT58V3AD',
        'Keychron K3 version 2 Hot-Swappable RGB Backlight Low Profile Switch - Brown Switch',
        'WHATS NEW ON K3 VERSION 2 ?',
        100,
        1650000,
        now (),
        now ()
    ),
    (
        '01HWSC8PVQ5MB4N7CKVXQ5CBA2',
        '01HWS9JHBNEZ286AF1AT58V3AD',
        'ROVER84 Lightyear Edition Wireless CNC Aluminium Mechanical Keyboard - Quark Matte',
        'WARNING: Jangan charge keyboard menggunakan stop kontak dan adapter charging HP. Sambungkan keyboard ke USB port laptop/PC untuk charging. Charging ke stop kontak/ adapter HP dapat mengakibatkan kerusakan motherboarddan akan menghanguskan warranty',
        100,
        1789000,
        now (),
        now ()
    ),
    (
        '01HWSC1J8M1TX79R12PRV4DXGD',
        '01HWS9M3GRQ3N4781W1MKRJGKB',
        'Logitech G PRO X Superlight 2 Wireless Gaming Mouse GPRO X PROX 2 - Black, Superlight 2',
        'Logitech G PRO X SUPERLIGHT 2 LIGHTSPEED WIreless Mouse Gaming,Sensor HERO 2, 32K DPI, for eSport',
        100,
        2199000,
        now (),
        now ()
    ),
    (
        '01HWSC3SA2P3PXSP31J6NVCJVY',
        '01HWS9MTPZSK8YYV5BFKF96DZ4',
        'Moondrop Starfield 2 / II 10mm Dynamic Driver in-Ear Monitor iem',
        'Moondrop is a brand we all are aware of. With their outstanding range of in-ear monitors, Moondrop has got a good reputation in the audiophile industry. Over the years, Moondrop has designed so many beautiful IEMs that are praised both for their sound performance as well as their exquisite craftsmanship. Today, Moondrop has released a successor to the classic Starfield IEMs that they released a few years back, introducing the all-new Moondrop Starfield 2. As a successor, the Moondrop Starfield 2 brings a lot of new updates to the pair that bring drastic improvements in its sound performance. The design is identical to the OG model, but the performance and the comfort are taken to an all-new level of awesomeness!!',
        100,
        1649500,
        now (),
        now ()
    );