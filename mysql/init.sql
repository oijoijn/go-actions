CREATE TABLE IF NOT EXISTS todos (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO todos (title, content, created_at, updated_at) 
VALUES 
    ('todo1 title', 'todo1 content', NOW(), NOW()),
    ('todo2 title', 'todo2 content', NOW(), NOW()),
    ('todo3 title', 'todo3 content', NOW(), NOW());

