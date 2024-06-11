CREATE TABLE IF NOT EXISTS leave_requests
(
    id             INTEGER PRIMARY KEY AUTOINCREMENT,
    org_id         INTEGER  NOT NULL,
    employee_id    TEXT     NOT NULL,
    leave_category TEXT     NOT NULL,
    start_date     TEXT     NOT NULL,
    end_date       TEXT     NOT NULL,
    reason         TEXT,
    status         TEXT     NOT NULL,
    updated_by     TEXT,
    created_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (org_id) REFERENCES orgs (id)
);
    