// import logo from './logo.svg';
import './App.css';

import React from "react";
import Plot from "plotly.js";
import "bootstrap/dist/css/bootstrap.min.css"

// import FileUpload from "./components/FileUpload"

// function App() {
//     return (
//         <div className="App">
//             <header className="App-header">
//                 <Plot
//                     data={[
//                         {
//                             x: [1, 2, 3],
//                             y: [2, 6, 3],
//                             type: 'scatter',
//                             mode: 'lines+markers',
//                             marker: {color: 'red'},
//                         },
//                         {type: 'bar', x: [1, 2, 3], y: [2, 5, 3]},
//                     ]}
//                     layout={ {width: 320, height: 240, title: 'A Fancy Plot'} }
//                 />
//             </header>
//         </div>
//     );
// }

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
