
    CREATE TABLE IF NOT EXISTS employees (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        org_id INTEGER NOT NULL,
        employee_id TEXT NOT NULL,
        name TEXT NOT NULL,
        position TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE,
        salary INTEGER NOT NULL,
        created_at TEXT NOT NULL,
        updated_at TEXT NOT NULL,
        FOREIGN KEY(org_id) REFERENCES orgs(id),
        UNIQUE(org_id, employee_id),
        UNIQUE(email)
    );
    