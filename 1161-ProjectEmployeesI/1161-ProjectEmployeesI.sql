-- Last updated: 11/7/2025, 2:47:03 PM

select project_id, 
       round(avg(experience_years)::numeric, 2) as average_years
from Project
join Employee using (employee_id)
group by project_id;