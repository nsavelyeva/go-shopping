BEGIN TRANSACTION;

CREATE TABLE IF NOT EXISTS "items" (
	"id"	integer,
	"name"	varchar(255),
	"price"	real,
	"sold"	bool,
    "created_at"	datetime,
    "updated_at"	datetime,
    "deleted_at"	datetime,
	PRIMARY KEY("id" AUTOINCREMENT)
);

INSERT INTO "items" VALUES (1,'Aladdin''s lamp',999.0,1,'2021-12-28 15:20:20.045968+01:00','2021-12-28 15:20:20.045968+01:00',NULL);
INSERT INTO "items" VALUES (2,'Rejuvenating apples',99.0,0,'2021-12-28 15:22:52.702091+01:00','2021-12-28 15:22:52.702091+01:00',NULL);
INSERT INTO "items" VALUES (3,'Magic table-cloth',399.0,0,'2021-12-28 15:23:13.876102+01:00','2021-12-28 15:23:13.876102+01:00',NULL);
INSERT INTO "items" VALUES (4,'Seven-league boots',59.0,0,'2021-12-28 15:23:29.858956+01:00','2021-12-28 15:23:29.858956+01:00',NULL);
INSERT INTO "items" VALUES (5,'Magic wand',799.0,0,'2021-12-28 15:23:46.370883+01:00','2021-12-28 15:23:46.370883+01:00',NULL);
INSERT INTO "items" VALUES (6,'Hat of invisibility',199.0,0,'2021-12-28 15:24:00.473628+01:00','2021-12-28 15:24:00.473628+01:00',NULL);
INSERT INTO "items" VALUES (7,'Flying carpet',399.0,0,'2021-12-28 15:24:14.116392+01:00','2021-12-28 15:24:14.116392+01:00',NULL);
INSERT INTO "items" VALUES (8,'Elixir of life',599.0,0,'2021-12-28 15:24:28.294611+01:00','2021-12-28 15:24:28.294611+01:00',NULL);

CREATE INDEX IF NOT EXISTS "idx_items_deleted_at" ON "items" (
	"deleted_at"
);

COMMIT;
