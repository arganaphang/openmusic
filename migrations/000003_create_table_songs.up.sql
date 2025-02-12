CREATE TABLE IF NOT EXISTS "songs" (
    "id" VARCHAR PRIMARY KEY,
    "title" VARCHAR NOT NULL,
    "year" SMALLINT NOT NULL,
    "genre" VARCHAR NOT NULL,
    "performer" VARCHAR NOT NULL,
    "duration" SMALLINT NOT NULL,
    "album_id" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY("album_id") REFERENCES "albums"("id")
);
