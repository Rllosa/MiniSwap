// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import "./IERC20.sol";
import "./ReentrancyGuard.sol";

contract MiniSwap is ReentrancyGuard {
    IERC20 public token1;
    IERC20 public token2;
    uint256 public rate = 1; // 1:1 rate for simplicity

    event Swap(
        address indexed user,
        address indexed tokenIn,
        address indexed tokenOut,
        uint256 amountIn,
        uint256 amountOut
    );

    constructor(address _token1, address _token2) {
        require(_token1 != address(0) && _token2 != address(0), "Invalid token addresses");
        token1 = IERC20(_token1);
        token2 = IERC20(_token2);
    }

    function addLiquidity(
        address token,
        uint256 amount
    ) external {
        require(
            token == address(token1) || token == address(token2),
            "Invalid token"
        );
        
        IERC20(token).transferFrom(msg.sender, address(this), amount);
    }

    function swap(
        address tokenIn,
        uint256 amountIn
    ) external nonReentrant {
        require(
            tokenIn == address(token1) || tokenIn == address(token2),
            "Invalid token"
        );
        
        address tokenOut = tokenIn == address(token1) ? address(token2) : address(token1);
        IERC20 tokenInContract = IERC20(tokenIn);
        IERC20 tokenOutContract = IERC20(tokenOut);
        
        require(amountIn > 0, "Amount must be greater than 0");
        require(
            tokenInContract.balanceOf(msg.sender) >= amountIn,
            "Insufficient balance"
        );
        require(
            tokenInContract.allowance(msg.sender, address(this)) >= amountIn,
            "Insufficient allowance"
        );

        uint256 amountOut = amountIn * rate;
        require(
            tokenOutContract.balanceOf(address(this)) >= amountOut,
            "Insufficient liquidity"
        );

        require(
            tokenInContract.transferFrom(msg.sender, address(this), amountIn),
            "TransferFrom failed"
        );
        require(
            tokenOutContract.transfer(msg.sender, amountOut),
            "Transfer failed"
        );

        emit Swap(msg.sender, tokenIn, tokenOut, amountIn, amountOut);
    }
}