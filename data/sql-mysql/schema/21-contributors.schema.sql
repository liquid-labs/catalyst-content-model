CREATE TABLE contributors (
  `id` INT,
  `content` INT,
  `role` VARCHAR(64),
  `summary_credit_order` TINYINT UNSIGNED,
  CONSTRAINT `contributors_key` PRIMARY KEY ( `id`, `content` ),
  CONSTRAINT `contributors_refs_persons` FOREIGN KEY ( `id` ) REFERENCES `persons` ( `id` ),
  CONSTRAINT `contributors_refs_content` FOREIGN KEY ( `content` ) REFERENCES `content_summary` ( `id` )
);

DELIMITER //
CREATE TRIGGER `contributors_credit_order_constraint`
-- We expect contributers to be cleared and rebuilt
  BEFORE INSERT ON contributors FOR EACH ROW
    BEGIN
      IF new.summary_credit_order <= 0 OR new.summary_credit_order > 3 THEN
        SIGNAL SQLSTATE '45000'
          SET MESSAGE_TEXT = 'Summary credit order must be 1, 2, or 3.';
      END IF;
    END;//

CREATE TRIGGER `contributors_update_disallowed`
-- We expect contributers to be cleared and rebuilt
  BEFORE UPDATE ON contributors FOR EACH ROW
    SIGNAL SQLSTATE '45000'
      SET MESSAGE_TEXT = 'Invalid update; current model expects contributors to be cleared and refreshed.';//

DELIMITER ;
