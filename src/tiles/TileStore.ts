import baseTiles from "./baseTiles.json";
import riverTiles from "./riverTiles.json";
import innsAndCatsTiles from "./innsAndCatsTiles.json";

export interface Tile {
    cathedral?: boolean,
    cities?: number[], // length 4 max 4
    fields?: number[], // length 8 max 4
    garden?: boolean,
    inn?: boolean[], // length 4
    monastery?: boolean,
    river?: boolean[], // length 4
    roads?: number[], // length 4 max 4
    shield?: number,
}


// 8 numbers, 2 for each half of an edge of a tile, number represents field index
// fields need to be splitable on a tile, then need to not touch when a city breaks between them

// road number represents road index, indexes are required due to C C road
// cities represents city index
// shield represents the index of the city with the shield

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