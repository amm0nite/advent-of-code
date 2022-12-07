import { getInputLines } from "../../input.js";

const lines = await getInputLines();

const pairs = lines.map((line) => line.split(/,|-/).map((element) => parseInt(element)));

const overlaps = pairs.filter((pair) => {
    return (pair[0] >= pair[2] && pair[1] <= pair[3]) || (pair[2] >= pair[0] && pair[3] <= pair[1]);
});

console.log('answer1', overlaps.length); //459

const allOverlaps = pairs.filter((pair) => {
    const distance1 = pair[1] - pair[0];
    const distance2 = pair[3] - pair[2];
    const limit = Math.max(...pair) - Math.min(...pair);
    return (distance1 + distance2) >= limit;
});

console.log('answer2', allOverlaps.length); //779
