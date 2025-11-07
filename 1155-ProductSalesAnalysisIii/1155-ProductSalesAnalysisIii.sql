-- Last updated: 11/7/2025, 2:47:04 PM
SELECT 
    product_id,
    year AS first_year,
    quantity,
    price
FROM (
    SELECT 
        product_id,
        year,
        quantity,
        price,
        RANK() OVER (PARTITION BY product_id ORDER BY year) AS year_rank
    FROM Sales
) ranked_sales
WHERE year_rank = 1;