-- V1__Create_employees_table.sql
CREATE TABLE IF NOT EXISTS employees (
                                         id TEXT PRIMARY KEY,
                                         name TEXT,
                                         age INTEGER,
                                         dept TEXT
);