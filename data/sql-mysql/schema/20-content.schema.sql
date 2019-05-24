CREATE TABLE content_namespaces (
  `id` INT,
  `name` VARCHAR(64),
  CONSTRAINT `content_namespaces_key` PRIMARY KEY ( `id` ),
  CONSTRAINT `content_namespaces_unique_name` UNIQUE ( `name` ),
  CONSTRAINT `content_namespaces_refs_entities` FOREIGN KEY ( `id` ) REFERENCES `entities` ( `id` )
);

CREATE TABLE content_sources (
  `id` INT,
  `namespace` INT,
  `source_type` ENUM('NONE', 'URL', 'GITLAB') NOT NULL,
  CONSTRAINT `content_sources_key` PRIMARY KEY ( `id` ),
  CONSTRAINT `content_sources_refs_entities` FOREIGN KEY ( `id` ) REFERENCES `entities` ( `id` ),
  CONSTRAINT `content_sources_refs_namespaces` FOREIGN KEY ( `namespace` ) REFERENCES `content_namespaces` ( `id` )
);

-- TODO: this would be much better stored as "just JSON". We don't need to
-- search it. The load on a "key-value" DB is so fast, we could move this to
-- something else and take the load off the primary DB.
CREATE TABLE content_sources_config (
  `source` INT,
  `key` VARCHAR(255),
  `value` VARCHAR(255),
  CONSTRAINT `content_sources_config_key` PRIMARY KEY ( `source`, `key` ),
  CONSTRAINT `content_sources_config_refs_content_sources` FOREIGN KEY ( `source` ) REFERENCES `content_sources` ( `id` )
);

CREATE TABLE content_summary (
  `id`          INT,
  `namespace`   INT NOT NULL,
  `source`      INT NOT NULL,
  `extern_path` VARCHAR(255),
  `type`        ENUM('TEXT') NOT NULL,
  `slug`        VARCHAR(40),
  `title`       VARCHAR(255) NOT NULL,
  `summary`     TEXT,
  `version_cookie` VARCHAR(255),
  CONSTRAINT `content_summary_key` PRIMARY KEY ( `id` ),
  CONSTRAINT `content_summary_refs_entities` FOREIGN KEY ( `id` ) REFERENCES `entities` ( `id` ),
  CONSTRAINT `content_summary_refs_namespaces` FOREIGN KEY ( `namespace` ) REFERENCES `content_namespaces` ( `id` ),
  CONSTRAINT `content_summary_refs_sources` FOREIGN KEY ( `source` ) REFERENCES `content_sources` ( `id` )
);

DELIMITER //
CREATE TRIGGER `content_summary_slug_constraint`
  BEFORE INSERT ON content_summary FOR EACH ROW
    BEGIN
      IF new.slug REGEXP '[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}' THEN
        SIGNAL SQLSTATE '45000'
          SET MESSAGE_TEXT = 'Cannot use slug which may be confused with a UUID.';
      END IF;
    END;//
CREATE TRIGGER `content_summary_type_constraint`
  BEFORE UPDATE ON content_summary FOR EACH ROW
    BEGIN
      IF new.type != old.type THEN
        SIGNAL SQLSTATE '45000'
          SET MESSAGE_TEXT = 'The content type cannot be changed.';
      END IF;
    END;//
DELIMITER ;

CREATE TABLE content_type_text (
  `id` INT,
  `format` ENUM('TEXT', 'MD', 'HTML', 'CODE'),
  `text` TEXT, -- TODO: in future, possibly expand to 'MEDIUMTEXT', but want to enusre we can enforce size limits to avoid DOS attacks.
               -- The other option would be to keep text and require large works (like a book) to be broken up into multiple content records.
  `last_sync` INT UNSIGNED,
  CONSTRAINT `content_type_text_key` PRIMARY KEY ( `id` ),
  CONSTRAINT `content_type_text_refs_content` FOREIGN KEY ( `id` ) REFERENCES `content_summary` ( `id` )
);

DELIMITER //
CREATE TRIGGER `content_type_text_last_sync_update`
  BEFORE UPDATE ON content_type_text FOR EACH ROW
    BEGIN
      IF new.last_sync = 0 THEN
        SET new.last_sync=UNIX_TIMESTAMP();
      END IF;
    END;//
DELIMITER ;
