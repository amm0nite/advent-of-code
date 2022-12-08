export function parseStacks(lines) {
    Array.prototype.last = function () { return this[this.length - 1]; };
    const stackCount = parseInt(lines.last().trim().split(/\s+/).last());

    const onlyInlineStacks = lines.slice(0, -1);
    const trimmedInlineStacks = onlyInlineStacks.map((line) => line.substring(1, line.length - 1));
    const cleanInlineStacks = trimmedInlineStacks.reduce((acc, val) => {
        const inlineStack = [];
        for (let i = 0; i < stackCount; i++) {
            inlineStack.push(val[i * 4]);
        }
        acc.push(inlineStack);
        return acc;
    }, []);

    const reversedInlineStacks = cleanInlineStacks.reverse();
    const initialMaxHeight = reversedInlineStacks.length;

    const stacks = [];
    for (let i = 0; i < stackCount; i++) {
        stacks[i] = [];
        for (let j = 0; j < initialMaxHeight; j++) {
            const crate = reversedInlineStacks[j][i];
            if (crate === ' ') break;
            stacks[i][j] = crate;
        }
    }

    return stacks;
}

export function parseMove(movePhrase) {
    const matches = movePhrase.match(/move (\d+) from (\d+) to (\d+)/);
    if (!matches) return null;

    const quantity = parseInt(matches[1]);
    const from = parseInt(matches[2]) - 1;
    const to = parseInt(matches[3]) - 1;

    return { quantity, from, to };
}