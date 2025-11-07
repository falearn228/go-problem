-- Last updated: 11/7/2025, 2:48:25 PM
SELECT 
    today.id 
FROM 
    Weather AS today
JOIN 
    Weather AS yesterday 
    ON today.recordDate = DATE_ADD(yesterday.recordDate, INTERVAL '1 day')
WHERE 
    today.temperature > yesterday.temperature;