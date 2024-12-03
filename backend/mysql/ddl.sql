CREATE DATABASE IF NOT EXISTS dexbackend;
USE dexbackend;

-- Table to track the latest processed block
CREATE TABLE IF NOT EXISTS blockInfo (
    LatestBlockNum BIGINT NOT NULL
);

-- Table to store swap events
CREATE TABLE IF NOT EXISTS swap_events (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_address VARCHAR(42) NOT NULL,
    token_in VARCHAR(42) NOT NULL,
    token_out VARCHAR(42) NOT NULL,
    amount_in DECIMAL(65,0) NOT NULL,
    amount_out DECIMAL(65,0) NOT NULL,
    transaction_hash VARCHAR(66) NOT NULL,
    block_number BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_user (user_address),
    INDEX idx_tokens (token_in, token_out),
    INDEX idx_block (block_number)
);

-- Table to store burning events
CREATE TABLE IF NOT EXISTS burning_events (
    id INT PRIMARY KEY AUTO_INCREMENT,
    token_address VARCHAR(42) NOT NULL,
    amount_burnt DECIMAL(65,0) NOT NULL,
    pool_remaining DECIMAL(65,0) NOT NULL,
    transaction_hash VARCHAR(66) NOT NULL,
    block_number BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_token (token_address),
    INDEX idx_block (block_number)
);