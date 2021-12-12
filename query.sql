SELECT
    u1.id as ID,
    u1.UserName,
    u2.UserName as ParentUserName
FROM
    USER u1
        LEFT JOIN USER u2
                  ON u1.Parent = u2.id
ORDER BY u1.id ASC