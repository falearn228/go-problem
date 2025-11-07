-- Last updated: 11/7/2025, 2:47:38 PM
-- Write your PostgreSQL query statement below
with ranked_people_count as (
    select id, visit_date, people,
    row_number() over (order by id) as rn
    from Stadium
    where people >= 100
), grouped as (
    select r.id, r.visit_date, r.people, r.rn, r.id - r.rn as grp
    from ranked_people_count as r
), sequence as (
    select grp, count(*) as length
    from grouped
    group by grp
    having count(*) >= 3
) 
select g.id, g.visit_date, g.people
from grouped as g
join sequence as s
    on g.grp = s.grp
order by g.visit_date