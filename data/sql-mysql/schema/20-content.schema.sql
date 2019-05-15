CREATE TABLE content_summary (
  `id` INT,
  `title` VARCHAR(255) NOT NULL,
  `summary` TEXT,
  `extern_path` VARCHAR(255),
  `namespace` VARCHAR(24),
  `slug` VARCHAR(40),
  `type` ENUM('TEXT') NOT NULL,
  `version_cookie` VARCHAR(255),
  CONSTRAINT `content_summary_key` PRIMARY KEY ( `id` ),
  CONSTRAINT `content_summary_ref_entities` FOREIGN KEY ( `id` ) REFERENCES `entities` ( `id` )
);

DELIMITER //
CREATE TRIGGER `content_summary_slug_constraint`
  BEFORE INSERT ON content FOR EACH ROW
    BEGIN
      IF new.slug REGEXP '[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}'
        SIGNAL SQLSTATE '45000'
          SET MESSAGE_TEXT = 'Cannot use slug which may be confused with a UUID.';
      END IF;
    END;//
CREATE TRIGGER `content_summary_type_constraint`
  BEFORE UPDATE ON content FOR EACH ROW
    BEGIN
      IF new.type != old.type
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
  CONSTRAINT `content_type_text_ref_content` FOREIGN KEY ( `id` ) REFERENCES `content` ( `id` )
);

CREATE TRIGGER `content_type_text_last_sync_update`
  BEFORE UPDATE ON content FOR EACH ROW
    BEGIN
      IF new.last_sync == 0
        SET new.last_sync=UNIX_TIMESTAMP();
      END IF;
    END;//
DELIMITER ;
