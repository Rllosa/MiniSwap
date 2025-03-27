// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import "./IERC20.sol";
import "./ReentrancyGuard.sol";

contract MiniSwap is ReentrancyGuard {
    IERC20 public token1;
    IERC20 public token2;
    uint256 public rate;  // Exchange rate with 3 decimals (1000 = 1:1)
    uint256 public constant FEE_DENOMINATOR = 1000;
    uint256 public feeNumerator = 2;      // 2/1000 = 0.2%
    address public owner;

    event Swap(
        address indexed user,
        address indexed tokenIn,
        address indexed tokenOut,
        uint256 amountIn,
        uint256 amountOut,
        uint256 feeAmount
    );

    event FeeUpdated(uint256 oldFee, uint256 newFee);
    event RateUpdated(uint256 oldRate, uint256 newRate);

    constructor(address _token1, address _token2) {
        require(_token1 != address(0) && _token2 != address(0), "Invalid token addresses");
        token1 = IERC20(_token1);
        token2 = IERC20(_token2);
        owner = msg.sender;
        rate = 1000;  // Initial 1:1 rate (1000 = 1.000)
    }

    function setFee(uint256 _newFeeNumerator) external {
        require(_newFeeNumerator <= 50, "Fee cannot exceed 5%");
        emit FeeUpdated(feeNumerator, _newFeeNumerator);
        feeNumerator = _newFeeNumerator;
    }

    function setRate(uint256 _newRate) external {
        require(_newRate > 0, "Rate must be greater than 0");
        emit RateUpdated(rate, _newRate);
        rate = _newRate;
    }

    function calculateFee(uint256 amount) public view returns (uint256) {
        return (amount * feeNumerator) / FEE_DENOMINATOR;
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

        // Calculate amount out after fee
        uint256 fee = calculateFee(amountIn);
        uint256 amountAfterFee = amountIn - fee;
        
        // Apply exchange rate (rate has 3 decimals, so divide by 1000)
        uint256 amountOut;
        if (tokenIn == address(token1)) {
            amountOut = (amountAfterFee * rate) / 1000;
        } else {
            amountOut = (amountAfterFee * 1000) / rate;
        }

        require(
            tokenOutContract.balanceOf(address(this)) >= amountOut,
            "Insufficient liquidity"
        );

        // Transfer all tokens in (including fee)
        require(
            tokenInContract.transferFrom(msg.sender, address(this), amountIn),
            "TransferFrom failed"
        );

        // Transfer amount minus fee to user
        require(
            tokenOutContract.transfer(msg.sender, amountOut),
            "Transfer failed"
        );

        emit Swap(msg.sender, tokenIn, tokenOut, amountIn, amountOut, fee);
    }
}