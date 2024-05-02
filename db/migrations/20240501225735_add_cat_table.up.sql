CREATE TABLE IF NOT EXISTS cats (
    id SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    race VARCHAR(30) NOT NULL CHECK (race IN ('Persian', 'Maine Coon', 'Siamese', 'Ragdoll', 'Bengal', 'Sphynx', 'British Shorthair', 'Abyssinian', 'Scottish Fold', 'Birman')),
    sex VARCHAR(6) NOT NULL CHECK (sex IN ('male', 'female')),
    age_in_month INT NOT NULL CHECK (age_in_month BETWEEN 1 AND 120082),
    description VARCHAR(200) NOT NULL,
    img_urls TEXT[] NOT NULL CHECK (array_length(img_urls, 1) >= 1),
    user_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE NO ACTION ON UPDATE NO ACTION
);
