CREATE DATABASE PandaMerch

-- Create User Table
CREATE TABLE Users (
    id INT IDENTITY(1,1) PRIMARY KEY,  -- Auto-incrementing ID
    firstName NVARCHAR(100) NOT NULL,  -- First Name with a max length of 100 characters
    lastName NVARCHAR(100) NOT NULL,   -- Last Name with a max length of 100 characters
    email NVARCHAR(255) NOT NULL UNIQUE, -- Email must be unique, max length 255 characters
    createdAt DATETIME DEFAULT GETDATE(),  -- Set creation date to current date and time
    modifiedAt DATETIME DEFAULT GETDATE()  -- Set modification date to current date and time
);

-- Create Merchandise Table
CREATE TABLE Merchandise (
    id INT IDENTITY(1,1) PRIMARY KEY,  -- Auto-incrementing ID
    name NVARCHAR(200) NOT NULL,       -- Merchandise name with a max length of 200 characters
    price DECIMAL(10, 2) NOT NULL,     -- Price with up to 10 digits and 2 decimal places
    availability BIT NOT NULL,         -- Boolean value (1 for available, 0 for not available)
    createdAt DATETIME DEFAULT GETDATE(),  -- Set creation date to current date and time
    modifiedAt DATETIME DEFAULT GETDATE()  -- Set modification date to current date and time
);

-- Insert dummy data into Users table
INSERT INTO Users (firstName, lastName, email, createdAt, modifiedAt)
VALUES 
('John', 'Doe', 'john.doe@example.com', GETDATE(), GETDATE()),
('Jane', 'Smith', 'jane.smith@example.com', GETDATE(), GETDATE()),
('Emily', 'Johnson', 'emily.johnson@example.com', GETDATE(), GETDATE()),
('Michael', 'Brown', 'michael.brown@example.com', GETDATE(), GETDATE()),
('Sarah', 'Davis', 'sarah.davis@example.com', GETDATE(), GETDATE());

-- Insert dummy data into Merchandise table
INSERT INTO Merchandise (name, price, availability, createdAt, modifiedAt)
VALUES 
('T-shirt', 19.99, 1, GETDATE(), GETDATE()),
('Hoodie', 49.99, 1, GETDATE(), GETDATE()),
('Coffee Mug', 12.50, 1, GETDATE(), GETDATE()),
('Sticker Pack', 4.99, 1, GETDATE(), GETDATE()),
('Baseball Cap', 24.99, 0, GETDATE(), GETDATE());  -- Not available

