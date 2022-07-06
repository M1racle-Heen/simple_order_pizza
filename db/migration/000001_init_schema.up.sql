CREATE TABLE "customers" (
  "id" bigserial PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone" bigint NOT NULL,
  "address" varchar NOT NULL
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "customer_id" bigint NOT NULL,
  "status" varchar NOT NULL DEFAULT 'Hold',
  "order_time" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "pizza" (
  "id" bigserial PRIMARY KEY,
  "order_id" bigint NOT NULL,
  "price" bigint NOT NULL DEFAULT '900',
  "pizza_type" varchar NOT NULL DEFAULT 'Salyami',
  "pizza_quant" bigint NOT NULL
);

CREATE TABLE "payment" (
  "id" bigserial PRIMARY KEY,
  "pizza_id" bigint NOT NULL,
  "customer_id" bigint NOT NULL,
  "payment_status" varchar NOT NULL DEFAULT 'Not Paid',
  "bill" bigint NOT NULL
);

CREATE INDEX ON "orders" ("customer_id");

CREATE INDEX ON "pizza" ("order_id");

CREATE INDEX ON "payment" ("pizza_id");

CREATE INDEX ON "payment" ("customer_id");

CREATE INDEX ON "payment" ("pizza_id", "customer_id");

ALTER TABLE "orders" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "pizza" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("id");

ALTER TABLE "payment" ADD FOREIGN KEY ("pizza_id") REFERENCES "pizza" ("id");

ALTER TABLE "payment" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");
