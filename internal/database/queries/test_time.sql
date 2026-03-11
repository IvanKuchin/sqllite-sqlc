-- name: InsertTestTime :one
INSERT INTO test_time (val, val2) VALUES (?, ?)
RETURNING id;

-- name: GetTestTimeByID :one
SELECT * FROM test_time WHERE id = ? LIMIT 1;
