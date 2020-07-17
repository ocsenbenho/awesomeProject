-- +migrate Up
CREATE table "user"(
"user_id" text PRIMARY KEY ,
"fullname" text,
"email" text UNIQUE ,
"password" text,
"role" text,
"create_at" TIMESTAMP NOT NULL ,
"update_at" TIMESTAMP NOT NULL
);
CREATE TABLE "repo"(
"name" text PRIMARY KEY ,
"decripstion" text,
"url" text,
"color" text,
"lang" text,
"fork" text,
"starts" text,
"start_today" text,
"build_by" text,
"create_at" TIMESTAMP NOT NULL ,
"update_at" TIMESTAMP NOT NULL
);

CREATE TABLE "bookmark"(
"bid" text PRIMARY KEY,
"user_id" text,
"repo_name" text,
"create_at" TIMESTAMP NOT NULL,
"update_at" TIMESTAMP NOT NULL
);

ALTER TABLE "bookmark" ADD FOREIGN KEY ("user_id") REFERENCES "user"("user_id");
ALTER TABLE "bookmark" ADD FOREIGN KEY ("repo_name") REFERENCES "repo"("name");

-- +migrate Down
DROP TABLE "bookmark";
DROP TABLE "user";
DROP TABLE "repo";

