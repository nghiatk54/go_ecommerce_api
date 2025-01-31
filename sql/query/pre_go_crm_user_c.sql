-- name: GetUserByEmailSqlc :one
SELECT `usr_email`, `usr_id` FROM `pre_go_crm_user_c` WHERE `usr_email` = sqlc.arg(email) LIMIT 1;

-- name: UpdateUserStatusByUserIdSqlc :exec
UPDATE `pre_go_crm_user_c` 
SET `usr_status` = sqlc.arg(status),
    `usr_updated_at` = sqlc.arg(updatedAt)
WHERE `usr_id` = sqlc.arg(userId);