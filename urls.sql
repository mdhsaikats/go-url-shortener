DROP TABLE IF EXISTS "public"."urls";
-- Table Definition
CREATE TABLE "public"."urls" (
    "id" uuid NOT NULL DEFAULT gen_random_uuid(),
    "short_code" varchar(10) NOT NULL,
    "original_url" text NOT NULL,
    "created_at" timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id")
);


-- Indices
CREATE UNIQUE INDEX urls_short_code_key ON public.urls USING btree (short_code);


