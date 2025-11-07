-- Last updated: 11/7/2025, 2:48:30 PM
-- Write your PostgreSQL query statement below
with ranked_salary as (
    select id, name, salary ,rank() over (partition by departmentId order by salary desc) as r_salary, departmentId
    from Employee
)

select d.name as Department, e.name as Employee, e.salary
from ranked_salary as e
join Department as d on d.id = e.departmentId
where e.r_salary = 1