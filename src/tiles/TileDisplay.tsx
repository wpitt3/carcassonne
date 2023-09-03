import React from 'react';
import isEqual from 'lodash/isEqual';

import './TileDisplay.css';
import { Tile } from './Tile';
import cathedral from '../assets/cathedral.png';
import city_full from '../assets/city_full.png';
import city_half from '../assets/city_half.png';
import city_opposite from '../assets/city_opposite.png';
import city_side from '../assets/city_side.png';
import city_three from '../assets/city_three.png';
import garden_mid from '../assets/garden_mid.png';
import garden_top from '../assets/garden_top.png';
import grass_side from '../assets/grass_side.png';
import inn_bottom_right from '../assets/inn_bottom_right.png';
import inn_top from '../assets/inn_top.png';
import monastery from '../assets/monastery.png';
import river_curved from '../assets/river_curved.png';
import river_end from '../assets/river_end.png';
import river_kinked from '../assets/river_kinked.png';
import river_straight from '../assets/river_straight.png';
import road_curved from '../assets/road_curved.png';
import road_half from '../assets/road_half.png';
import shield_mid from '../assets/shield_mid.png';
import shield_top from '../assets/shield_top.png';
import village from '../assets/village.png';

interface Params {
    tile: Tile
}

const image = (imageName: string, rotation? : number): React.JSX.Element => {
    if ( !rotation ) {
        return <img src={imageName}/>
    } else {
        return <img className={"rot"+ 90 * rotation} src={imageName}/>
    }
}

const multiImage = (...images: React.JSX.Element[]): React.JSX.Element => {
    return <div>{images.map((element, i) => <div key={i}>{element}</div>)}</div>
}

const roads = (r: number[], c: number[]) => {
    if (isEqual(r, [0, 0, 1, 0]) && isEqual(c,[1, 0, 0, 0])) {
        return multiImage(image(road_half), image(road_half, 2));
    }
    if (isEqual(r, [0, 1, 0, 2]) && isEqual(c,[1, 0, 2, 0])) {
        return multiImage(image(road_half), image(road_half, 1), image(road_half, 2), image(road_half, 3), image(village));
    }

    const toPrint: React.JSX.Element[] = [];
    if (r.filter(num => num === 1).length == 2 && (r[0] === r[1] || r[1] === r[2])) {
        for (let i = 0; i < 4; i++) {
            if (!!r[i] && r[i] === r[(i+1)%4]) {
                toPrint.push(image(road_curved, ((i+2)%4)));
            }
        }
    } else {
        for (let i = 0; i < 4; i++) {
            if (!!r[i]) {
                toPrint.push(image(road_half, ((i+2)%4)));
            }
        }
        if(r.filter(num => num === 2).length > 0) {
            toPrint.push(image(village));
        }
    }
    return multiImage(...toPrint);
}

const cities = (c: number[], f: number[]) => {
    if (isEqual(c,[1, 0, 2, 1])) {
        return multiImage(image(city_half), image(city_side, 2));
    }

    if (isEqual(c,[0, 0, 0, 1]) && f[1] !== f[2]) {
        return multiImage(image(city_half), image(grass_side));
    }

    if (c.filter(num => num === 1).length < 2) {
        const toPrint: React.JSX.Element[] = [];
        for (let i = 0; i < 4; i++) {
            if (!!c[i]) {
                toPrint.push(image(city_side, i));
            }
        }
        return multiImage(...toPrint);
    }

    if (isEqual(c, [1, 1, 1, 1])) {
        return image(city_full);
    }
    if (isEqual(c, [1, 0, 0, 1])) {
        return image(city_half);
    }
    if (isEqual(c, [0, 1, 0, 1])) {
        return image(city_opposite);
    }
    if (isEqual(c, [1, 1, 0, 1])) {
        return image(city_three, 2);
    }
}

const shield = (c: number[]) => {
    if (isEqual(c, [0, 1, 0, 1])) {
        return image(shield_mid);
    }
    return image(shield_top);
}

const garden = (r: number[], c: number[] ) => {
    if(!isEqual(r, [])) {
        return !r[1] ? image(garden_top, 1) : image(garden_top, 2);
    }
    if( c.length > 0 && !!c[0] && !!c[2]) {
        return image(garden_mid);
    }
    return image(garden_top, 2);
}

const river = (ri: boolean[], m: boolean ) => {
    if (isEqual(ri, [false, false, true, false])) {
        return image(river_end);
    }
    if (isEqual(ri, [false, true, false, true])) {
        return m ? image(river_kinked) : image(river_straight);
    }
    if (isEqual(ri, [false, true, true, false])) {
        return image(river_curved, 3);
    }
    if (isEqual(ri, [true, true, false, false])) {
        return image(river_curved, 2);
    }
}

const inn = (hasCities: boolean) => {
    return hasCities ? image(inn_bottom_right) : image(inn_top);
}

function TileDisplay( {tile}: Params ) {
    return (
        <div className="tile-block">
            { !tile.river || <div className="image-holder">{river(tile.river, !!tile.monastery )}</div>}
            { !tile.roads || <div className="image-holder">{roads(tile.roads, tile.cities || [])}</div>}
            { !tile.cities || <div className="image-holder">{cities(tile.cities, tile.fields || [])}</div>}
            { !tile.shield || !tile.cities || <div className="image-holder">{shield(tile.cities)}</div>}
            { !tile.inn || <div className="image-holder">{inn(!!tile.cities)}</div>}
            { !tile.garden || <div className="image-holder">{garden(tile.roads || [], tile.cities || [])}</div>}
            { !tile.monastery || <div className="image-holder">{image(monastery)}</div>}
            { !tile.cathedral || <div className="image-holder">{image(cathedral)}</div>}
        </div>
    );
}

export default TileDisplay;
