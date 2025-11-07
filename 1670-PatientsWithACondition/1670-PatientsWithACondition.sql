-- Last updated: 11/7/2025, 2:46:31 PM
SELECT *
FROM Patients
WHERE conditions LIKE 'DIAB1%'           -- начало строки
   OR conditions LIKE '% DIAB1%'         -- начало слова (после пробела)
   OR conditions = 'DIAB1'               -- точное совпадение