-- Last updated: 11/7/2025, 2:46:54 PM
with first_order as (
    select customer_id,
    order_date,
    customer_pref_delivery_date,
    row_number() over (PARTITION by customer_id order by order_date) as order_rank
    from Delivery
)

SELECT
    ROUND(
        100.0 * COUNT(CASE WHEN order_date = customer_pref_delivery_date THEN 1 END) / COUNT(*), 
        2
    ) AS immediate_percentage
FROM
    first_order
where order_rank = 1