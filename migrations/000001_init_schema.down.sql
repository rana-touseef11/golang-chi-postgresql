-- drop trigger first
DROP TRIGGER IF EXISTS update_users_updated_at ON users;

-- drop function
DROP FUNCTION IF EXISTS update_updated_at_column;

-- drop table
DROP TABLE IF EXISTS users;

-- 🔴 IMPORTANT: drop enum type
DROP TYPE IF EXISTS user_status;