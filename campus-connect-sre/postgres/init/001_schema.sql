CREATE TABLE IF NOT EXISTS posts (
    id BIGSERIAL PRIMARY KEY,
    author VARCHAR(100) NOT NULL,
    title VARCHAR(150) NOT NULL,
    content TEXT NOT NULL,
    likes INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO posts (author, title, content, likes)
VALUES
    ('Sanzhar', 'Welcome to Campus Connect', 'This project demonstrates a production-style stack for the midterm.', 3),
    ('Beka', 'Monitoring matters', 'Do not forget to take screenshots of Grafana and Prometheus.', 2),
    ('Syrym', 'Alert test', 'You can manually trigger warning and critical alerts from the debug endpoints.', 1)
ON CONFLICT DO NOTHING;
