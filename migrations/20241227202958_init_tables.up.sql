CREATE TABLE users(
	user_id SERIAL PRIMARY KEY,
	user_name VARCHAR(50) NOT NULL UNIQUE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE decks(
	deck_id SERIAL PRIMARY KEY,
	user_id INT NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
	deck_name VARCHAR(100) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	UNIQUE (user_id, deck_name)
);

CREATE TABLE cards(
	card_id SERIAL PRIMARY KEY,
	deck_id INT NOT NULL REFERENCES decks(deck_id) ON DELETE CASCADE,
	front TEXT NOT NULL,
	back TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
