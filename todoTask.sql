CREATE OR REPLACE FUNCTION getOverdueTasks(userId int) RETURNS TABLE (
    id int,
    name varchar(20),
    description text,
    done bool,
    is_active bool,
    deadline timestamp,
    user_id int
)
LANGUAGE plpgsql
    AS
$$
begin
    return query SELECT t.id, t.name, t.description, t.done, t.is_active, t.deadline, t.user_id
        FROM tasks as t
            WHERE t.user_id = $1 AND t.deadline < current_date;
end;
$$;

SELECT * from getOverdueTasks(1) order by id;
