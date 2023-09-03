import './Tiles.css';
import TileDisplay from './TileDisplay';
import {TileStore} from './TileStore';

function Tiles() {
  const baseTiles = new TileStore().getBaseTiles();
  const riverTiles = new TileStore().getRiverTiles();
  const innsAndCatsTiles = new TileStore().getInnsAndCatsTiles();
  return (
    <div>
        <div className="tile-block-wrapper">
          {baseTiles.map((tile, i) => <TileDisplay key={i} tile={tile}/>)}
        </div>
        <br/>
        <div className="tile-block-wrapper">
          {riverTiles.map((tile, i) => <TileDisplay key={i} tile={tile}/>)}
        </div>
        <div className="tile-block-wrapper">
          {innsAndCatsTiles.map((tile, i) => <TileDisplay key={i} tile={tile}/>)}
        </div>
    </div>
  );
}

export default Tiles;
