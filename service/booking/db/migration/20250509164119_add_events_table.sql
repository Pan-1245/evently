-- Add new schema named "booking"
CREATE SCHEMA "booking";
-- Create "event" table
CREATE TABLE "booking"."event" (
  "id" uuid NOT NULL DEFAULT gen_random_uuid(),
  "title" text NOT NULL,
  "description" text NULL,
  "start_time" timestamptz NOT NULL,
  "end_time" timestamptz NOT NULL,
  "location" text NOT NULL,
  "is_active" boolean NOT NULL,
  "created_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NULL,
  "organizer_id" text NOT NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_booking_event_deleted_at" to table: "event"
CREATE INDEX "idx_booking_event_deleted_at" ON "booking"."event" ("deleted_at");
