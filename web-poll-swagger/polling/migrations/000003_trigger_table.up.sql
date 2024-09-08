-- -- Function to renumber questions after deletion
-- CREATE OR REPLACE FUNCTION renumber_questions() RETURNS TRIGGER AS $$
-- BEGIN
--     WITH ordered_questions AS (
--         SELECT id, ROW_NUMBER() OVER (PARTITION BY poll_id ORDER BY num) AS new_num
--         FROM questions
--     )
--     UPDATE questions
--     SET num = ordered_questions.new_num
--     FROM ordered_questions
--     WHERE questions.id = ordered_questions.id;
--     RETURN NULL;
-- END;
-- $$ LANGUAGE plpgsql;

-- -- Function to renumber polls after deletion
-- CREATE OR REPLACE FUNCTION renumber_polls() RETURNS TRIGGER AS $$
-- BEGIN
--     WITH ordered_polls AS (
--         SELECT id, ROW_NUMBER() OVER (ORDER BY poll_num) AS new_num
--         FROM polls
--     )
--     UPDATE polls
--     SET poll_num = ordered_polls.new_num
--     FROM ordered_polls
--     WHERE polls.id = ordered_polls.id;
--     RETURN NULL;
-- END;
-- $$ LANGUAGE plpgsql;

-- -- Conditional creation of the after_question_delete trigger
-- DO $$
-- BEGIN
--     IF NOT EXISTS (
--         SELECT 1 
--         FROM pg_trigger 
--         WHERE tgname = 'after_question_delete'
--     ) THEN
--         CREATE TRIGGER after_question_delete
--         AFTER DELETE ON questions
--         FOR EACH STATEMENT
--         EXECUTE FUNCTION renumber_questions();
--     END IF;
-- END
-- $$;

-- -- Conditional creation of the after_poll_delete trigger
-- DO $$
-- BEGIN
--     IF NOT EXISTS (
--         SELECT 1 
--         FROM pg_trigger 
--         WHERE tgname = 'after_poll_delete'
--     ) THEN
--         CREATE TRIGGER after_poll_delete
--         AFTER DELETE ON polls
--         FOR EACH STATEMENT
--         EXECUTE FUNCTION renumber_polls();
--     END IF;
-- END
-- $$;




-- Function to get the next available poll number
CREATE OR REPLACE FUNCTION get_next_poll_num() RETURNS INT AS $$
DECLARE
    next_num INT;
BEGIN
    SELECT COALESCE(MAX(poll_num), 0) + 1 INTO next_num FROM polls;
    RETURN next_num;
END;
$$ LANGUAGE plpgsql;

-- Function to renumber polls after deletion
CREATE OR REPLACE FUNCTION renumber_polls() RETURNS TRIGGER AS $$
BEGIN
    WITH ordered_polls AS (
        SELECT id, ROW_NUMBER() OVER (ORDER BY poll_num) AS new_num
        FROM polls
    )
    UPDATE polls
    SET poll_num = ordered_polls.new_num
    FROM ordered_polls
    WHERE polls.id = ordered_polls.id;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Trigger to renumber polls after deletion
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 
        FROM pg_trigger 
        WHERE tgname = 'after_poll_delete'
    ) THEN
        CREATE TRIGGER after_poll_delete
        AFTER DELETE ON polls
        FOR EACH STATEMENT
        EXECUTE FUNCTION renumber_polls();
    END IF;
END
$$;

-- Modify INSERT to use the next available poll number
CREATE OR REPLACE FUNCTION insert_poll(p_title VARCHAR, p_options JSONB, p_feedbacks JSONB) RETURNS VOID AS $$
DECLARE
    next_num INT;
BEGIN
    next_num := get_next_poll_num();
    INSERT INTO polls (poll_num, title, options, feedbacks) VALUES (next_num, p_title, p_options, p_feedbacks);
END;
$$ LANGUAGE plpgsql;


-- %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%% 


-- Function to get the next available question number for a poll
CREATE OR REPLACE FUNCTION get_next_question_num(p_poll_id UUID) RETURNS INT AS $$
DECLARE
    next_num INT;
BEGIN
    SELECT COALESCE(MAX(num), 0) + 1 INTO next_num FROM questions WHERE poll_id = p_poll_id;
    RETURN next_num;
END;
$$ LANGUAGE plpgsql;

-- Function to renumber questions after deletion
CREATE OR REPLACE FUNCTION renumber_questions() RETURNS TRIGGER AS $$
BEGIN
    WITH ordered_questions AS (
        SELECT id, ROW_NUMBER() OVER (PARTITION BY poll_id ORDER BY num) AS new_num
        FROM questions
    )
    UPDATE questions
    SET num = ordered_questions.new_num
    FROM ordered_questions
    WHERE questions.id = ordered_questions.id;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Trigger to renumber questions after deletion
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 
        FROM pg_trigger 
        WHERE tgname = 'after_question_delete'
    ) THEN
        CREATE TRIGGER after_question_delete
        AFTER DELETE ON questions
        FOR EACH STATEMENT
        EXECUTE FUNCTION renumber_questions();
    END IF;
END
$$;

-- Modify INSERT to use the next available question number
CREATE OR REPLACE FUNCTION insert_question(p_content TEXT, p_poll_id UUID) RETURNS VOID AS $$
DECLARE
    next_num INT;
BEGIN
    next_num := get_next_question_num(p_poll_id);
    INSERT INTO questions (num, content, poll_id) VALUES (next_num, p_content, p_poll_id);
END;
$$ LANGUAGE plpgsql;
