-- file: pre_go_acc_user_two_factor_9999.sql

-- Enalbled two factor
-- name: EnabledTwoFactorTypeEmail :exec
INSERT INTO `pre_go_acc_user_two_factor_9999` (
    user_id,
    two_factor_auth_type,
    two_factor_auth_secret,
    two_factor_email,
    two_factor_is_active,
    two_factor_created_at,
    two_factor_updated_at
) VALUES (
    ?,
    ?,
    "OTP",
    ?,
    FALSE,
    NOW(),
    NOW()
);

-- Disable two factor
-- name: DisableTwoFactor :exec
UPDATE `pre_go_acc_user_two_factor_9999` SET
    two_factor_is_active = TRUE,
    two_factor_updated_at = NOW()
WHERE user_id = ? AND two_factor_auth_type = ?;

-- Update two factor status verification
-- name: UpdateTwoFactorStatus :exec
UPDATE `pre_go_acc_user_two_factor_9999` SET
    two_factor_is_active = TRUE,
    two_factor_updated_at = NOW()
WHERE user_id = ? AND two_factor_auth_type = ? AND two_factor_is_active = FALSE;

-- Verify two factor
-- name: VerifyTwoFactor :one
SELECT COUNT(*)
FROM `pre_go_acc_user_two_factor_9999`
WHERE user_id = ? AND two_factor_auth_type = ? AND two_factor_is_active = TRUE;

-- Get two factor status
-- name: GetTwoFactorStatus :one
SELECT two_factor_is_active
FROM `pre_go_acc_user_two_factor_9999`
WHERE user_id = ? AND two_factor_auth_type = ?;

-- Is two factor enabled
-- name: IsTwoFactorEnabled :one
SELECT COUNT(*)
FROM `pre_go_acc_user_two_factor_9999`
WHERE user_id = ? AND two_factor_is_active = TRUE;

-- Add or update phone number
-- name: AddOrUpdatePhoneNumber :exec
INSERT INTO `pre_go_acc_user_two_factor_9999` (
    user_id,
    two_factor_phone,
    two_factor_is_active
) VALUES (
    ?,
    ?,
    TRUE
)
ON DUPLICATE KEY UPDATE
    two_factor_phone = ?,
    two_factor_updated_at = NOW();

-- add or update email
-- name: AddOrUpdateEmail :exec
INSERT INTO `pre_go_acc_user_two_factor_9999` (
    user_id,
    two_factor_email,
    two_factor_is_active
) VALUES (?, ?, TRUE)
ON DUPLICATE KEY UPDATE
    two_factor_email = ?,
    two_factor_updated_at = NOW();

-- get user two factor methods
-- name: GetUserTwoFactorMethods :many
SELECT two_factor_id, user_id, two_factor_auth_type, two_factor_auth_secret,
        two_factor_phone, two_factor_email, two_factor_is_active,
        two_factor_created_at, two_factor_updated_at
FROM `pre_go_acc_user_two_factor_9999`
WHERE user_id = ?;

-- reactive two factor
-- name: ReactiveTwoFactor :exec
UPDATE `pre_go_acc_user_two_factor_9999` SET
    two_factor_is_active = TRUE,
    two_factor_updated_at = NOW()
WHERE user_id = ? AND two_factor_auth_type = ?;

-- remove two factor
-- name: RemoveTwoFactor :exec
DELETE FROM `pre_go_acc_user_two_factor_9999`
WHERE user_id = ? AND two_factor_auth_type = ?;

-- count active two factor methods
-- name: CountActiveTwoFactorMethods :one
SELECT COUNT(*)
FROM `pre_go_acc_user_two_factor_9999`
WHERE user_id = ? AND two_factor_is_active = TRUE;

-- get two factor method by id
-- name: GetTwoFactorMethodById :one
SELECT two_factor_id, user_id, two_factor_auth_type, two_factor_auth_secret,
        two_factor_phone, two_factor_email, two_factor_is_active,
        two_factor_created_at, two_factor_updated_at
FROM `pre_go_acc_user_two_factor_9999`
WHERE two_factor_id = ?;

-- get two factor method by user id and auth type
-- name: GetTwoFactorMethodByUserIdAndAuthType :one
SELECT two_factor_id, user_id, two_factor_auth_type, two_factor_auth_secret,
        two_factor_phone, two_factor_email, two_factor_is_active,
        two_factor_created_at, two_factor_updated_at
FROM `pre_go_acc_user_two_factor_9999`
WHERE user_id = ? AND two_factor_auth_type = ?;