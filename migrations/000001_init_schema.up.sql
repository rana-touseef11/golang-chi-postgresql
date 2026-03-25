CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TYPE user_status AS ENUM('Active', 'Inactive', 'Blocked', 'Deleted');
CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),

  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL UNIQUE,
  phone varchar(255) NOT NULL UNIQUE,

  address varchar(255),
  status user_status NOT NULL DEFAULT 'Active',

  created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- function to update updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$ language 'plpgsql';

-- trigger
CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();