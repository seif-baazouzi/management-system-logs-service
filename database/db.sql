CREATE DATABASE logs;

\c logs

CREATE TABLE logs (
    logID UUID PRIMARY KEY,
    label VARCHAR NOT NULL, 
    description VARCHAR NOT NULL, 
    value INT NOT NULL, 
    date TIMESTAMP NOT NULL, 
    userID UUID NOT NULL,
    workspaceID UUID NOT NULL,
    createdAt TIMESTAMP DEFAULT NOW() 
);

CREATE INDEX UserIdInex ON logs USING hash (userID);
CREATE INDEX WorkspaceInex ON logs USING hash (workspaceID);
