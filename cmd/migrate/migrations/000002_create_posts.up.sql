CREATE TABLE IF NOT EXISTS posts (
  id bigserial PRIMARY KEY,
  title varchar(255) NOT NULL,
  user_id bigint NOT NULL,
  content text NOT NULL,
  tags varchar(100)[],
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  version int default 0,
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id)
);
