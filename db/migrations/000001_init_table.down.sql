ALTER TABLE "order_items"
DROP CONSTRAINT IF EXISTS "order_items_order_id_fkey";

ALTER TABLE "order_items"
DROP CONSTRAINT IF EXISTS "order_items_product_id_fkey";

ALTER TABLE "orders"
DROP CONSTRAINT IF EXISTS "orders_user_id_fkey";

ALTER TABLE "products"
DROP CONSTRAINT IF EXISTS "products_category_id_fkey";

ALTER TABLE "cart_items"
DROP CONSTRAINT IF EXISTS "cart_items_cart_id_fkey";

ALTER TABLE "cart_items"
DROP CONSTRAINT IF EXISTS "cart_items_product_id_fkey";

ALTER TABLE "carts"
DROP CONSTRAINT IF EXISTS "carts_user_id_fkey";

DROP TABLE IF EXISTS "order_items";

DROP TABLE IF EXISTS "orders";

DROP TABLE IF EXISTS "users";

DROP TABLE IF EXISTS "carts";

DROP TABLE IF EXISTS "cart_items";

DROP TABLE IF EXISTS "products";

DROP TABLE IF EXISTS "categories";