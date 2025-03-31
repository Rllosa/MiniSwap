import React, { useState } from 'react';
import './HomePage.css';

const HomePage = () => {
  const [walletConnected, setWalletConnected] = useState<boolean>(false);
  const [tokenETH, setTokenETH] = useState<number | undefined>(undefined);
  const [walletContent, setWalletContent] = useState<number | undefined>(undefined);
  const [wallet2Content, setWallet2Content] = useState<number | undefined>(undefined);
  const [status, setStatus] = useState<string>('');

  const connectWallet = () => {
    setWalletConnected(true);
    setStatus('Wallet connected (simulation)');
  };

  const handleSwap = () => {
    // Placeholder for backend
    if (tokenETH) {
      setWalletContent(tokenETH);
      setWallet2Content((tokenETH - (0.002 * tokenETH)));
    }
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
          <p>ETH Amount Value</p>
          <input
            type="number"
            placeholder="ETH Amount"
            value={tokenETH ? tokenETH : ""}
            onChange={(e) => {
              const value = Number(e.target.value);
              if (!isNaN(value)) {
                setTokenETH(value);
              }
            }}
            className="swap-input"
          />
          <p>USDT Amount conversion</p>
          <input
            type="number"
            placeholder="USDT Amount"
            readOnly
            value={tokenETH ? (tokenETH - (0.002 * tokenETH)) : ""}
            className="swap-input"
          />
          <button className="home-button" onClick={handleSwap}>
            Swap Tokens
          </button>
        </div>
      )}
      <p>{walletContent ? "Your ETH wallet amount : " + walletContent : "We did not receive your ETH wallet infos."}</p>
      <p>{walletContent ? "Your USDT wallet amount : " + wallet2Content : "We did not receive your USDT wallet infos."}</p>
    </div>
  );
};

export default HomePage;
