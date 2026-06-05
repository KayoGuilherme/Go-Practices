CREATE TABLE IF NOT EXISTS "products" (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(255) NOT NULL,
    price           NUMERIC(10,2) NOT NULL,
    description     VARCHAR(255) NOT NULL,
    stock           INT NOT NULL DEFAULT 0,
    is_on_sale      BOOLEAN NOT NULL DEFAULT FALSE,
    weight          INT,
    height          INT,
    width           INT,
    diameter        INT,
    length          INT,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);