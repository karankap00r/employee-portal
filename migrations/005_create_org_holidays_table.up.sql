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
    