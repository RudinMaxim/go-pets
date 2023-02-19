CREATE TABLE "cats" (
    "id" bigserial PRIMARY KEY NOT NULL,
    "name" varchar NOT NULL,
    "is_strip" boolean  NOT NULL DEFAULT false,
    "color" varchar
);

CREATE TABLE "dogs" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "breed" varchar NOT NULL,
    "color" varchar NOT NULL,
    "ear_length" int NOT NULL
);