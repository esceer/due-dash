-- migrate:up
CREATE TABLE templates (
    id INTEGER PRIMARY KEY AUTOINCREMENT, 
    name TEXT NOT NULL,
    frequency TEXT DEFAULT 'ONCE',
    day_of_month INTEGER NOT NULL,
    enabled BOOLEAN DEFAULT 1
);

CREATE TABLE tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT, 
    template_id INTEGER REFERENCES templates(id) ON DELETE SET NULL,
    name TEXT NOT NULL,
    due_date DATE NOT NULL,
    completed_at DATETIME,
    status TEXT DEFAULT 'PENDING',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- migrate:down
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS templates;
