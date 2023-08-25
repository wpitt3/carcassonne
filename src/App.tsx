import React, { useState } from 'react';
import './App.css';
import TileDisplay from './TileDisplay';
import {TileStore} from './TileStore';

function App() {
  const [page, setPage] = useState < string > ('build');

  const baseTiles = new TileStore().baseTiles();
  const riverTiles = new TileStore().riverTiles();
  const innsAndCatsTiles = new TileStore().innsAndCatsTiles();
  return (
    <div>
        <div className="tile-block-wrapper">
          {baseTiles.map((tile, i) => <TileDisplay tile={tile}/>)}
        </div>
        <br/>
        <div className="tile-block-wrapper">
          {riverTiles.map((tile, i) => <TileDisplay tile={tile}/>)}
        </div>
        <div className="tile-block-wrapper">
          {innsAndCatsTiles.map((tile, i) => <TileDisplay tile={tile}/>)}
        </div>
    </div>
  );
}

export default App;
