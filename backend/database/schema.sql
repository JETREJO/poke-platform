CREATE TABLE pokemon_type (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE generation (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50) NOT NULL,
  region VARCHAR(50) NOT NULL
);

CREATE TABLE pokemon_type_effectiveness (
  attacking_type_id INTEGER NOT NULL,
  defending_type_id INTEGER NOT NULL,
  multiplier NUMERIC(3,1) NOT NULL,
  PRIMARY KEY (
      attacking_type_id,
      defending_type_id
  ),
  FOREIGN KEY (attacking_type_id) REFERENCES pokemon_type(id),
  FOREIGN KEY (defending_type_id) REFERENCES pokemon_type(id)
);

CREATE TABLE pokemon (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  primary_type_id INTEGER NOT NULL,
  secondary_type_id INTEGER,
  generation_id INTEGER NOT NULL,
  sprite VARCHAR(255),
  shiny BOOLEAN DEFAULT FALSE,
  FOREIGN KEY (primary_type_id) REFERENCES pokemon_type(id),
  FOREIGN KEY (secondary_type_id) REFERENCES pokemon_type(id),
  FOREIGN KEY (generation_id) REFERENCES generation(id)
);

-- CREATE TABLE trainer_pokemon (
--   id SERIAL PRIMARY KEY,
--   trainer_id INTEGER NOT NULL,
--   pokemon_id INTEGER NOT NULL,
--   nickname VARCHAR(100),
--   level INTEGER NOT NULL,
--   battles_won INTEGER DEFAULT 0,
--   shiny BOOLEAN DEFAULT FALSE,
--   FOREIGN KEY (pokemon_id) REFERENCES pokemon(id)
-- );

-- CREATE TABLE attack (
--   id SERIAL PRIMARY KEY,
--   name VARCHAR(100) NOT NULL,
--   power INTEGER,
--   accuracy INTEGER,
--   pp INTEGER NOT NULL,
--   description TEXT,
--   type_id INTEGER NOT NULL,
--   FOREIGN KEY (type_id) REFERENCES pokemon_type(id)
-- );

-- CREATE TABLE pokemon_attack (
--   pokemon_id INTEGER,
--   attack_id INTEGER,
--   PRIMARY KEY (
--       pokemon_id,
--       attack_id
--   )
-- );

-- CREATE TABLE evolution (
--   id SERIAL PRIMARY KEY,
--   pokemon_id INTEGER NOT NULL,
--   evolves_to_id INTEGER NOT NULL,
--   level_required INTEGER
-- );

-- CREATE TABLE pokemon_type_effectiveness (
--     attacking_type_id INTEGER NOT NULL,
--     defending_type_id INTEGER NOT NULL,
--     multiplier NUMERIC(3,1) NOT NULL,
--     PRIMARY KEY (
--         attacking_type_id,
--         defending_type_id
--     ),
--     FOREIGN KEY (attacking_type_id) REFERENCES pokemon_type(id),
--     FOREIGN KEY (defending_type_id) REFERENCES pokemon_type(id)
-- );