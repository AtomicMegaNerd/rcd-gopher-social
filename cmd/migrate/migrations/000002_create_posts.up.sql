CREATE TABLE IF NOT EXISTS posts (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  user_id bigint NOT NULL,
  content TEXT NOT NULL,
  tags VARCHAR(100)[],
  created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);
