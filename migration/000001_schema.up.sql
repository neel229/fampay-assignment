CREATE TABLE "videos"(
    "id" UUID NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "video_id" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NULL,
    "published_at" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL,
    "thumbnail_url" VARCHAR(255) NOT NULL
);
CREATE INDEX "videos_title_description_index" ON
    "videos"("title", "description");
ALTER TABLE
    "videos" ADD PRIMARY KEY("id");