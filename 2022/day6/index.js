import { getInputLines } from "../../input.js";

const lines = await getInputLines();

const data = lines[0];

Array.prototype.repeats = function() {
    for (const current of this) {
        const rest = this.filter((element) => element !== current);
        if (rest.length !== this.length-1) return true;
    }
    return false;
}

function answer(limit) {
    let counter = 0;
    const stack = [];
    for (const c of data) {
        if (stack.length === limit) stack.shift();
        
        stack.push(c);
        counter++;

        if (stack.length !== limit) continue;
        if (!stack.repeats()) return counter;
    }
}

console.log('answer1', answer(4));  //1702
console.log('answer2', answer(14)); //3559