CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    allow_comments BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    parent_id INTEGER REFERENCES comments(id) ON DELETE CASCADE,
    author VARCHAR(255) NOT NULL,
    text VARCHAR(2000) NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);

ALTER TABLE comments
ADD CONSTRAINT fk_post FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE;

ALTER TABLE comments
ADD CONSTRAINT fk_parent_comment FOREIGN KEY (parent_id) REFERENCES comments(id) ON DELETE CASCADE;
CREATE TABLE post_comments (
    post_id INTEGER REFERENCES posts(id) ON DELETE CASCADE,
    comment_id INTEGER REFERENCES comments(id) ON DELETE CASCADE,
    PRIMARY KEY (post_id, comment_id)
);

CREATE TABLE comment_replies (
    parent_id INTEGER REFERENCES comments(id) ON DELETE CASCADE,
    child_id INTEGER REFERENCES comments(id) ON DELETE CASCADE,
    PRIMARY KEY (parent_id, child_id)
);