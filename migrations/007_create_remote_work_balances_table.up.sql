CREATE TABLE IF NOT EXISTS remote_work_balances
(
    id             INTEGER PRIMARY KEY AUTOINCREMENT,
    org_id         INTEGER  NOT NULL,
    employee_id    TEXT     NOT NULL,
    type           TEXT     NOT NULL,
    annual_balance INTEGER  NOT NULL,
    created_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (org_id) REFERENCES orgs (id),
    UNIQUE (org_id, employee_id, type)
);
    