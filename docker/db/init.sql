CREATE TABLE payments (
  id SERIAL PRIMARY KEY,
  idempotency_key SERIAL UNIQUE,
  amount VARCHAR(64) NOT NULL,
  currency VARCHAR(3) NOT NULL,
  payment_status VARCHAR(20) NOT NULL,
  payment_method VARCHAR(1024) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
