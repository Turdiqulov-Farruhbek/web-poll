-- Create the enum type for gender
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender_type') THEN
        CREATE TYPE gender_type AS ENUM ('male', 'female');
    END IF;
END
$$;


-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    gender gender_type NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    working_experience INT CHECK (working_experience >= 0),
    level_type VARCHAR NOT NULL,
    is_confirmed BOOLEAN DEFAULT FALSE,
    role VARCHAR(20) DEFAULT 'user',
    confirmed_at TIMESTAMP
);

-- Create polls table
CREATE TABLE IF NOT EXISTS polls (
    id UUID PRIMARY KEY,
    poll_num INT UNIQUE NOT NULL,
    title VARCHAR(255) NOT NULL,
    options JSONB NOT NULL
);

-- Create questions table
CREATE TABLE IF NOT EXISTS questions (
    id UUID PRIMARY KEY,
    num INT NOT NULL,
    content TEXT NOT NULL,
    poll_id UUID REFERENCES polls(id) ON DELETE CASCADE,
    UNIQUE (poll_id, num)
);

-- ############# Storing Results
CREATE TABLE IF NOT EXISTS results (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    poll_id UUID REFERENCES polls(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS poll_answers (
    id UUID PRIMARY KEY,
    result_id UUID REFERENCES results(id) ON DELETE CASCADE,
    question_num INT,
    answer INT
);
