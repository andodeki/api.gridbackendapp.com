-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE insurance_details(
    insurance_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users,
    car_id UUID NOT NULL REFERENCES cars,
    insurance_type TEXT NOT NULL,
    basic_cover JSON NOT NULL,
    add_ons JSON NOT NULL,
    start_date DATETIME NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

-- CREATE UNIQUE INDEX user_email
--     ON insurance_details (insurance_id);