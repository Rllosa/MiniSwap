import React, { useState } from "react";
import "./HomePage.css";

const HomePage = () => {
  const [walletConnected, setWalletConnected] = useState<boolean>(false);
  const [tokenAmount, setTokenAmount] = useState<number | undefined>(undefined);
  const [walletContent, setWalletContent] = useState<number | undefined>(undefined);
  const [wallet2Content, setWallet2Content] = useState<number | undefined>(undefined);
  const [status, setStatus] = useState<string>("");
  const [isEthToUsdt, setIsEthToUsdt] = useState<boolean>(true);

  const [message, setMessage] = useState<string | null>(null);
  const [messageType, setMessageType] = useState<"success" | "error" | null>(null);

  const showMessage = (msg: string, type: "success" | "error") => {
    setMessage(msg);
    setMessageType(type);

    setTimeout(() => {
      setMessage(null);
      setMessageType(null);
    }, 3000);
  };

  const connectWallet = async () => {
    try {
      const response = await fetch("http://localhost:8080/getBalance");
  
      if (!response.ok) {
        throw new Error(`Erreur ${response.status} : ${response.statusText}`);
      }
  
      const json = await response.json();
      setWalletContent(json.token1);
      setWallet2Content(json.amount);
  
      setWalletConnected(true);
      setStatus("Wallet connected (simulation)");
      showMessage("Wallet connected successfully!", "success");
    } catch (err) {
      console.log(err);
      showMessage("Failed to connect to the wallet", "error");
    }
  };

  const handleSwap = async () => {
    if (tokenAmount && walletContent !== undefined && wallet2Content !== undefined) {
      try {
        const fromToken = isEthToUsdt ? "base" : "quote";
        const response = await fetch(
          `http://localhost:8080/swap?amount=${tokenAmount}&token=${fromToken}`
        );
  
        if (!response.ok) {
          throw new Error(`Erreur ${response.status} : ${response.statusText}`);
        }
  
        const json = await response.json();
        setWalletContent(json.token1);
        setWallet2Content(json.token2);
        showMessage("Swap successful!", "success");
      } catch (err) {
        console.log(err);
        showMessage("Swap failed", "error");
      }
    }
  };

  const handleSwitch = () => {
    setIsEthToUsdt((prev) => !prev);
    setTokenAmount(undefined);
  };

  const getConversionValue = () => {
    if (!tokenAmount) return "";
    const fee = 0.002;
    return isEthToUsdt
      ? tokenAmount - tokenAmount * fee
      : tokenAmount - tokenAmount * fee;
  };

  const getMaxInput = () => {
    return isEthToUsdt ? walletContent : wallet2Content;
  };

  return (
    <div className="home-container">
    <h1 className="home-title">Token Swap</h1>
    <p className="home-status">{status}</p>

    {message && (
      <div className={`message ${messageType}`}>
        {message}
      </div>
    )}
  
    {!walletConnected ? (
      <button className="home-button" onClick={connectWallet}>
        Connect Wallet
      </button>
    ) : (
      <div className="swap-form">
        <button className="home-button switch-button" onClick={handleSwitch}>
          Switch to {isEthToUsdt ? "USDT → ETH" : "ETH → USDT"}
        </button>
  
        <p className="wallet-status">Wallet connected (simulated).</p>
  
        <p>{isEthToUsdt ? "ETH Amount Value" : "USDT Amount Value"}</p>
        <input
          type="number"
          placeholder="Amount"
          value={tokenAmount ?? ""}
          onChange={(e) => {
            const value = Number(e.target.value);
            const max = getMaxInput();
            if (!isNaN(value) && max !== undefined && value <= max) {
              setTokenAmount(value);
            }
          }}
          className="swap-input"
        />
  
        <p>{isEthToUsdt ? "USDT Amount conversion" : "ETH Amount conversion"}</p>
        <input
          type="number"
          placeholder="Converted Amount"
          readOnly
          value={getConversionValue()}
          className="swap-input"
        />
  
        <button className="home-button" onClick={handleSwap}>
          Swap Tokens
        </button>
      </div>
    )}
    {walletConnected && (
      <div className="wallet-summary">
        <div className="wallet-item">
          <span className="wallet-label">ETH:</span>
          <span className="wallet-value">{walletContent ?? "-"}</span>
        </div>
        <div className="wallet-item">
          <span className="wallet-label">USDT:</span>
          <span className="wallet-value">{wallet2Content ?? "-"}</span>
        </div>
      </div>
    )}
  </div>
  );
};

export default HomePage;
