CREATE TABLE payments (
  id SERIAL PRIMARY KEY,
  amount NUMERIC(10,2) NOT NULL,
  currency VARCHAR(3) NOT NULL,
  status VARCHAR(20) NOT NULL,
  payment_method VARCHAR(20) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE credit_cards (
  id SERIAL PRIMARY KEY,
  payment_id INTEGER NOT NULL,
  card_number VARCHAR(20) NOT NULL,
  expiration_date VARCHAR(5) NOT NULL,
  cardholder_name VARCHAR(50) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  FOREIGN KEY (payment_id) REFERENCES payments (id)
);