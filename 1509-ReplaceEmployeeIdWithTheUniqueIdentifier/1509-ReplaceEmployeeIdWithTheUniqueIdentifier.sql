-- Last updated: 11/7/2025, 2:46:40 PM
-- Write your PostgreSQL query statement below
select u.unique_id, e.name 
from Employees as e left join EmployeeUNI as u
on e.id = u.id
