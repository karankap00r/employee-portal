CREATE TABLE IF NOT EXISTS employees (
                                         id INTEGER PRIMARY KEY AUTOINCREMENT,
                                         employee_id TEXT NOT NULL UNIQUE,
                                         name TEXT NOT NULL,
                                         position TEXT NOT NULL,
                                         email TEXT NOT NULL UNIQUE,
                                         salary INTEGER NOT NULL,
                                         created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER IF NOT EXISTS update_employee_updated_at
AFTER UPDATE ON employees
                            FOR EACH ROW
BEGIN
UPDATE employees SET updated_at = CURRENT_TIMESTAMP WHERE id = OLD.id;
END;
