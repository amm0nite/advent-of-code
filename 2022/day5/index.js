import { readInput } from "../../input.js";
import { parseStacks, parseMove } from "./parsing.js";

const lines = (await readInput()).split("\n");

const inlineStackLines = lines.slice(0, 9);
const moveLines = lines.slice(10);

function answer1() {
    const stacks = parseStacks(inlineStackLines);

    for (const movePhrase of moveLines) {
        const move = parseMove(movePhrase);
        if (!move) continue;

        for (let i=0; i<move.quantity; i++) {
            const crate = stacks[move.from].pop();
            stacks[move.to].push(crate);
        }
    }

    const tops = stacks.map((stack) => stack.pop());
    console.log('answer1', tops.join(''));
}

function answer2() {
    const stacks = parseStacks(inlineStackLines);

    for (const movePhrase of moveLines) {
        const move = parseMove(movePhrase);
        if (!move) continue;

        const bundle = stacks[move.from].splice(-move.quantity);
        stacks[move.to] = stacks[move.to].concat(bundle);
    }

    const tops = stacks.map((stack) => stack.pop());
    console.log('answer2', tops.join(''));
}

answer1(); //TWSGQHNHL
answer2(); //JNRSCDWPP