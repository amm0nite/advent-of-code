import fs from 'fs';
import { promisify } from 'util';

export async function getInputLines() {
    const input = await readInput();
    return input.trim().split("\n");
}

export async function readInput() {
    return (promisify(fs.readFile))('input.txt', { encoding: 'utf8' });
}