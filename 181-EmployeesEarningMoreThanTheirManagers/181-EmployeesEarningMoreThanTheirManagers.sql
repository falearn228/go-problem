-- Last updated: 11/7/2025, 2:48:32 PM
-- Write your PostgreSQL query statement below
-- id.salary > managerId.salary
select a.Name as Employee 
from Employee a join Employee b on a.ManagerId = b.Id
where a.Salary > b.Salary