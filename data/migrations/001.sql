CREATE TABLE test_time (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    val DATETIME DEFAULT (datetime('now')),
    val2 DATETIME not NULL DEFAULT (datetime('now'))
);
