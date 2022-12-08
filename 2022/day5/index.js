import { readInput } from "../../input.js";
import { parseStacks } from "./parsing.js";

const lines = (await readInput()).split("\n");

const inlineStackLines = lines.slice(0, 9);
const moveLines = lines.slice(10);

const stacks = parseStacks(inlineStackLines);

console.log(stacks);
console.log(moveLines);