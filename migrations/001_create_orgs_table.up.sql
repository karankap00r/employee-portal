
    CREATE TABLE IF NOT EXISTS orgs (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE,
        client_id TEXT NOT NULL UNIQUE,
        contact_email TEXT NOT NULL UNIQUE,
        contact_phone TEXT NOT NULL UNIQUE,
        created_at TEXT NOT NULL,
        updated_at TEXT NOT NULL
    );
    
    INSERT INTO orgs (name, client_id, contact_email, contact_phone, created_at, updated_at)
    VALUES
    ('Cercli', 'cercli-client-id', 'contact@cercli.com', '+1234567890', DATETIME('now'), DATETIME('now')),
    ('Google', 'google-client-id', 'contact@google.com', '+1987654321', DATETIME('now'), DATETIME('now'));
    