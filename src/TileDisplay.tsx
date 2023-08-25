import React from 'react';
import isEqual from 'lodash/isEqual';

import './TileDisplay.css';
import { Tile } from './TileStore';
import monastery from './icons/monastery.png';
import a_road from './icons/1-road.png';
import c_road from './icons/c-road.png';
import village from './icons/village.png';
import a_city from './icons/1-city.png';
import all_city from './icons/all-city.png';
import half_city from './icons/half-city.png';
import three_city from './icons/3-city.png';
import b_city from './icons/b-city.png';
import a_grass from './icons/1-grass.png';
import top_shield from './icons/top-shield.png';
import mid_shield from './icons/mid-shield.png';
import t_garden from './icons/t-garden.png';
import m_garden from './icons/m-garden.png';
import b_river from './icons/b-river.png';
import c_river from './icons/c-river.png';
import e_river from './icons/e-river.png';
import s_river from './icons/s-river.png';
import cathedral from './icons/cat.png';
import bottom_right_inn from './icons/b-r-inn.png';
import on_top_inn from './icons/mid-inn.png';

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
        return multiImage(image(a_road), image(a_road, 2));
    }
    if (isEqual(r, [0, 1, 0, 2]) && isEqual(c,[1, 0, 2, 0])) {
        return multiImage(image(a_road), image(a_road, 1), image(a_road, 2), image(a_road, 3), image(village));
    }

    const toPrint: React.JSX.Element[] = [];
    if (r.filter(num => num === 1).length == 2 && (r[0] === r[1] || r[1] === r[2])) {
        for (let i = 0; i < 4; i++) {
            if (!!r[i] && r[i] === r[(i+1)%4]) {
                toPrint.push(image(c_road, ((i+2)%4)));
            }
        }
    } else {
        for (let i = 0; i < 4; i++) {
            if (!!r[i]) {
                toPrint.push(image(a_road, ((i+2)%4)));
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
        return multiImage(image(half_city), image(a_city, 2));
    }

    if (isEqual(c,[0, 0, 0, 1]) && f[1] !== f[2]) {
        return multiImage(image(half_city), image(a_grass));
    }

    if (c.filter(num => num === 1).length < 2) {
        const toPrint: React.JSX.Element[] = [];
        for (let i = 0; i < 4; i++) {
            if (!!c[i]) {
                toPrint.push(image(a_city, i));
            }
        }
        return multiImage(...toPrint);
    }

    if (isEqual(c, [1, 1, 1, 1])) {
        return image(all_city);
    }
    if (isEqual(c, [1, 0, 0, 1])) {
        return image(half_city);
    }
    if (isEqual(c, [0, 1, 0, 1])) {
        return image(b_city);
    }
    if (isEqual(c, [1, 1, 0, 1])) {
        return image(three_city, 2);
    }
}

const shield = (c: number[]) => {
    if (isEqual(c, [0, 1, 0, 1])) {
        return image(mid_shield);
    }
    return image(top_shield);
}

const garden = (r: number[], c: number[] ) => {
    if(!isEqual(r, [])) {
        return !r[1] ? image(t_garden, 1) : image(t_garden, 2);
    }
    if( c.length > 0 && !!c[0] && !!c[2]) {
        return image(m_garden);
    }
    return image(t_garden, 2);
}

const river = (ri: boolean[], m: boolean ) => {
    if (isEqual(ri, [false, false, true, false])) {
        return image(e_river);
    }
    if (isEqual(ri, [false, true, false, true])) {
        return m ? image(b_river) : image(s_river);
    }
    if (isEqual(ri, [false, true, true, false])) {
        return image(c_river, 3);
    }
    if (isEqual(ri, [true, true, false, false])) {
        return image(c_river, 2);
    }
}

const inn = (hasCities: boolean) => {
    return hasCities ? image(bottom_right_inn) : image(on_top_inn);
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
