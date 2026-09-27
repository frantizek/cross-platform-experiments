/**
 * @fileoverview Recursively reads all files in a directory and its subdirectories.
 * Demonstrates Node.js's async/await and non-blocking I/O.
 */

const fs = require('fs').promises;
const path = require('path');
const os = require('os');
const { format } = require('date-fns');

/**
 * Prints system information such as OS, architecture, CPU count, and current time.
 */
async function printSystemInfo() {
    console.log("--- System Info ---");
    console.log(`OS: ${os.platform()}`);
    console.log(`Architecture: ${os.arch()}`);
    console.log(`Number of CPUs: ${os.cpus().length}`);
    console.log(`Current Time: ${format(new Date(), 'EEE, dd MMM yyyy HH:mm:ss zzz')}`);
    console.log("-------------------");
}

/**
 * Recursively reads all files in a directory and its subdirectories.
 * @param {string} dir - The directory to start reading from.
 * @param {Array<string>} fileList - Accumulator for file paths.
 * @returns {Promise<Array<string>>} - List of file paths.
 */
async function readFilesRecursively(dir, fileList = []) {
    const files = await fs.readdir(dir);
    for (const file of files) {
        const filePath = path.join(dir, file);
        const stat = await fs.stat(filePath);
        if (stat.isDirectory()) {
            await readFilesRecursively(filePath, fileList);
        } else {
            fileList.push(filePath);
        }
    }
    return fileList;
}

/**
 * Main function to demonstrate recursive file reading.
 */
async function main() {
    console.log("Hello from Node.js!");

    await printSystemInfo();

    console.log("\nRecursively reading files in current directory...");
    try {
        const files = await readFilesRecursively(process.cwd());
        console.log(`Found ${files.length} files:`);
        files.forEach(file => console.log(`  - ${file}`));
    } catch (err) {
        console.error("Error reading files:", err);
    }

    console.log("\nExiting...");
}

main();