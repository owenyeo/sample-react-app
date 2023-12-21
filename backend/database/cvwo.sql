-- Drop the table if it exists
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS posts

-- Create the album table with an auto-incrementing id (using SERIAL)
CREATE TABLE users (
  name       VARCHAR(50) NOT NULL PRIMARY KEY
);
