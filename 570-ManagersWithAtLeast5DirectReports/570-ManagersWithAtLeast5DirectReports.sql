-- Last updated: 11/7/2025, 2:47:45 PM
-- Write your PostgreSQL query statement below
select s.name
from Employee as s
join Employee as e on s.id = e.managerId
group by s.id, s.name
having count(e.managerId) > 4