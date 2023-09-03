import baseTiles from "./baseTiles.json";
import riverTiles from "./riverTiles.json";
import innsAndCatsTiles from "./innsAndCatsTiles.json";
import { Tile } from "./Tile";

export class TileStore {

    baseTiles(): Tile[] {
        return baseTiles as Tile[];
    }

    riverTiles(): Tile[] {
        return riverTiles as Tile[];
    }

    innsAndCatsTiles(): Tile[] {
        return innsAndCatsTiles as Tile[];
    }
}