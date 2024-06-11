CREATE TABLE IF NOT EXISTS org_holidays
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    org_id     INTEGER  NOT NULL,
    start_date TEXT     NOT NULL,
    end_date   TEXT     NOT NULL,
    name       TEXT     NOT NULL,
    created_by TEXT     NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (org_id) REFERENCES orgs (id),
    UNIQUE (org_id, start_date, end_date)
);

-- Insert sample entries into the org_holidays table
INSERT INTO org_holidays (org_id, start_date, end_date, name, created_by, created_at, updated_at)
VALUES (1, '2024-01-01', '2024-01-01', 'New Years Day', 'admin', DATETIME('now'), DATETIME('now')),
       (1, '2024-12-25', '2024-12-25', 'Christmas Day', 'admin', DATETIME('now'), DATETIME('now')),
       (1, '2024-07-04', '2024-07-04', 'Independence Day', 'admin', DATETIME('now'), DATETIME('now')),
       (2, '2024-01-01', '2024-01-01', 'New Years Day', 'admin', DATETIME('now'), DATETIME('now')),
       (2, '2024-12-25', '2024-12-25', 'Christmas Day', 'admin', DATETIME('now'), DATETIME('now')),
       (2, '2024-07-04', '2024-07-04', 'Independence Day', 'admin', DATETIME('now'), DATETIME('now'));
