
# Carcassonnne

## Concept
The idea for this repo is to store the representation for each of the tiles and to give a visual representation of each of them.

- [Base Game Tiles](https://wikicarpedia.com/car/Base_game#C2_Edition)
- [The River Tiles](https://wikicarpedia.com/car/River#The_River_I_C2_(Base_game_version))
- [Inns and Cathedrals Game Tiles](https://wikicarpedia.com/car/Inns_and_Cathedrals#C2_Edition)

## Schema
Each tile is made up of 9 different possible icons.
I have represented each item on the tile by an array of a set length 4 (8 for fields) starting on the north side (north north west edge for field) and working around clockwise.

[Schema](src/tiles/Tile.ts)

### Variables
booleans - based on presence
 - cathedral
 - garden
 - monastery

boolean arrays - based on the sides it is present
 - river
 - inn - index of roads with inn boost

number arrays - index representing the side it is present, value representing the group it is part of
 - cities
 - fields
 - roads

numbers - value of affected
 - shield - cities group value with shield boost

### Reasoning for schema decisions
 - There need to be 8 field positions rather than 4 as there are fields which are broken across a corner
 
   ![Gross](https://wikicarpedia.com/images/a/a9/Inns_And_Cathedrals_C2_Tile_G.jpg)
 - Roads need to have an index as multiple roads can be present which do not join

   ![Ewww](https://wikicarpedia.com/images/8/8a/Inns_And_Cathedrals_C2_Tile_E.jpg)
 - Shields need an index as a tile may have multiple cities but only one shield

   ![Vom](https://wikicarpedia.com/images/3/3e/Inns_And_Cathedrals_C2_Tile_P.jpg)
 - Inns need to map to roads as the inn on this tile only affects the top 2 roads

   ![Vom](https://wikicarpedia.com/images/e/e6/Inns_And_Cathedrals_C2_Tile_C.jpg)
### Open Questions
 - Not sure what to do about tile with central field

   ![Vom](https://wikicarpedia.com/images/9/9d/Inns_And_Cathedrals_C2_Tile_H.jpg)
 - Not sure about shield and inn representation, I would prefer if they matched


### Json files for tiles
- [Base Game json](src/tiles/baseTiles.json)
- [The River json](src/tiles/riverTiles.json)
- [Inns and Cathedrals Game json](src/tiles/innsAndCatsTiles.json)

