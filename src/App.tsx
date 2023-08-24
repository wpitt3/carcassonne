import React, { useState } from 'react';
import './App.css';
import TileDisplay from './TileDisplay';

function App() {
  const [page, setPage] = useState < string > ('build');
  return (
    <div>
      <TileDisplay tile={{}} />
    </div>
  );
}

export default App;
