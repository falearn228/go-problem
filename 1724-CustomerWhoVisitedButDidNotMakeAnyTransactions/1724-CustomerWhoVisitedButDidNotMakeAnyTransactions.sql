-- Last updated: 11/7/2025, 2:46:30 PM
-- Write your PostgreSQL query statement below
select v.customer_id, count(v.visit_id) as count_no_trans
from Visits as v left join Transactions as t on v.visit_id = t.visit_id
where t.transaction_id is null
group by v.customer_id



-- left join Transactions as t on v.visit_id = t.visit_id