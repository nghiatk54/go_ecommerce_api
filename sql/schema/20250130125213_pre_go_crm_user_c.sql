-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `pre_go_crm_user_c` (
    `usr_id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Account ID',
    `usr_email` varchar(30) NOT NULL DEFAULT '' COMMENT 'Email',
    `usr_phone` varchar(15) NOT NULL DEFAULT '' COMMENT 'Phone number',
    `usr_username` varchar(255) NOT NULL DEFAULT '' COMMENT 'Username',
    `usr_password` varchar(255) NOT NULL DEFAULT '' COMMENT 'Password',
    `usr_created_at` int(11) NOT NULL DEFAULT 0 COMMENT 'Creation time',
    `usr_updated_at` int(11) NOT NULL DEFAULT 0 COMMENT 'Update time',
    `usr_create_ip` varchar(12) NOT NULL DEFAULT '' COMMENT 'Creation IP',
    `usr_last_login_at` int(11) NOT NULL DEFAULT 0 COMMENT 'Last login time',
    `usr_last_login_ip` varchar(12) NOT NULL DEFAULT '' COMMENT 'Last login IP',
    `usr_login_times` int(11) NOT NULL DEFAULT 0 COMMENT 'Login times',
    `usr_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'Status 1:enable, 0:disable, -1:deleted',
    PRIMARY KEY (`usr_id`),
    KEY `idx_usr_email` (`usr_email`),
    KEY `idx_usr_phone` (`usr_phone`),
    KEY `idx_usr_username` (`usr_username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT="Account";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `pre_go_crm_user_c`;
-- +goose StatementEnd
