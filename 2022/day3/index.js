import { getInputLines } from "../../input.js";

const lines = await getInputLines();

const couples = lines.map((line) => [line.substring(0, line.length/2).split(''), line.substring(line.length/2).split('')]);
const collisions = couples.map((couple) => couple[0].find((c1) => couple[1].find((c2) => c1 === c2)));

const lowerCaseOffset = 1 - 'a'.charCodeAt(0);
const upperCaseOffset = 27 - 'A'.charCodeAt(0);

const lowers = collisions.filter((c) => c === c.toLowerCase());
const uppers = collisions.filter((c) => c === c.toUpperCase());

const lowerScores = lowers.map((c) => c.charCodeAt(0) + lowerCaseOffset);
const lowerTotalScore = lowerScores.reduce((acc, val) => acc += val, 0);

const upperScores = uppers.map((c) => c.charCodeAt(0) + upperCaseOffset);
const upperTotalScore = upperScores.reduce((acc, val) => acc += val, 0);

console.log('answer1', lowerTotalScore + upperTotalScore); //7917

Array.prototype.last = function() { return this[this.length-1]; };

const groups = lines.reduce((acc, val) => {
    if (acc.last().length === 3) acc.push([]);
    acc.last().push(val);
    return acc;
}, [[]]);

console.log(groups);