-- Drop the table if it exists
DROP TABLE IF EXISTS users;

-- Create the album table with an auto-incrementing id (using SERIAL)
CREATE TABLE users (
  id         SERIAL PRIMARY KEY,
  name       VARCHAR(50) NOT NULL
);

