
    CREATE TABLE IF NOT EXISTS leave_balances (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        org_id INTEGER NOT NULL,
        employee_id TEXT NOT NULL,
        leave_type TEXT NOT NULL,
        annual_balance INTEGER NOT NULL,
        created_at TEXT NOT NULL,
        updated_at TEXT NOT NULL,
        FOREIGN KEY(org_id) REFERENCES orgs(id),
        UNIQUE(org_id, employee_id, leave_type)
    );
    