-- Last updated: 11/7/2025, 2:46:57 PM
SELECT DISTINCT ON (author_id) author_id AS id
FROM Views
WHERE author_id = viewer_id
ORDER BY author_id ASC;