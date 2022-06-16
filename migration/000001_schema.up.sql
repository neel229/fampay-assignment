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
ALTER TABLE 
    "videos" ADD COLUMN tsv tsvector;
UPDATE "videos" SET tsv =
    setweight(to_tsvector(title), 'A') ||
    setweight(to_tsvector(description), 'B');
CREATE INDEX ix_videos_tsv ON "videos" USING GIN(tsv);

CREATE FUNCTION videos_tsvector_trigger() RETURNS trigger as $$
begin
    new.tsv :=
    setweight(to_tsvector('english', new.title), 'A')
    || setweight(to_tsvector('english', new.description), 'B');
    return new;
end
$$ LANGUAGE plpgsql;

CREATE TRIGGER tsvectorupdate BEFORE INSERT OR UPDATE
    ON "videos" FOR EACH ROW EXECUTE PROCEDURE videos_tsvector_trigger();