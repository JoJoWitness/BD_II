CREATE TABLE IF NOT EXISTS phrases (
  id          SERIAL PRIMARY KEY,
  text        TEXT NOT NULL
);

INSERT INTO phrases (text) VALUES
  ('Hola Mundo'),
  ('Comunicaciones II'),
  ('Practica II'),
  ('HTTP'),
  ('Go'),
  ('Javascript');
