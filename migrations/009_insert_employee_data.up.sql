INSERT INTO employees (employee_id, org_id, name, position, email, salary)
VALUES ('E001', 1, 'John Doe', 'Software Engineer', 'john.doe@example.com', 70000),
       ('E002', 1, 'Jane Doe', 'Product Manager', 'jane.doe@example.com', 80000),
       ('E003', 2, 'Alice Johnson', 'HR Specialist', 'alice.johnson@example.com', 60000),
       ('E004', 2, 'Bob Smith', 'Finance Analyst', 'bob.smith@example.com', 75000),
       ('E005', 2, 'Carol White', 'Sales Manager', 'carol.white@example.com', 68000);

INSERT INTO leave_balances (org_id, employee_id, leave_type, annual_balance, created_at, updated_at)
VALUES (1, 'E001', 'Sick Leave', 10, DATETIME('now'), DATETIME('now')),
       (1, 'E002', 'Vacation Leave', 15, DATETIME('now'), DATETIME('now')),
       (2, 'E003', 'Sick Leave', 12, DATETIME('now'), DATETIME('now')),
       (2, 'E005', 'Vacation Leave', 18, DATETIME('now'), DATETIME('now'));

INSERT INTO remote_work_balances (org_id, employee_id, type, annual_balance, created_at, updated_at)
VALUES (1, 'E001', 'LOCAL', 20, DATETIME('now'), DATETIME('now')),
       (1, 'E001', 'CROSS_BORDER', 45, DATETIME('now'), DATETIME('now')),
       (2, 'E003', 'LOCAL', 20, DATETIME('now'), DATETIME('now')),
       (2, 'E003', 'CROSS_BORDER', 45, DATETIME('now'), DATETIME('now'));