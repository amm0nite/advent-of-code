import { getInputLines } from "../../input.js";

const lines = await getInputLines();

const couples = lines.map((line) => [line.substring(0, line.length/2).split(''), line.substring(line.length/2).split('')]);

Array.prototype.findCommonElement = function(...others) {
    return this.find((element) => others.every((other) => other.includes(element)));
};

const collisions = couples.map((couple) => couple[0].findCommonElement(couple[1]));

const lowerCaseOffset = 1 - 'a'.charCodeAt(0);
const upperCaseOffset = 27 - 'A'.charCodeAt(0);

const lowers = collisions.filter((c) => c === c.toLowerCase());
const uppers = collisions.filter((c) => c === c.toUpperCase());

Array.prototype.sum = function() { return this.reduce((acc, val) => acc += val, 0); }

const lowerScores = lowers.map((c) => c.charCodeAt(0) + lowerCaseOffset);
const upperScores = uppers.map((c) => c.charCodeAt(0) + upperCaseOffset);

console.log('answer1', lowerScores.sum() + upperScores.sum()); //7917

Array.prototype.last = function() { return this[this.length-1]; };

const groups = lines.reduce((acc, val) => {
    if (acc.last().length === 3) acc.push([]);
    acc.last().push(val.split(''));
    return acc;
}, [[]]);

const badges = groups.map((group) => group[0].findCommonElement(group[1], group[2]));

const badgeLowers = badges.filter((c) => c === c.toLowerCase());
const badgeUppers = badges.filter((c) => c === c.toUpperCase());

const badgeLowerScores = badgeLowers.map((c) => c.charCodeAt(0) + lowerCaseOffset);
const badgeUpperScores = badgeUppers.map((c) => c.charCodeAt(0) + upperCaseOffset);

console.log('answer2', badgeLowerScores.sum() + badgeUpperScores.sum()); //2585