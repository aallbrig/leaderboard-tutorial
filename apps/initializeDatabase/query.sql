CREATE TABLE games (
   game_id INT AUTO_INCREMENT PRIMARY KEY,
   username VARCHAR(50) NOT NULL,
   gamedate TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   score INT NOT NULL,
   level INT NOT NULL
);