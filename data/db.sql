-- Create Employee Table
CREATE TABLE Employee (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    profileImage TEXT,
    salary REAL
);