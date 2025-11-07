-- Last updated: 11/7/2025, 2:46:32 PM
SELECT *
FROM Users
WHERE mail ~ '^[A-Za-z][A-Za-z0-9_.-]*@leetcode\.com$'