-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE cars(
    car_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users,
    car_value TEXT NOT NULL,
    car_make TEXT NOT NULL,
    car_model TEXT NOT NULL,
    year_of_manufacture DATETIME NOT NULL,
    car_use TEXT NOT NULL,
    policy_period DATETIME NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
    -- CONSTRAINT fk_car_id FOREIGN KEY (car_id) users (user_id)
);

-- CREATE UNIQUE INDEX user_email
--     ON cars (car_id);