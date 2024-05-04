CREATE TYPE match_status AS ENUM ('pending', 'approved', 'rejected');

CREATE TABLE cat_matches (
    id SERIAL PRIMARY KEY,
    user_cat_id INT NOT NULL,
    match_cat_id INT NOT NULL,
    status match_status DEFAULT 'pending',
    message VARCHAR(200),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_cat_id) REFERENCES cats(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (match_cat_id) REFERENCES cats(id) ON DELETE CASCADE ON UPDATE CASCADE
);
