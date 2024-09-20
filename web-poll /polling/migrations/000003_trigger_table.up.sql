-- Polllar uchun keyingi bo'sh raqamni olish funksiyasi
CREATE OR REPLACE FUNCTION get_next_poll_num() RETURNS INT AS $$
DECLARE
    next_num INT;
BEGIN
    -- Polllar jadvalidagi eng katta poll_num qiymatini olish va 1 qo'shish
    SELECT COALESCE(MAX(poll_num), 0) + 1 INTO next_num FROM polls;
    RETURN next_num;
END;
$$ LANGUAGE plpgsql;

-- Polllar o'chirilgandan so'ng raqamlarni qayta tartiblash funksiyasi
CREATE OR REPLACE FUNCTION renumber_polls() RETURNS TRIGGER AS $$
BEGIN
    -- Polllar jadvalini yangi tartib bilan yangilash
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

-- Polllar o'chirilgandan so'ng trigger yaratish
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

-- Poll qo'shish funksiyasini yaratish (keyingi bo'sh poll_num bilan)
CREATE OR REPLACE FUNCTION insert_poll(p_title VARCHAR, p_subtitle VARCHAR, p_options JSONB, p_feedbacks JSONB) RETURNS VOID AS $$
DECLARE
    next_num INT;
BEGIN
    -- Keyingi bo'sh poll raqamini olish
    next_num := get_next_poll_num();
    -- Pollni qo'shish
    INSERT INTO polls (poll_num, title, subtitle, options, feedbacks) VALUES (next_num, p_title, p_subtitle, p_options, p_feedbacks);
END;
$$ LANGUAGE plpgsql;

-- Savollar uchun keyingi bo'sh raqamni olish funksiyasi (poll_id bo'yicha)
CREATE OR REPLACE FUNCTION get_next_question_num(p_poll_id UUID) RETURNS INT AS $$
DECLARE
    next_num INT;
BEGIN
    -- Berilgan poll_id uchun eng katta num qiymatini olish va 1 qo'shish
    SELECT COALESCE(MAX(num), 0) + 1 INTO next_num FROM questions WHERE poll_id = p_poll_id;
    RETURN next_num;
END;
$$ LANGUAGE plpgsql;

-- Savollar o'chirilgandan so'ng raqamlarni qayta tartiblash funksiyasi
CREATE OR REPLACE FUNCTION renumber_questions() RETURNS TRIGGER AS $$
BEGIN
    -- Savollar jadvalini yangi tartib bilan yangilash
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

-- Savollar o'chirilgandan so'ng trigger yaratish
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

-- Savol qo'shish funksiyasini yaratish (keyingi bo'sh num bilan)
CREATE OR REPLACE FUNCTION insert_question(p_content TEXT, p_poll_id UUID) RETURNS VOID AS $$
DECLARE
    next_num INT;
BEGIN
    -- Keyingi bo'sh savol raqamini olish
    next_num := get_next_question_num(p_poll_id);
    -- Savolni qo'shish
    INSERT INTO questions (num, content, poll_id) VALUES (next_num, p_content, p_poll_id);
END;
$$ LANGUAGE plpgsql;
