CREATE TABLE IF NOT EXISTS duck_entries (
  id VARCHAR(20) NOT NULL,
  fed_time TIMESTAMPTZ NOT NULL,
  food VARCHAR(50) NOT NULL,
  kind_of_food VARCHAR(50) NOT NULL,
  amount_of_food INT NOT NULL,
  location VARCHAR(50) NOT NULL,
  number_of_ducks INT NOT NULL
);

