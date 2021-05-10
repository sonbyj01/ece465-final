import logo from './logo.svg';
import './App.css';

import React, { Component } from "react";
import "bootstrap/dist/css/bootstrap.min.css"

import FileUpload from "./components/FileUpload"

function App() {
  return (
      <div className={"container"} style={{ width: "600px" }}>
        <FileUpload />
      </div>

    // <div className="App">
    //   <header className="App-header">
    //     <img src={logo} className="App-logo" alt="logo" />
    //     <p>
    //       Edit <code>src/App.js</code> and save to reload.
    //     </p>
    //     <a
    //       className="App-link"
    //       href="https://reactjs.org"
    //       target="_blank"
    //       rel="noopener noreferrer"
    //     >
    //       Learn React
    //     </a>
    //   </header>
    //   <UploadUI />
    // </div>
  );
}

export default App;
