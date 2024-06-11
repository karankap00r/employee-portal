-- Create the email_reports table
CREATE TABLE IF NOT EXISTS email_reports
(
    id             INTEGER PRIMARY KEY AUTOINCREMENT,
    org_id         INTEGER  NOT NULL,
    report_type    TEXT     NOT NULL,
    cron_frequency TEXT     NOT NULL,
    status         TEXT     NOT NULL,
    email          TEXT     NOT NULL,
    created_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at     DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (org_id) REFERENCES orgs (id)
);
