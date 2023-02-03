DO $$
DECLARE
  i INTEGER;
  opinion VARCHAR(20);
  size VARCHAR(20);
  shape VARCHAR(20);
  origin VARCHAR(20);
  item VARCHAR(20);
  rand1 NUMERIC;
  rand2 NUMERIC;
  rand3 NUMERIC;
  rand4 NUMERIC;
  rand5 NUMERIC;
  name VARCHAR(100);
  subcategory VARCHAR(100);
  brand VARCHAR(100);
  retail_price NUMERIC;
  status BOOLEAN;
BEGIN
  FOR i IN 1..200000 LOOP
    rand1 := RANDBETWEEN(1, 5, random()::NUMERIC);
    rand2 := RANDBETWEEN(1, 5, random()::NUMERIC);
    rand3 := RANDBETWEEN(1, 5, random()::NUMERIC);
    rand4 := RANDBETWEEN(1, 5, random()::NUMERIC);
    rand5 := RANDBETWEEN(1, 5, random()::NUMERIC);
    SELECT (
      CASE
        WHEN rand1 = 1 THEN 'Unusual'
        WHEN rand1 = 2 THEN 'Lovely'
        WHEN rand1 = 3 THEN 'Beautiful'
        WHEN rand1 = 4 THEN 'Exquisite'
        ELSE 'Ugly'
      END ) INTO opinion;
    
    SELECT (
      CASE
        WHEN rand2 = 1 THEN 'Big'
        WHEN rand2 = 2 THEN 'Huge'
        WHEN rand2 = 3 THEN 'Small'
        WHEN rand2 = 4 THEN 'Tiny'
        ELSE 'Average'
      END) INTO size;

    SELECT (
      CASE
        WHEN rand3 = 1 THEN 'Round'
        WHEN rand3 = 2 THEN 'Circular'
        WHEN rand3 = 3 THEN 'Rectangular'
        WHEN rand3 = 4 THEN 'Oval'
        ELSE 'Cylindrical'
      END) INTO shape;

    SELECT (
      CASE
        WHEN rand4 = 1 THEN 'Australian'
        WHEN rand4 = 2 THEN 'Belgian'
        WHEN rand4 = 3 THEN 'Indonesian'
        WHEN rand4 = 4 THEN 'Norwegian'
        ELSE 'Singaporean'
      END) INTO origin;

    SELECT (
      CASE
        WHEN rand5 = 1 THEN 'Table'
        WHEN rand5 = 2 THEN 'Monitor'
        WHEN rand5 = 3 THEN 'Earphones'
        WHEN rand5 = 4 THEN 'Bookshelf'
        ELSE 'Fan'
      END) INTO item;

    name := opinion || ' ' || size || ' ' || shape || ' ' || origin || ' ' || item;

    SELECT (
      CASE
        WHEN rand5 = 1 THEN 'Furniture'
        WHEN rand5 = 2 THEN 'Electronics'
        WHEN rand5 = 3 THEN 'Electronics'
        WHEN rand5 = 4 THEN 'Furniture'
        ELSE 'Electronics'
      END) INTO subcategory;
    
    SELECT (
      CASE
        WHEN rand3 = 1 THEN 'Panasonic'
        WHEN rand3 = 2 THEN 'LG'
        WHEN rand3 = 3 THEN 'Toshiba'
        WHEN rand3 = 4 THEN 'Xiaomi'
        ELSE 'Samsung'
      END) INTO brand;
    retail_price := RANDBETWEEN(100000, 50000000, random()::NUMERIC);
    status := (i % 2) = 0;

    INSERT INTO products (name, subcategory, brand, retail_price, status)
    VALUES (name, subcategory, brand, retail_price, status);
  END LOOP;
END $$;