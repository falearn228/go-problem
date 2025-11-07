-- Last updated: 11/7/2025, 2:46:06 PM
-- Write your PostgreSQL query statement below
select teacher_id, count(distinct subject_id) as cnt
from Teacher
group by teacher_id