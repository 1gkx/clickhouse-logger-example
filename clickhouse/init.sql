CREATE DATABASE IF NOT EXISTS Test;
USE Test;

CREATE TABLE IF NOT EXISTS test (
                                    body String,
                                    created_at DateTime,
                                    id UInt64
) Engine = MergeTree()
      ORDER BY (created_at);
