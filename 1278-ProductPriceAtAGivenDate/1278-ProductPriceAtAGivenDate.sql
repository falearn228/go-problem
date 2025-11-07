-- Last updated: 11/7/2025, 2:46:55 PM
with last_changes as (
    select distinct on (product_id) product_id,
    new_price
    from Products
    where change_date <= '2019-08-16'
    order by product_id, change_date desc
)

select product_id,
coalesce(new_price, 10) as price
from (select distinct product_id from Products)
left join last_changes using (product_id)