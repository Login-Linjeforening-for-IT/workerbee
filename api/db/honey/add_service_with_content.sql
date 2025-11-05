INSERT INTO honey (
    service,
    language,
    page,
    text
) VALUES (
    :service,
    :language,
    :page,
    :text
)
RETURNING *;