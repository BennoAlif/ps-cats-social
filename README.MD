# Create migration
    
    ```bash
    migrate create -ext sql -dir db/migrations add_user_table
    
    ```
    
    This command will create two new files named `add_user_table.up.sql` and `add_user_table.down.sql` inside the `db/migrations` folder
    
    - `.up.sql` can be filled with database queries to create / delete / change the table
    - `.down.sql` can be filled with database queries to perform a `rollback` or return to the state before the table from `.up.sql` was created

# Execute migration
    
    ```bash
    migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations up
    
    ```
    
# Rollback migration
 ```bash
    migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations down
    
    ```
