-- Last updated: 11/7/2025, 2:46:18 PM
-- Write your PostgreSQL query statement below
with primary_flags as (
    select employee_id, department_id
    from Employee
    where primary_flag = 'Y'
)

 select distinct employee_id, coalesce(p.department_id, e1.department_id) as department_id 
 from Employee as e1
 left join primary_flags as p using (employee_id)
 order by employee_id