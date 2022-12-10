import { getInputLines } from "../../input.js";

const lines = await getInputLines();

const data = lines[0];
const stack = [];

Array.prototype.repeats = function() {
    for (const current of this) {
        const rest = this.filter((element) => element !== current);
        if (rest.length !== this.length-1) return true;
    }
    return false;
}

function answer1() {
    let counter = 0;
    for (const c of data) {
        if (stack.length === 4) stack.shift();
        
        stack.push(c);
        counter++;

        if (stack.length !== 4) continue;
        if (!stack.repeats()) return counter;
    }
}

console.log('answer1', answer1()); //1702