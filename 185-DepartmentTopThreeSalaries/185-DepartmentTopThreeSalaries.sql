-- Last updated: 11/7/2025, 2:48:29 PM
-- Write your PostgreSQL query statement below
with ranked_salary as (
    select name, salary, departmentId, dense_rank() over (partition by departmentId order by salary desc) as ranked_s
    from Employee
)

select d.name as Department, e.name as Employee, e.salary as Salary
from ranked_salary e
join Department d on d.id = e.departmentId
where ranked_s < 4