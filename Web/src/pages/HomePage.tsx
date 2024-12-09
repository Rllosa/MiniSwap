import React, { useState } from 'react';
import './HomePage.css';

const HomePage = () => {
  const [walletConnected, setWalletConnected] = useState<boolean>(false);
  const [tokenA, setTokenA] = useState<string>('');
  const [tokenB, setTokenB] = useState<string>('');
  const [amount, setAmount] = useState<string>('');
  const [status, setStatus] = useState<string>('');

  const connectWallet = () => {
    setWalletConnected(true);
    setStatus('Wallet connected (simulation)');
  };

  const handleSwap = () => {
    if (!walletConnected) {
      setStatus('Please connect your wallet first.');
      return;
    }

    if (!tokenA || !tokenB || !amount) {
      setStatus('Please fill all fields.');
      return;
    }

    setStatus(`Simulating swap: ${amount} of Token A (${tokenA}) to Token B (${tokenB})`);
    setTimeout(() => {
      setStatus('Swap simulation complete!');
    }, 1500);
  };

  return (
    <div className="home-container">
      <h1 className="home-title">Token Swap</h1>
      <p className="home-status">{status}</p>

      {!walletConnected ? (
        <button className="home-button" onClick={connectWallet}>
          Connect Wallet
        </button>
      ) : (
        <div className="swap-form">
          <p className="wallet-status">Wallet connected (simulated).</p>
          <input
            type="text"
            placeholder="Token A Address"
            value={tokenA}
            onChange={(e) => setTokenA(e.target.value)}
            className="swap-input"
          />
          <input
            type="text"
            placeholder="Token B Address"
            value={tokenB}
            onChange={(e) => setTokenB(e.target.value)}
            className="swap-input"
          />
          <input
            type="number"
            placeholder="Amount to Swap"
            value={amount}
            onChange={(e) => setAmount(e.target.value)}
            className="swap-input"
          />
          <button className="home-button" onClick={handleSwap}>
            Swap Tokens
          </button>
        </div>
      )}
    </div>
  );
};

export default HomePage;
