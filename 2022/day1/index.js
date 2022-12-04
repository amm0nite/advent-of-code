import { getInputLines } from "../../input.js";

const lines = await getInputLines();
const elves = lines.reduce((accumulator, value) => {
    if (value !== '') accumulator[accumulator.length - 1] += parseInt(value);
    else accumulator.push(0);
    return accumulator;
}, [0]);

console.log('answer1', Math.max(...elves));

elves.sort((a, b) => a - b);
const top3 = elves.slice(-3);
const sum = top3.reduce((accumulator, value) => accumulator += value, 0);

console.log('answer2', sum);