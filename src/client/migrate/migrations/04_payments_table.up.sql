-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE payments(
    payment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users,
    insurance_id UUID NOT NULL REFERENCES insurance_details,
    car_id UUID NOT NULL REFERENCES cars,
    payment_frequency TEXT NOT NULL,
    min_deposit INTEGER NOT NULL,
    refund_amount INTEGER NOT NULL,
    monthly_repayment INTEGER NOT NULL,
    total_interest INTEGER NOT NULL,
    payment_mode TEXT NOT NULL,
    currency TEXT,
    status TEXT NOT NULL,
    payment_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP    
);

-- CREATE UNIQUE INDEX user_email
--     ON payments (payments_id);