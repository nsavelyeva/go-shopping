BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "items" (
	"id"	integer,
	"name"	varchar(255),
	"price"	real,
	"sold"	bool,
	PRIMARY KEY("id" AUTOINCREMENT)
);
INSERT INTO "items" VALUES (1,'Aladdin''s lamp',999.0,1);
INSERT INTO "items" VALUES (2,'Rejuvenating apples',99.0,0);
INSERT INTO "items" VALUES (3,'Magic table-cloth',399.0,0);
INSERT INTO "items" VALUES (4,'Seven-league boots',59.0,0);
INSERT INTO "items" VALUES (5,'Magic wand',799.0,0);
INSERT INTO "items" VALUES (6,'Hat of invisibility',199.0,0);
INSERT INTO "items" VALUES (7,'Flying carpet',399.0,0);
INSERT INTO "items" VALUES (8,'Elixir of life',599.0,0);
COMMIT;
