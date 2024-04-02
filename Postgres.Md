# Postgress

### Token

In PostgreSQL, a token generally refers to a sequence of characters that represent a specific element in SQL, such as a keyword, identifier, literal value, or operator. Tokens are the basic building blocks of an SQL statement. For example, in the SQL statement SELECT * FROM users WHERE id = 1;, the tokens are SELECT, *, FROM, users, WHERE, id, =, and 1.

### CTID

In PostgreSQL, CTID is a pseudo-column that represents the physical location of a tuple within its table. It stands for "tuple identifier." Each row (tuple) in a PostgreSQL table has a unique CTID that identifies its physical location on disk. The CTID is a tuple's physical address and is useful in certain scenarios, such as when you need a unique identifier for a row that doesn't change, even if the row is updated.

### Primary key

In PostgreSQL, a primary key is a column or a group of columns that uniquely identifies each row in a table.meaning that the primary key column(s) must contain unique values and cannot be null.

### Transactions

A transaction is the propagation of one or more changes to the database. For example, if you are creating a record, updating a record, or deleting a record from the table, then you are performing transaction on the table.

Transactions have the following four standard properties, usually referred to by the acronym ACID −

- Atomicity
- Consistency
- Isolation
- Durability

Transaction Control

- BEGIN TRANSACTION − To start a transaction.
- COMMIT − To save the changes, alternatively you can use END TRANSACTION command.
- ROLLBACK − To rollback the changes.

### What is the difference between serial and sequence in postgress

In PostgreSQL, SERIAL and SEQUENCE are both used to create auto-incrementing integer columns, typically used for primary keys. However, they differ in how they are implemented and used:

SERIAL:

SERIAL is a shorthand notation for creating an auto-incrementing column.
When you define a column as SERIAL, PostgreSQL internally creates a sequence object and sets the column's default value to be generated from the next value of the sequence.
SERIAL columns are commonly used for primary keys and are often used when you want a simple, auto-incrementing integer primary key.
Example:
Copy code
CREATE TABLE example_table (
    id SERIAL PRIMARY KEY,
    name TEXT
);

SEQUENCE:

A SEQUENCE is an independent database object that generates a sequence of integers.
Unlike SERIAL, you explicitly create and manage the sequence object, and then use it to populate the column values as needed.
Using a sequence directly gives you more control over how the sequence is used and allows you to use the same sequence for multiple tables if desired.
Example:
CREATE SEQUENCE my_sequence START 1;
CREATE TABLE example_table (
    id INTEGER PRIMARY KEY DEFAULT nextval('my_sequence'),
    name TEXT
);
In summary, SERIAL is a shorthand way to create an auto-incrementing column with a sequence behind the scenes, while SEQUENCE is a standalone object that you can use to generate unique values for columns manually.

### How do you handle concurrent updates in PostgreSQL?

PostgreSQL handles concurrent updates through its multi-version concurrency control (MVCC) mechanism. MVCC allows multiple transactions to access the same data simultaneously without blocking each other or causing conflicts.

When two transactions attempt to modify the same data concurrently, PostgreSQL ensures isolation by creating separate copies of the data for each transaction. This way, each transaction sees a consistent snapshot of the data as it appeared at the beginning of the transaction.

In cases where conflicts can occur, PostgreSQL provides different isolation levels and locking mechanisms to manage concurrent updates. Developers can choose appropriate transaction isolation levels, such as READ COMMITTED, REPEATABLE READ, or SERIALIZABLE, based on their application's requirements and trade-offs between concurrency and data consistency.

### Windows function

In PostgreSQL, window functions are a powerful feature that allow you to perform calculations across a set of rows related to the current row, without the need for self-joins or subqueries. Window functions are used in conjunction with the OVER clause and can be used to calculate running totals, rankings, averages, and more.

### Stored Procedure

In PostgreSQL, a stored procedure is a pre-compiled and stored database object that encapsulates a set of SQL statements. It allows you to perform complex operations and execute them as a single unit. The benefits of using stored procedures include code reusability, improved performance, and enhanced security

### Error Handler

In PostgreSQL, you can handle errors using exception blocks in PL/pgSQL, which is the procedural language used in PostgreSQL for writing functions and stored procedures.

DO $$
BEGIN
    -- Perform some operations
    RAISE EXCEPTION 'This is a test exception';
EXCEPTION
    WHEN OTHERS THEN
        -- Handle the exception
        RAISE NOTICE 'An error occurred: %', SQLERRM;
END $$;

In this example, the RAISE EXCEPTION statement is used to deliberately raise an exception. The EXCEPTION block then catches any exception that occurs within the block. You can replace RAISE NOTICE with RAISE EXCEPTION to re-raise the exception and stop further execution, or you can perform other actions such as logging the error, rolling back a transaction, or returning a specific error message to the client application.

### Difference between Union and Union ALL

UNION: The UNION operator is used to combine the result sets of two or more SELECT statements and remove duplicate rows from the final result set.

UNION ALL: The UNION ALL operator is used to combine the result sets of two or more SELECT statements, including all rows (including duplicates) from all SELECT statements.

Ex: SELECT column1 FROM table1 UNION ALL SELECT column1 FROM table2;

### How do you use EXPLAIN in PostgreSQL to analyze query execution plans?

In PostgreSQL, you can use the EXPLAIN command to analyze the execution plan of a query without actually executing it. This helps in understanding how the query will be processed by the database engine and allows for query optimization. Here's an example:

EXPLAIN query;

### How do you implement table partianing in postgress

Table partitioning in PostgreSQL allows you to split a large table into smaller, more manageable parts called partitions. Each partition is a separate table with its own storage characteristics, but from the application's perspective, it appears as a single table. 

```
CREATE TABLE measurement (
    city_id         int not null,
    logdate         date not null,
    peaktemp        int,
    unitsales       int
) PARTITION BY RANGE (logdate);

CREATE TABLE measurement_2016 PARTITION OF measurement
    FOR VALUES FROM ('2016-01-01') TO ('2017-01-01');

CREATE TABLE measurement_2017 PARTITION OF measurement
    FOR VALUES FROM ('2017-01-01') TO ('2018-01-01');

CREATE TABLE measurement_2018 PARTITION OF measurement
    FOR VALUES FROM ('2018-01-01') TO ('2019-01-01');

INSERT INTO measurement (city_id, logdate, peaktemp, unitsales)
VALUES (1, '2016-01-01', 55, 10);

SELECT * FROM measurement WHERE logdate >= '2016-01-01' AND logdate < '2017-01-01';
```

### What is the role of VACUUM in PostgreSQL?

In PostgreSQL, VACUUM is a crucial process for managing and reclaiming disk space occupied by deleted or outdated data. It performs two main tasks:

Freeing up disk space: When data is updated or deleted in PostgreSQL, it is not immediately removed from the disk. Instead, it becomes marked as reusable by future inserts. The VACUUM process identifies these reusable data pages and makes the space available for future use.

Updating statistics: VACUUM analyzes the distribution of data and updates the statistics used by the query planner. Accurate statistics help PostgreSQL choose optimal query plans, resulting in improved performance.

VACUUM is essential for maintaining the performance and efficiency of PostgreSQL databases, especially in scenarios with frequent updates and deletes.

### How do you use advisory locks in PostgreSQL?

Advisory locks in PostgreSQL are user-defined locks that allow coordination between database
sessions. Unlike regular row-level locks, advisory locks do not conflict with other types of locks and can be acquired and released voluntarily.

To acquire an advisory lock, use the pg_advisory_lock function:

SELECT pg_advisory_lock(lock_id);

Replace lock_id with a unique identifier for the lock. Multiple sessions can acquire the same advisory lock by using the same lock_id.

To release an advisory lock, use the pg_advisory_unlock function:

SELECT pg_advisory_unlock(lock_id);

Advisory locks are often used for application-level synchronization or to implement custom concurrency control mechanisms.

### Explain the use of PostgreSQL triggers.

A trigger can be defined as a function that is called automatically when the insertion, updation, or deletion event occurs. They serve as a way to check the data integrity. Triggers are capable of handling any errors that occur in the database. Another advantage of triggers is: Any table that is present in a PostgreSQL database can be forced to receive security approvals with the use of PostgreSQL triggers.

