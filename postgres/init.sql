CREATE TABLE IF NOT EXISTS logs (
    id SERIAL PRIMARY KEY,
    request_body TEXT,
    response_body TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
