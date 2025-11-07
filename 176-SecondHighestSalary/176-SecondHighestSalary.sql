-- Last updated: 11/7/2025, 2:48:36 PM
-- Write your PostgreSQL query statement below
select (
    select salary from (
        select salary, dense_rank() over (order by salary desc) as rank
        from employee
    )
    where rank = 2
    limit 1
) as SecondHighestSalary 