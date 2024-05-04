CREATE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_id ON users (id);
CREATE INDEX idx_cats_id ON cats (id);
CREATE INDEX idx_cats_user_id ON cats (user_id);
CREATE INDEX idx_cat_matches_id ON cat_matches (id);
CREATE INDEX idx_cat_matches_user_cat_id ON cat_matches (user_cat_id);
CREATE INDEX idx_cat_matches_match_cat_id ON cat_matches (match_cat_id);
