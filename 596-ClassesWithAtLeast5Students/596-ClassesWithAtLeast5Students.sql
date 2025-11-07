-- Last updated: 11/7/2025, 2:47:39 PM
select class
from Courses
group by class
having count(student) > 4