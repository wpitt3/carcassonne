

export interface Tile {
    fields: number[], // length 8 max 4
    monastery: boolean,
    garden: boolean,
    roads: number[], // length 4 max 4
    cities: number[], // length 4 max 4
    shield: number,
    river: boolean[], // length 4
    inn: boolean[], // length 4
}

// 8 numbers, 2 for each half of an edge of a tile, number represents field index
// fields need to be splitable on a tile, then need to not touch when a city breaks between them

// road number represents road index, indexes are required due to C C road
// cities represents city index
// shield represents the index of the city with the shield
