-- Manually create database with name: 'go_crm'
-- Connect to the database by postbird..
-- Create the Customer table, import this script to create database..
CREATE TABLE IF NOT EXISTS Customer (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    Role VARCHAR(50),
    Email VARCHAR(100) NOT NULL,
    Phone VARCHAR(20),
    Contacted BOOLEAN DEFAULT FALSE
);

-- Insert fake users
INSERT INTO Customer (Name, Role, Email, Phone, Contacted) VALUES
('Alice Johnson', 'Manager', 'alice.johnson@example.com', '555-0100', TRUE),
('Bob Smith', 'Developer', 'bob.smith@example.com', '555-0101', FALSE),
('Charlie Brown', 'Designer', 'charlie.brown@example.com', '555-0102', TRUE),
('Diana Prince', 'Analyst', 'diana.prince@example.com', '555-0103', FALSE),
('Edward Norton', 'Tester', 'edward.norton@example.com', '555-0104', TRUE),
('Fiona Apple', 'Developer', 'fiona.apple@example.com', '555-0105', FALSE),
('George Martin', 'Manager', 'george.martin@example.com', '555-0106', TRUE),
('Helen Hunt', 'Designer', 'helen.hunt@example.com', '555-0107', FALSE),
('Ian McKellen', 'Tester', 'ian.mckellen@example.com', '555-0108', TRUE),
('Jane Doe', 'Analyst', 'jane.doe@example.com', '555-0109', FALSE);