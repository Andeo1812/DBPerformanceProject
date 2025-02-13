CREATE EXTENSION IF NOT EXISTS citext;

-- Main
CREATE UNLOGGED TABLE IF NOT EXISTS users
(
    user_id  bigserial,
    nickname citext COLLATE "ucs_basic" NOT NULL UNIQUE PRIMARY KEY,
    fullname text                       NOT NULL,
    about    text,
    email    citext                     NOT NULL UNIQUE
);

CREATE UNLOGGED TABLE IF NOT EXISTS forums
(
    forum_id       bigserial,
    users_nickname citext NOT NULL REFERENCES users (nickname),
    slug           citext NOT NULL PRIMARY KEY,
    title          text   NOT NULL,
    posts          int DEFAULT 0,
    threads        int DEFAULT 0
);

CREATE UNLOGGED TABLE IF NOT EXISTS threads
(
    thread_id bigserial PRIMARY KEY NOT NULL,
    author    citext                NOT NULL REFERENCES users (nickname),
    forum     citext                NOT NULL REFERENCES forums (slug),
    title     text                  NOT NULL,
    message   text                  NOT NULL,
    votes     integer                  DEFAULT 0,
    slug      citext,
    created   timestamp with time zone DEFAULT now()
);

CREATE UNLOGGED TABLE IF NOT EXISTS posts
(
    post_id   bigserial PRIMARY KEY NOT NULL UNIQUE,
    forum     citext REFERENCES forums (slug),
    thread_id integer REFERENCES threads (thread_id),
    author    citext                NOT NULL REFERENCES users (nickname),
    parent    int                      DEFAULT 0,
    message   text                  NOT NULL,
    is_edited bool                     DEFAULT FALSE,
    created   timestamp with time zone DEFAULT now(),
    path      bigint[]                 DEFAULT ARRAY []::INTEGER[]
);

-- M:N
CREATE UNLOGGED TABLE IF NOT EXISTS user_votes
(
    nickname  citext NOT NULL REFERENCES users (nickname),
    thread_id int    NOT NULL REFERENCES threads (thread_id),
    voice     int    NOT NULL
);

CREATE UNLOGGED TABLE IF NOT EXISTS user_forums
(

    nickname citext COLLATE "ucs_basic" NOT NULL REFERENCES users (nickname),
    forum    citext                     NOT NULL REFERENCES forums (slug),
    fullname text,
    about    text,
    email    citext,
    CONSTRAINT user_forum_key unique (nickname, forum)
);

-- Storage features
CREATE OR REPLACE FUNCTION function_path_update() RETURNS TRIGGER AS
$$
BEGIN
    new.path = (SELECT path FROM posts WHERE post_id = new.parent) || new.post_id;
    RETURN new;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER path_upd
    BEFORE INSERT
    ON posts
    FOR EACH ROW
EXECUTE PROCEDURE function_path_update();

-- Update denormal fields
-- For votes
CREATE OR REPLACE FUNCTION function_insert_votes_into_threads()
    RETURNS TRIGGER AS
$$
BEGIN
    UPDATE threads
    SET votes = votes + NEW.voice
    WHERE thread_id = NEW.thread_id;
    RETURN NEW;
END;
$$ language plpgsql;

CREATE TRIGGER insert_votes
    AFTER INSERT
    ON user_votes
    FOR EACH ROW
EXECUTE PROCEDURE function_insert_votes_into_threads();

CREATE OR REPLACE FUNCTION function_update_votes_in_threads()
    RETURNS TRIGGER AS
$$
BEGIN
    UPDATE threads
    SET votes = votes + NEW.voice - OLD.voice
    WHERE thread_id = NEW.thread_id;
    RETURN NEW;
END;
$$ language plpgsql;

CREATE TRIGGER update_votes
    AFTER UPDATE
    ON user_votes
    FOR EACH ROW
EXECUTE PROCEDURE function_update_votes_in_threads();

-- For counters
CREATE OR REPLACE FUNCTION function_count_posts()
    RETURNS TRIGGER AS
$$
BEGIN
    UPDATE forums
    SET posts = forums.posts + 1
    WHERE slug = NEW.forum;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_count_posts
    AFTER INSERT
    ON posts
    FOR EACH ROW
EXECUTE PROCEDURE function_count_posts();

CREATE OR REPLACE FUNCTION function_count_threads()
    RETURNS TRIGGER AS
$$
BEGIN
    UPDATE forums
    SET threads = forums.threads + 1
    WHERE slug = NEW.forum;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_count_threads
    AFTER INSERT
    ON threads
    FOR EACH ROW
EXECUTE PROCEDURE function_count_threads();

CREATE OR REPLACE FUNCTION function_update_user_forum()
    RETURNS TRIGGER AS

-- info posts, threads
$$
DECLARE
    _nickname citext;
    _fullname text;
    _about    text;
    _email    citext;
BEGIN
    SELECT u.nickname, u.fullname, u.about, u.email
    FROM users u
    WHERE u.nickname = NEW.author
    INTO _nickname, _fullname, _about, _email;

    INSERT INTO user_forums (nickname, fullname, about, email, forum)
    VALUES (_nickname, _fullname, _about, _email, NEW.forum)
    ON CONFLICT DO NOTHING;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER update_user_forum
    AFTER INSERT
    ON threads
    FOR EACH ROW
EXECUTE PROCEDURE function_update_user_forum();

CREATE TRIGGER update_users_forum
    AFTER INSERT
    ON posts
    FOR EACH ROW
EXECUTE PROCEDURE function_update_user_forum();

-- Optimization
-- forums
-- Хеши для быстрого доступа и полного сравнение на равенство
CREATE INDEX IF NOT EXISTS forum_slug_hash ON forums USING hash (slug);
CREATE INDEX IF NOT EXISTS forum_user_hash ON forums USING hash (users_nickname);
-- user_forums
CREATE INDEX IF NOT EXISTS users_to_forums_forum_compare ON user_forums (forum);
CREATE INDEX IF NOT EXISTS users_to_forums_nickname_compare ON user_forums (nickname);
CREATE INDEX IF NOT EXISTS users_to_forums_forum_nickname ON user_forums (forum, nickname);
CREATE INDEX IF NOT EXISTS users_to_forum_nickname_forum ON user_forums (nickname, fullname, about, email);
-- users
CREATE INDEX IF NOT EXISTS user_nickname_compare ON users (nickname);
CREATE INDEX IF NOT EXISTS user_all ON users (nickname, fullname, about, email);
-- posts
CREATE INDEX IF NOT EXISTS post_thread ON posts USING hash (thread_id);
CREATE INDEX IF NOT EXISTS post_parent ON posts (thread_id, post_id, (path[1]), parent);
CREATE INDEX IF NOT EXISTS post_path_1_path ON posts ((path[1]), path);
CREATE INDEX IF NOT EXISTS post_thread_path ON posts (thread_id, path);
-- votes
CREATE UNIQUE INDEX IF NOT EXISTS votes_all ON user_votes (nickname, thread_id, voice);
CREATE UNIQUE INDEX IF NOT EXISTS votes ON user_votes (nickname, thread_id);
-- threads
CREATE INDEX IF NOT EXISTS th_slug_hash ON threads USING hash (slug);
CREATE INDEX IF NOT EXISTS th_user_hash ON threads USING hash (author);
CREATE INDEX IF NOT EXISTS th_created ON threads (created);
CREATE INDEX IF NOT EXISTS th_forum ON threads USING hash (forum);
CREATE INDEX IF NOT EXISTS th_forum_created ON threads (forum, created);

VACUUM ANALYSE;
