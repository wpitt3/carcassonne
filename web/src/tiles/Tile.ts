
export interface Tile {
    cathedral?: boolean,
    cities?: number[], // length 4
    fields?: number[], // length 8
    garden?: boolean,
    inn?: boolean[], // length 4
    monastery?: boolean,
    river?: boolean[], // length 4
    roads?: number[], // length 4
    shield?: number,
}

