CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    price NUMERIC,
    brand VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);