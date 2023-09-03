import baseTiles from "../assets/json/baseTiles.json";
import riverTiles from "../assets/json/riverTiles.json";
import innsAndCatsTiles from "../assets/json/innsAndCatsTiles.json";
import { Tile } from "./Tile";

interface TileWithMeta {
    id: number,
    count: number,
    props: Tile
}

export class TileStore {
    private baseTiles: TileWithMeta[];
    private riverTiles: TileWithMeta[];
    private innsAndCatsTiles: TileWithMeta[];

    constructor() {
        this.baseTiles = baseTiles as TileWithMeta[];
        this.riverTiles = riverTiles as TileWithMeta[];
        this.innsAndCatsTiles = innsAndCatsTiles as TileWithMeta[];
    }

    getBaseTiles(): Tile[] {
        return this.baseTiles.map((it) => it.props);
    }

    getRiverTiles(): Tile[] {
        return this.riverTiles.map((it) => it.props);
    }

    getInnsAndCatsTiles(): Tile[] {
        return this.innsAndCatsTiles.map((it) => it.props);
    }
}