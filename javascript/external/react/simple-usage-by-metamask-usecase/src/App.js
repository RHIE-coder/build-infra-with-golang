import { useEffect, useState } from 'react';
import './App.css';
import { MetaMaskSDK } from '@metamask/sdk';
import detectEthereumProvider from '@metamask/detect-provider';

function App() {
  const initialState = { accounts: [] }               
  const [wallet, setWallet] = useState(initialState)  
  useEffect(()=>{
    if (typeof window.ethereum !== 'undefined') {
      // alert('MetaMask is installed!');
    } else {
      alert('MetaMask is not installed. Please install it')
    }
  }, [])

  const connect = async () => {
    const options = {

    }
    const MMSDK = new MetaMaskSDK(options);
    const provider = await detectEthereumProvider();
    const result = await ethereum.request({ method: 'eth_requestAccounts', params: [] });
    console.log(result)
  }

  return (
    <div className="App">
      <h1>Hello, Metamask</h1>
      <button onClick={connect}>메타마스크 연결</button>
    </div>
  );
}

export default App;
