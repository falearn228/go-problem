-- Last updated: 11/7/2025, 2:48:33 PM
select distinct num as ConsecutiveNums
from (
    select
    num,
    lag(num, 1) over (order by id) = num and
    lag(num, 2) over (order by id) = num as porno
    from Logs 
)
where porno