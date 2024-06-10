-- 0001_create_employees_table.up.sql
CREATE TABLE IF NOT EXISTS employees (
                                         id TEXT PRIMARY KEY,
                                         name TEXT,
                                         age INTEGER,
                                         dept TEXT
);