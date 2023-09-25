CREATE DATABASE IF NOT EXISTS salad_bowl_db
    DEFAULT CHARACTER SET = 'utf8mb4';

USE salad_bowl_db;

CREATE TABLE games (
    id              INT AUTO_INCREMENT,
    current_team    INT,
    turn_end_time   DATETIME,
    PRIMARY KEY (`id`)
);

DROP TABLE words;

CREATE TABLE words (
    id          INT AUTO_INCREMENT NOT NULL,
    game_id     INT NOT NULL,
    word        VARCHAR(255) NOT NULL,
    team        INT,
    PRIMARY KEY (`id`),
    INDEX game_idx (game_id),
    FOREIGN KEY (game_id)
        REFERENCES games(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
/*
    CONSTRAINT game_fk
*/

INSERT INTO words (id, game_id, word, team) VALUES ROW(4, 2, "word1", 1);


DROP TABLE child;
CREATE TABLE child (
    id INT NOT NULL AUTO_INCREMENT,
    parent_id INT NOT NULL,
    INDEX par_ind (parent_id),
    PRIMARY KEY (`id`),
    CONSTRAINT parent_fk
        FOREIGN KEY (parent_id)
            REFERENCES parent(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE
) ENGINE=INNODB;
