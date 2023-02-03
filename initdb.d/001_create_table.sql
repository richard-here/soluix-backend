CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS products (
  id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  subcategory VARCHAR(100) NOT NULL,
  brand VARCHAR(100) NOT NULL,
  retail_price NUMERIC NOT NULL,
  status BOOLEAN NOT NULL
);