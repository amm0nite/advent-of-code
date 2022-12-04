import fs from 'fs';
import { promisify } from 'util';

export async function getInputLines() {
    const input = await (promisify(fs.readFile))('input.txt', { encoding: 'utf8' });
    return input.trim().split("\n");
}