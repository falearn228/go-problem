-- Last updated: 11/7/2025, 2:46:21 PM
-- Write your PostgreSQL query statement below
select e1.employee_id, e1.name, count(*) as reports_count , round(avg(e2.age)) as average_age 
from Employees as e1
join Employees as e2 on e1.employee_id = e2.reports_to
group by e1.employee_id, e1.name
order by e1.employee_id