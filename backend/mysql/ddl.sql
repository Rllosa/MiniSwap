CREATE DATABASE IF NOT EXISTS MiniSwap;
USE MiniSwap;

-- Table to track the latest processed block
CREATE TABLE IF NOT EXISTS blockInfo (
    LatestBlockNum BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS swap_events (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_address VARCHAR(42) NOT NULL,
    token_in VARCHAR(42) NOT NULL,
    token_out VARCHAR(42) NOT NULL,
    amount_in DECIMAL(65,0) NOT NULL,
    amount_out DECIMAL(65,0) NOT NULL,
    fees DECIMAL(65,0) NOT NULL,
    created_at BIGINT NOT NULL,
    INDEX idx_user (user_address),
    INDEX idx_created_at (created_at)
);

-- Table to store burning events
CREATE TABLE IF NOT EXISTS burning_events (
    id INT PRIMARY KEY AUTO_INCREMENT,
    token_address VARCHAR(42) NOT NULL,
    amount_burnt DECIMAL(65,0) NOT NULL,
    pool_remaining DECIMAL(65,0) NOT NULL,
    created_at BIGINT NOT NULL,
    INDEX idx_token (token_address),
    INDEX idx_created_at (created_at)
);
