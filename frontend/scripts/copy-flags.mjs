import {cpSync, existsSync, mkdirSync, readdirSync} from "fs";
import {dirname, join} from "path";
import {fileURLToPath} from "url";

const root = join(dirname(fileURLToPath(import.meta.url)), "..");
const src = join(root, "node_modules/flag-icons/flags/4x3");
const dest = join(root, "public/flags/4x3");

if (!existsSync(src)) {
    console.warn("flag-icons not installed, skipping flag copy");
    process.exit(0);
}

mkdirSync(dest, {recursive: true});

let copied = 0;
for (const file of readdirSync(src)) {
    if (!file.endsWith(".svg")) {
        continue;
    }
    cpSync(join(src, file), join(dest, file));
    copied++;
}

console.log(`Copied ${copied} flag icons to public/flags/4x3`);
