-- Last updated: 11/7/2025, 2:46:14 PM
-- Write your PostgreSQL query statement below
select employee_id
from Employees
where salary < 30000 and manager_id not in (
    select employee_id
    from Employees
)
order by employee_id asc;